package domain

type Tenant struct {
	TenantCode string
	TenantName string
}

type User struct {
	TenantCode string
	Username   string
	Mobile     string
	Password   string
}

func (t *Tenant) generateUser(username, password, mobile string) User {
	return User{
		TenantCode: t.TenantCode,
		Username:   username,
		Mobile:     mobile,
		Password:   password,
	}
}
