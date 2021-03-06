// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "oneday-infrastructure/internal/pkg/tenant/domain"

	mock "github.com/stretchr/testify/mock"
)

// TenantRepo is an autogenerated mock type for the TenantRepo type
type TenantRepo struct {
	mock.Mock
}

// FindByName provides a mock function with given fields: tenantName
func (_m *TenantRepo) FindByName(tenantName string) (domain.Tenant, bool) {
	ret := _m.Called(tenantName)

	var r0 domain.Tenant
	if rf, ok := ret.Get(0).(func(string) domain.Tenant); ok {
		r0 = rf(tenantName)
	} else {
		r0 = ret.Get(0).(domain.Tenant)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(tenantName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetByCode provides a mock function with given fields: tenantCode
func (_m *TenantRepo) GetByCode(tenantCode string) domain.Tenant {
	ret := _m.Called(tenantCode)

	var r0 domain.Tenant
	if rf, ok := ret.Get(0).(func(string) domain.Tenant); ok {
		r0 = rf(tenantCode)
	} else {
		r0 = ret.Get(0).(domain.Tenant)
	}

	return r0
}

// InsertTenant provides a mock function with given fields: tenant
func (_m *TenantRepo) InsertTenant(tenant domain.Tenant) {
	_m.Called(tenant)
}

// InsertUser provides a mock function with given fields: user
func (_m *TenantRepo) InsertUser(user domain.User) {
	_m.Called(user)
}
