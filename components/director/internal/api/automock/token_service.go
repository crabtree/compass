// Code generated by mockery 2.9.0. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"

	tokens "github.com/kyma-incubator/compass/components/director/internal/tokens"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// RegenerateOneTimeToken provides a mock function with given fields: ctx, authID, token
func (_m *TokenService) RegenerateOneTimeToken(ctx context.Context, authID string, token tokens.TokenType) (model.OneTimeToken, error) {
	ret := _m.Called(ctx, authID, token)

	var r0 model.OneTimeToken
	if rf, ok := ret.Get(0).(func(context.Context, string, tokens.TokenType) model.OneTimeToken); ok {
		r0 = rf(ctx, authID, token)
	} else {
		r0 = ret.Get(0).(model.OneTimeToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, tokens.TokenType) error); ok {
		r1 = rf(ctx, authID, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
