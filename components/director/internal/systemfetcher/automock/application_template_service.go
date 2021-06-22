// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// ApplicationTemplateService is an autogenerated mock type for the ApplicationTemplateService type
type ApplicationTemplateService struct {
	mock.Mock
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *ApplicationTemplateService) GetByName(ctx context.Context, name string) (*model.ApplicationTemplate, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.ApplicationTemplate
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ApplicationTemplate); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ApplicationTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}