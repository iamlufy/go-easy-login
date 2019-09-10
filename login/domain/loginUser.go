package domain

import "errors"

type EncryptWay string

type LoginUserE struct {
	Username     string
	IsLock       bool
	UniqueCode   string
	Mobile       string
	canLoginFunc func() bool
	EncryptWay
}

type LoginUserDO struct {
	Username   string
	Password   string
	TenantId   string
	IsLock     bool
	UniqueCode string
	Mobile     string
}

type LoginHandler interface {
	Before()
	Do() bool
	After()
}

type EncryptHelper interface {
	Encrypt(password string) string
	Decrypt(password string) string
	Match(source, encryptedString string) bool
}

var EncryptMap = make(map[EncryptWay]EncryptHelper)

func (user *LoginUserE) Lock() {
	user.IsLock = true
}

//TODO return specific reason of not being allowed to login
func (user *LoginUserE) CanLogin() bool {
	var can bool
	if user.canLoginFunc != nil {
		can = user.canLoginFunc()
	} else {
		can = !user.IsLock
	}
	return can
}

func (user *LoginUserE) HandlerVerify(handler LoginHandler) string {
	handler.Before()

	if !user.CanLogin() {
		panic("user can not user")
	}

	match := handler.Do()
	if !match {
		panic("do not match")
	}
	handler.After()
	return ""
}

func (user *LoginUserE) DoVerify(sourceCode string, encryptedCode string) (bool, error) {
	if !user.CanLogin() {
		return false, errors.New("can not login")
	}
	match := user.EncryptHelper().Match(sourceCode, encryptedCode)
	return match, nil
}

func (encryptWay EncryptWay) EncryptHelper() EncryptHelper {
	if helper, ok := EncryptMap[encryptWay]; ok {
		return helper
	} else {
		panic("can not find helper")
	}
}

func AddEncryptHelper(encryptWay EncryptWay, helper EncryptHelper) {
	EncryptMap[encryptWay] = helper
}
