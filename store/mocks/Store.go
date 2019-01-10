// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "gitlack/model"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateIssue provides a mock function with given fields: _a0
func (_m *Store) CreateIssue(_a0 *model.Issue) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Issue) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMergeRequest provides a mock function with given fields: _a0
func (_m *Store) CreateMergeRequest(_a0 *model.MergeRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.MergeRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProject provides a mock function with given fields: _a0
func (_m *Store) CreateProject(_a0 *model.Project) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Project) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: _a0
func (_m *Store) CreateUser(_a0 *model.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIssue provides a mock function with given fields: _a0, _a1
func (_m *Store) GetIssue(_a0 int, _a1 int) (*model.Issue, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Issue
	if rf, ok := ret.Get(0).(func(int, int) *model.Issue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Issue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMergeRequest provides a mock function with given fields: _a0, _a1
func (_m *Store) GetMergeRequest(_a0 int, _a1 int) (*model.MergeRequest, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.MergeRequest
	if rf, ok := ret.Get(0).(func(int, int) *model.MergeRequest); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MergeRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByID provides a mock function with given fields: _a0
func (_m *Store) GetProjectByID(_a0 int) (*model.Project, error) {
	ret := _m.Called(_a0)

	var r0 *model.Project
	if rf, ok := ret.Get(0).(func(int) *model.Project); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByPath provides a mock function with given fields: _a0
func (_m *Store) GetProjectByPath(_a0 string) (*model.Project, error) {
	ret := _m.Called(_a0)

	var r0 *model.Project
	if rf, ok := ret.Get(0).(func(string) *model.Project); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: _a0
func (_m *Store) GetUserByEmail(_a0 string) (*model.User, error) {
	ret := _m.Called(_a0)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: _a0
func (_m *Store) GetUserByID(_a0 int) (*model.User, error) {
	ret := _m.Called(_a0)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(int) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProjectDefaultChannel provides a mock function with given fields: _a0, _a1
func (_m *Store) UpdateProjectDefaultChannel(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserDefaultChannel provides a mock function with given fields: _a0, _a1
func (_m *Store) UpdateUserDefaultChannel(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
