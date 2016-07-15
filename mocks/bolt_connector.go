package mocks

import (
	"github.com/klarna/eremetic/types"
	"github.com/stretchr/testify/mock"
)

// BoltConnectorInterface is an autogenerated mock type for the BoltConnectorInterface type
type BoltConnectorInterface struct {
	mock.Mock
}

// Open provides a mock function with given fields: path
func (_m *BoltConnectorInterface) Open(path string) (types.BoltConnection, error) {
	ret := _m.Called(path)

	var r0 types.BoltConnection
	if rf, ok := ret.Get(0).(func(string) types.BoltConnection); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.BoltConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
