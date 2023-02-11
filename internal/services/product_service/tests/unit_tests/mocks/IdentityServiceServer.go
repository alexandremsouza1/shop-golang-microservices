// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	identity_service "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/product/grpc_client/protos"
	mock "github.com/stretchr/testify/mock"
)

// IdentityServiceServer is an autogenerated mock type for the IdentityServiceServer type
type IdentityServiceServer struct {
	mock.Mock
}

// GetUserById provides a mock function with given fields: _a0, _a1
func (_m *IdentityServiceServer) GetUserById(_a0 context.Context, _a1 *identity_service.GetUserByIdReq) (*identity_service.GetUserByIdRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *identity_service.GetUserByIdRes
	if rf, ok := ret.Get(0).(func(context.Context, *identity_service.GetUserByIdReq) *identity_service.GetUserByIdRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identity_service.GetUserByIdRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *identity_service.GetUserByIdReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIdentityServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewIdentityServiceServer creates a new instance of IdentityServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIdentityServiceServer(t mockConstructorTestingTNewIdentityServiceServer) *IdentityServiceServer {
	mock := &IdentityServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}