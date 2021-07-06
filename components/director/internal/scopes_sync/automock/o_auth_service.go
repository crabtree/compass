// Code generated by mockery 2.9.0. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"

	models "github.com/ory/hydra-client-go/models"

	oauth20 "github.com/kyma-incubator/compass/components/director/internal/domain/oauth20"
)

// OAuthService is an autogenerated mock type for the OAuthService type
type OAuthService struct {
	mock.Mock
}

// GetClientDetails provides a mock function with given fields: objType
func (_m *OAuthService) GetClientDetails(objType model.SystemAuthReferenceObjectType) (*oauth20.ClientDetails, error) {
	ret := _m.Called(objType)

	var r0 *oauth20.ClientDetails
	if rf, ok := ret.Get(0).(func(model.SystemAuthReferenceObjectType) *oauth20.ClientDetails); ok {
		r0 = rf(objType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth20.ClientDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SystemAuthReferenceObjectType) error); ok {
		r1 = rf(objType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClients provides a mock function with given fields:
func (_m *OAuthService) ListClients() ([]*models.OAuth2Client, error) {
	ret := _m.Called()

	var r0 []*models.OAuth2Client
	if rf, ok := ret.Get(0).(func() []*models.OAuth2Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.OAuth2Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClient provides a mock function with given fields: ctx, clientID, objectType
func (_m *OAuthService) UpdateClient(ctx context.Context, clientID string, objectType model.SystemAuthReferenceObjectType) error {
	ret := _m.Called(ctx, clientID, objectType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.SystemAuthReferenceObjectType) error); ok {
		r0 = rf(ctx, clientID, objectType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
