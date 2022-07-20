// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/ijalalfrz/sirclo-weight-test/model"
	mock "github.com/stretchr/testify/mock"

	response "github.com/ijalalfrz/sirclo-weight-test/response"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// FindMany provides a mock function with given fields: ctx
func (_m *Usecase) FindMany(ctx context.Context) response.Response {
	ret := _m.Called(ctx)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context) response.Response); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

// FindOne provides a mock function with given fields: ctx, key
func (_m *Usecase) FindOne(ctx context.Context, key int64) response.Response {
	ret := _m.Called(ctx, key)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64) response.Response); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

// InsertOne provides a mock function with given fields: ctx, payload
func (_m *Usecase) InsertOne(ctx context.Context, payload model.WeightPayload) response.Response {
	ret := _m.Called(ctx, payload)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, model.WeightPayload) response.Response); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

// UpdateOne provides a mock function with given fields: ctx, key, payload
func (_m *Usecase) UpdateOne(ctx context.Context, key int64, payload model.WeightPayload) response.Response {
	ret := _m.Called(ctx, key, payload)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64, model.WeightPayload) response.Response); ok {
		r0 = rf(ctx, key, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}