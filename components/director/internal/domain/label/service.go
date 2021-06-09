package label

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/pkg/apperrors"
	"github.com/kyma-incubator/compass/components/director/pkg/jsonschema"
	"github.com/kyma-incubator/compass/components/director/pkg/log"
	"github.com/kyma-incubator/compass/components/director/pkg/str"
	"github.com/pkg/errors"
)

//go:generate mockery --name=LabelRepository --output=automock --outpkg=automock --case=underscore
type LabelRepository interface {
	Upsert(ctx context.Context, label *model.Label) error
	GetByKey(ctx context.Context, tenant string, objectType model.LabelableObject, objectID, key string) (*model.Label, error)
	Delete(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, key string) error
	ListByObjectTypeAndMatchAnyScenario(ctx context.Context, tenantId string, objectType model.LabelableObject, scenarios []string) ([]model.Label, error)
	GetBundleInstanceAuthsScenarioLabels(ctx context.Context, tenant, appId, runtimeId string) ([]model.Label, error)
}

//go:generate mockery --name=LabelDefinitionRepository --output=automock --outpkg=automock --case=underscore
type LabelDefinitionRepository interface {
	Create(ctx context.Context, def model.LabelDefinition) error
	Exists(ctx context.Context, tenant string, key string) (bool, error)
	GetByKey(ctx context.Context, tenant string, key string) (*model.LabelDefinition, error)
}

//go:generate mockery --name=UIDService --output=automock --outpkg=automock --case=underscore
type UIDService interface {
	Generate() string
}

type labelUpsertService struct {
	labelRepo           LabelRepository
	labelDefinitionRepo LabelDefinitionRepository
	uidService          UIDService
}

func NewLabelUpsertService(labelRepo LabelRepository, labelDefinitionRepo LabelDefinitionRepository, uidService UIDService) *labelUpsertService {
	return &labelUpsertService{labelRepo: labelRepo, labelDefinitionRepo: labelDefinitionRepo, uidService: uidService}
}

func (s *labelUpsertService) UpsertMultipleLabels(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, labels map[string]interface{}) error {
	for key, val := range labels {
		err := s.UpsertLabel(ctx, tenant, &model.LabelInput{
			Key:        key,
			Value:      val,
			ObjectID:   objectID,
			ObjectType: objectType,
		})
		if err != nil {
			return errors.Wrap(err, "while upserting multiple Labels")
		}
	}

	return nil
}

func (s *labelUpsertService) UpsertLabel(ctx context.Context, tenant string, labelInput *model.LabelInput) error {
	var labelDef *model.LabelDefinition

	labelDef, err := s.labelDefinitionRepo.GetByKey(ctx, tenant, labelInput.Key)
	if err != nil && !apperrors.IsNotFoundError(err) {
		return errors.Wrapf(err, "while reading LabelDefinition for key '%s'", labelInput.Key)
	}

	if labelDef == nil {
		// Create new LabelDefinition
		labelDefinitionID := s.uidService.Generate()
		labelDef = &model.LabelDefinition{
			ID:     labelDefinitionID,
			Tenant: tenant,
			Key:    labelInput.Key,
			Schema: nil,
		}
		err := s.labelDefinitionRepo.Create(ctx, *labelDef)
		if err != nil {
			return errors.Wrapf(err, "while creating a new LabelDefinition for Label with key: '%s'", labelInput.Key)
		}
		log.C(ctx).Debugf("Successfully created LabelDefinition with id %s and key %s for Label with key %s", labelDef.ID, labelDef.Key, labelInput.Key)
	}

	err = s.validateLabelInputValue(ctx, tenant, labelInput, labelDef)
	if err != nil {
		return errors.Wrapf(err, "while validating Label value for '%s'", labelInput.Key)
	}

	label := labelInput.ToLabel(s.uidService.Generate(), tenant)

	err = s.labelRepo.Upsert(ctx, label)
	if err != nil {
		return errors.Wrapf(err, "while creating Label with id %s for %s with id %s", label.ID, label.ObjectType, label.ObjectID)
	}
	log.C(ctx).Debugf("Successfully created Label with id %s for %s with id %s", label.ID, label.ObjectType, label.ObjectID)

	return nil
}

func (s *labelUpsertService) UpsertScenarios(ctx context.Context, tenantID string, labels []model.Label, newScenarios []string, mergeFn func(scenarios, diffScenario []string) []string) error {
	for _, label := range labels {
		if model.ScenariosKey != label.Key {
			continue
		}

		scenariosString, err := GetScenariosAsStringSlice(label)
		if err != nil {
			return err
		}

		scenariosToUpsert := mergeFn(scenariosString, newScenarios)

		err = s.updateScenario(ctx, tenantID, label, scenariosToUpsert)
		if err != nil {
			return errors.Wrap(err, "while updating scenarios label")
		}
	}
	return nil
}

func GetScenariosAsStringSlice(label model.Label) ([]string, error) {
	return GetScenariosFromValueAsStringSlice(label.Value)
}

func GetScenariosFromValueAsStringSlice(labelValue interface{}) ([]string, error) {
	var result []string

	switch value := labelValue.(type) {
	case []string:
		result = value
	case []interface{}:
		convertedScenarios, err := str.InterfaceSliceToStringSlice(value)
		if err != nil {
			return nil, errors.Wrap(err, "while converting array of interfaces to array of strings")
		}
		result = convertedScenarios
	default:
		return nil, errors.Errorf("scenarios value is invalid type: %t", labelValue)
	}

	return result, nil
}

func UniqueScenarios(scenarios, newScenarios []string) []string {
	scenarios = append(scenarios, newScenarios...)
	return str.Unique(scenarios)
}

func (s *labelUpsertService) updateScenario(ctx context.Context, tenantID string, label model.Label, scenarios []string) error {
	if len(scenarios) == 0 {
		return s.labelRepo.Delete(ctx, tenantID, label.ObjectType, label.ObjectID, model.ScenariosKey)
	} else {
		labelInput := model.LabelInput{
			Key:        label.Key,
			Value:      scenarios,
			ObjectID:   label.ObjectID,
			ObjectType: label.ObjectType,
		}
		return s.UpsertLabel(ctx, tenantID, &labelInput)
	}
}

func (s *labelUpsertService) validateLabelInputValue(ctx context.Context, tenant string, labelInput *model.LabelInput, labelDef *model.LabelDefinition) error {
	if labelDef == nil || labelDef.Schema == nil {
		// nothing to validate
		return nil
	}

	validator, err := jsonschema.NewValidatorFromRawSchema(*labelDef.Schema)
	if err != nil {
		return errors.Wrapf(err, "while creating JSON Schema validator for schema %+v", *labelDef.Schema)
	}

	jsonSchema, err := json.Marshal(*labelDef.Schema)
	if err != nil {
		return apperrors.InternalErrorFrom(err, "while marshalling json schema")
	}

	result, err := validator.ValidateRaw(labelInput.Value)
	if err != nil {
		return apperrors.InternalErrorFrom(err, "while validating value=%+v against JSON Schema=%s", labelInput.Value, string(jsonSchema))
	}
	if !result.Valid {
		return apperrors.NewInvalidDataError(fmt.Sprintf("input value=%+v, key=%s, is not valid against JSON Schema=%s,result=%s", labelInput.Value, labelInput.Key, jsonSchema, result.Error.Error()))
	}

	return nil
}
