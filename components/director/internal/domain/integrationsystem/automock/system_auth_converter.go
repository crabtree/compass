// Code generated by mockery 2.9.0. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// SystemAuthConverter is an autogenerated mock type for the SystemAuthConverter type
type SystemAuthConverter struct {
	mock.Mock
}

// ToGraphQL provides a mock function with given fields: in
func (_m *SystemAuthConverter) ToGraphQL(in *model.SystemAuth) (graphql.SystemAuth, error) {
	ret := _m.Called(in)

	var r0 graphql.SystemAuth
	if rf, ok := ret.Get(0).(func(*model.SystemAuth) graphql.SystemAuth); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graphql.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SystemAuth) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
