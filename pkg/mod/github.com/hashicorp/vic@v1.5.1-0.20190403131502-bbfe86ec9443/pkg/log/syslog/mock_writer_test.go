// Copyright 2016-2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package syslog

import mock "github.com/stretchr/testify/mock"

// MockWriter is an autogenerated mock type for the Writer type
type MockWriter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockWriter) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Crit provides a mock function with given fields: _a0
func (_m *MockWriter) Crit(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Debug provides a mock function with given fields: _a0
func (_m *MockWriter) Debug(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Emerg provides a mock function with given fields: _a0
func (_m *MockWriter) Emerg(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Err provides a mock function with given fields: _a0
func (_m *MockWriter) Err(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Info provides a mock function with given fields: _a0
func (_m *MockWriter) Info(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Warning provides a mock function with given fields: _a0
func (_m *MockWriter) Warning(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithPriority provides a mock function with given fields: priority
func (_m *MockWriter) WithPriority(priority Priority) Writer {
	ret := _m.Called(priority)

	var r0 Writer
	if rf, ok := ret.Get(0).(func(Priority) Writer); ok {
		r0 = rf(priority)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Writer)
		}
	}

	return r0
}

// WithTag provides a mock function with given fields: tag
func (_m *MockWriter) WithTag(tag string) Writer {
	ret := _m.Called(tag)

	var r0 Writer
	if rf, ok := ret.Get(0).(func(string) Writer); ok {
		r0 = rf(tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Writer)
		}
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *MockWriter) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ Writer = (*MockWriter)(nil)
