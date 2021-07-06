// Code generated by mockery 2.9.0. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// ModelConverter is an autogenerated mock type for the ModelConverter type
type ModelConverter struct {
	mock.Mock
}

// FromGraphQL provides a mock function with given fields: input, tenant
func (_m *ModelConverter) FromGraphQL(input graphql.LabelDefinitionInput, tenant string) (model.LabelDefinition, error) {
	ret := _m.Called(input, tenant)

	var r0 model.LabelDefinition
	if rf, ok := ret.Get(0).(func(graphql.LabelDefinitionInput, string) model.LabelDefinition); ok {
		r0 = rf(input, tenant)
	} else {
		r0 = ret.Get(0).(model.LabelDefinition)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.LabelDefinitionInput, string) error); ok {
		r1 = rf(input, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGraphQL provides a mock function with given fields: definition
func (_m *ModelConverter) ToGraphQL(definition model.LabelDefinition) (graphql.LabelDefinition, error) {
	ret := _m.Called(definition)

	var r0 graphql.LabelDefinition
	if rf, ok := ret.Get(0).(func(model.LabelDefinition) graphql.LabelDefinition); ok {
		r0 = rf(definition)
	} else {
		r0 = ret.Get(0).(graphql.LabelDefinition)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.LabelDefinition) error); ok {
		r1 = rf(definition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
