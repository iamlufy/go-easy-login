// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "oneday-infrastructure/internal/pkg/authenticate/domain"

	mock "github.com/stretchr/testify/mock"
)

// LoginUserRepo is an autogenerated mock type for the LoginUserRepo type
type LoginUserRepo struct {
	mock.Mock
}

// FindOne provides a mock function with given fields: username
func (_m *LoginUserRepo) FindOne(username string) (domain.LoginUser, bool) {
	ret := _m.Called(username)

	var r0 domain.LoginUser
	if rf, ok := ret.Get(0).(func(string) domain.LoginUser); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.LoginUser)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// FindSmsCode provides a mock function with given fields: mobile
func (_m *LoginUserRepo) FindSmsCode(mobile string) string {
	ret := _m.Called(mobile)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(mobile)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOne provides a mock function with given fields: username
func (_m *LoginUserRepo) GetOne(username string) domain.LoginUser {
	ret := _m.Called(username)

	var r0 domain.LoginUser
	if rf, ok := ret.Get(0).(func(string) domain.LoginUser); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.LoginUser)
	}

	return r0
}

// UpdateByUsername provides a mock function with given fields: user
func (_m *LoginUserRepo) UpdateByUsername(user domain.LoginUser) domain.LoginUser {
	ret := _m.Called(user)

	var r0 domain.LoginUser
	if rf, ok := ret.Get(0).(func(domain.LoginUser) domain.LoginUser); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(domain.LoginUser)
	}

	return r0
}
