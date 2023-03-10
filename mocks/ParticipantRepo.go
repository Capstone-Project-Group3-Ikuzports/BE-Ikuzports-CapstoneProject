// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	participant "ikuzports/features/participant"

	mock "github.com/stretchr/testify/mock"
)

// ParticipantRepo is an autogenerated mock type for the RepositoryInterface type
type ParticipantRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *ParticipantRepo) Create(data participant.ParticipantCore) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(participant.ParticipantCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(participant.ParticipantCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMember provides a mock function with given fields: eventID, userID
func (_m *ParticipantRepo) FindMember(eventID int, userID int) (participant.ParticipantCore, error) {
	ret := _m.Called(eventID, userID)

	var r0 participant.ParticipantCore
	if rf, ok := ret.Get(0).(func(int, int) participant.ParticipantCore); ok {
		r0 = rf(eventID, userID)
	} else {
		r0 = ret.Get(0).(participant.ParticipantCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(eventID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateParticipant provides a mock function with given fields: data
func (_m *ParticipantRepo) UpdateParticipant(data participant.ParticipantCore) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(participant.ParticipantCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(participant.ParticipantCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: data
func (_m *ParticipantRepo) UpdateStatus(data participant.ParticipantCore) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(participant.ParticipantCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(participant.ParticipantCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewParticipantRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewParticipantRepo creates a new instance of ParticipantRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewParticipantRepo(t mockConstructorTestingTNewParticipantRepo) *ParticipantRepo {
	mock := &ParticipantRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
