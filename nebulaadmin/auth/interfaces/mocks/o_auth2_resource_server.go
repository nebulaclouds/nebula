// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/nebulaclouds/nebula/nebulaadmin/auth/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// OAuth2ResourceServer is an autogenerated mock type for the OAuth2ResourceServer type
type OAuth2ResourceServer struct {
	mock.Mock
}

type OAuth2ResourceServer_ValidateAccessToken struct {
	*mock.Call
}

func (_m OAuth2ResourceServer_ValidateAccessToken) Return(_a0 interfaces.IdentityContext, _a1 error) *OAuth2ResourceServer_ValidateAccessToken {
	return &OAuth2ResourceServer_ValidateAccessToken{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2ResourceServer) OnValidateAccessToken(ctx context.Context, expectedAudience string, tokenStr string) *OAuth2ResourceServer_ValidateAccessToken {
	c_call := _m.On("ValidateAccessToken", ctx, expectedAudience, tokenStr)
	return &OAuth2ResourceServer_ValidateAccessToken{Call: c_call}
}

func (_m *OAuth2ResourceServer) OnValidateAccessTokenMatch(matchers ...interface{}) *OAuth2ResourceServer_ValidateAccessToken {
	c_call := _m.On("ValidateAccessToken", matchers...)
	return &OAuth2ResourceServer_ValidateAccessToken{Call: c_call}
}

// ValidateAccessToken provides a mock function with given fields: ctx, expectedAudience, tokenStr
func (_m *OAuth2ResourceServer) ValidateAccessToken(ctx context.Context, expectedAudience string, tokenStr string) (interfaces.IdentityContext, error) {
	ret := _m.Called(ctx, expectedAudience, tokenStr)

	var r0 interfaces.IdentityContext
	if rf, ok := ret.Get(0).(func(context.Context, string, string) interfaces.IdentityContext); ok {
		r0 = rf(ctx, expectedAudience, tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.IdentityContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, expectedAudience, tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
