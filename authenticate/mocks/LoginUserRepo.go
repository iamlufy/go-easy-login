// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "oneday-infrastructure/authenticate/domain"

	mock "github.com/stretchr/testify/mock"
)

// LoginUserRepo is an autogenerated mock type for the LoginUserRepo type
type LoginUserRepo struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *LoginUserRepo) Add(_a0 *domain.LoginUserDO) domain.LoginUserDO {
	ret := _m.Called(_a0)

	var r0 domain.LoginUserDO
	if rf, ok := ret.Get(0).(func(*domain.LoginUserDO) domain.LoginUserDO); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.LoginUserDO)
	}

	return r0
}

// FindOne provides a mock function with given fields: username
func (_m *LoginUserRepo) FindOne(username string) (domain.LoginUserDO, bool) {
	ret := _m.Called(username)

	var r0 domain.LoginUserDO
	if rf, ok := ret.Get(0).(func(string) domain.LoginUserDO); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.LoginUserDO)
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
func (_m *LoginUserRepo) GetOne(username string) domain.LoginUserDO {
	ret := _m.Called(username)

	var r0 domain.LoginUserDO
	if rf, ok := ret.Get(0).(func(string) domain.LoginUserDO); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.LoginUserDO)
	}

	return r0
}

// Update provides a mock function with given fields: model, updateFields
func (_m *LoginUserRepo) Update(model domain.LoginUserDO, updateFields map[string]interface{}) domain.LoginUserDO {
	ret := _m.Called(model, updateFields)

	var r0 domain.LoginUserDO
	if rf, ok := ret.Get(0).(func(*domain.LoginUserDO, map[string]interface{}) domain.LoginUserDO); ok {
		r0 = rf(model, updateFields)
	} else {
		r0 = ret.Get(0).(domain.LoginUserDO)
	}

	return r0
}
