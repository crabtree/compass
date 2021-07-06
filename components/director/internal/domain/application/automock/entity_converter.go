// Code generated by mockery 2.9.0. DO NOT EDIT.

package automock

import (
	application "github.com/kyma-incubator/compass/components/director/internal/domain/application"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// EntityConverter is an autogenerated mock type for the EntityConverter type
type EntityConverter struct {
	mock.Mock
}

// FromEntity provides a mock function with given fields: entity
func (_m *EntityConverter) FromEntity(entity *application.Entity) *model.Application {
	ret := _m.Called(entity)

	var r0 *model.Application
	if rf, ok := ret.Get(0).(func(*application.Entity) *model.Application); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Application)
		}
	}

	return r0
}

// ToEntity provides a mock function with given fields: in
func (_m *EntityConverter) ToEntity(in *model.Application) (*application.Entity, error) {
	ret := _m.Called(in)

	var r0 *application.Entity
	if rf, ok := ret.Get(0).(func(*model.Application) *application.Entity); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.Entity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Application) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
