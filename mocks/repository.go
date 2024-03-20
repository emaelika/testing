// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	jobs "testing/jobs"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: newJobs
func (_m *Repository) Create(newJobs jobs.Jobs) (jobs.Jobs, error) {
	ret := _m.Called(newJobs)

	var r0 jobs.Jobs
	var r1 error
	if rf, ok := ret.Get(0).(func(jobs.Jobs) (jobs.Jobs, error)); ok {
		return rf(newJobs)
	}
	if rf, ok := ret.Get(0).(func(jobs.Jobs) jobs.Jobs); ok {
		r0 = rf(newJobs)
	} else {
		r0 = ret.Get(0).(jobs.Jobs)
	}

	if rf, ok := ret.Get(1).(func(jobs.Jobs) error); ok {
		r1 = rf(newJobs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: jobID, role
func (_m *Repository) GetJob(jobID uint, role string) (jobs.Jobs, error) {
	ret := _m.Called(jobID, role)

	var r0 jobs.Jobs
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) (jobs.Jobs, error)); ok {
		return rf(jobID, role)
	}
	if rf, ok := ret.Get(0).(func(uint, string) jobs.Jobs); ok {
		r0 = rf(jobID, role)
	} else {
		r0 = ret.Get(0).(jobs.Jobs)
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(jobID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// // GetJobs provides a mock function with given fields: userID, role, page, pagesize
// func (_m *Repository) GetJobs(userID uint, role string, page int, pagesize int) ([]jobs.Jobs, int, error) {
// 	ret := _m.Called(userID, role, page, pagesize)

// 	var r0 []jobs.Jobs
// 	var r1 int
// 	var r2 error
// 	if rf, ok := ret.Get(0).(func(uint, string, int, int) ([]jobs.Jobs, int, error)); ok {
// 		return rf(userID, role, page, pagesize)
// 	}
// 	if rf, ok := ret.Get(0).(func(uint, string, int, int) []jobs.Jobs); ok {
// 		r0 = rf(userID, role, page, pagesize)
// 	} else {
// 		if ret.Get(0) != nil {
// 			r0 = ret.Get(0).([]jobs.Jobs)
// 		}
// 	}

// 	if rf, ok := ret.Get(1).(func(uint, string, int, int) int); ok {
// 		r1 = rf(userID, role, page, pagesize)
// 	} else {
// 		r1 = ret.Get(1).(int)
// 	}

// 	if rf, ok := ret.Get(2).(func(uint, string, int, int) error); ok {
// 		r2 = rf(userID, role, page, pagesize)
// 	} else {
// 		r2 = ret.Error(2)
// 	}

// 	return r0, r1, r2
// }

// // GetJobsByStatus provides a mock function with given fields: userID, status, role, page, pagesize
// func (_m *Repository) GetJobsByStatus(userID uint, status string, role string, page int, pagesize int) ([]jobs.Jobs, int, error) {
// 	ret := _m.Called(userID, status, role, page, pagesize)

// 	var r0 []jobs.Jobs
// 	var r1 int
// 	var r2 error
// 	if rf, ok := ret.Get(0).(func(uint, string, string, int, int) ([]jobs.Jobs, int, error)); ok {
// 		return rf(userID, status, role, page, pagesize)
// 	}
// 	if rf, ok := ret.Get(0).(func(uint, string, string, int, int) []jobs.Jobs); ok {
// 		r0 = rf(userID, status, role, page, pagesize)
// 	} else {
// 		if ret.Get(0) != nil {
// 			r0 = ret.Get(0).([]jobs.Jobs)
// 		}
// 	}

// 	if rf, ok := ret.Get(1).(func(uint, string, string, int, int) int); ok {
// 		r1 = rf(userID, status, role, page, pagesize)
// 	} else {
// 		r1 = ret.Get(1).(int)
// 	}

// 	if rf, ok := ret.Get(2).(func(uint, string, string, int, int) error); ok {
// 		r2 = rf(userID, status, role, page, pagesize)
// 	} else {
// 		r2 = ret.Error(2)
// 	}

// 	return r0, r1, r2
// }

// // UpdateJob provides a mock function with given fields: update
// func (_m *Repository) UpdateJob(update jobs.Jobs) (jobs.Jobs, error) {
// 	ret := _m.Called(update)

// 	var r0 jobs.Jobs
// 	var r1 error
// 	if rf, ok := ret.Get(0).(func(jobs.Jobs) (jobs.Jobs, error)); ok {
// 		return rf(update)
// 	}
// 	if rf, ok := ret.Get(0).(func(jobs.Jobs) jobs.Jobs); ok {
// 		r0 = rf(update)
// 	} else {
// 		r0 = ret.Get(0).(jobs.Jobs)
// 	}

// 	if rf, ok := ret.Get(1).(func(jobs.Jobs) error); ok {
// 		r1 = rf(update)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return r0, r1
// }

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
