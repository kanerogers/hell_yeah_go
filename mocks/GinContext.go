package mocks

import "github.com/stretchr/testify/mock"

import "github.com/stretchr/testify/mock"

type GinContext struct {
	mock.Mock
}

func (m *GinContext) JSON(code int, obj interface{}) (bool, error) {
	ret := m.Called(code, obj)

	r0 := ret.Get(0).(bool)
	r1 := ret.Error(1)

	return r0, r1
}
