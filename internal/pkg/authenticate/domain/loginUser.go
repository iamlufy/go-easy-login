package domain

import (
	"fmt"
	"oneday-infrastructure/tools"
)

type LoginUser struct {
	Username string
	IsLock   bool
	PassCode
}

type PassCode struct {
	LoginMode
	Password
	SmsCode
}

type LoginMode string
type Password string
type SmsCode string

func (passCode *PassCode) resetPassword(encryptWay, oldPassword string, newPassword string) bool {
	if passCode.Password.getComparor(encryptWay)(oldPassword) {
		passCode.doResetPassword(newPassword, encryptWay)
		return true
	}
	return false
}

func (user LoginUser) isAvailable() bool {
	return !user.IsLock
}

func NewPassCode(password string) PassCode {
	return PassCode{Password: Password(password)}
}

func (passCode *PassCode) setLoginMode(loginMode string) {
	passCode.LoginMode = LoginMode(loginMode)
}

func (passCode *PassCode) doResetPassword(newPassword, encryptWay string) {
	passCode.Password = Password(newPassword)
	passCode.encryptPassword(encryptWay)
}

func (passCode *PassCode) encryptPassword(encryptWay string) {
	s := tools.ChooseEncrypter(encryptWay)(string(passCode.Password))
	passCode.Password = Password(s)
}

func (user LoginUser) authenticate(loginMode string, encryptWay, beComparedCode string) AuthenticateResult {
	user.setLoginMode(loginMode)
	if !user.isAvailable() {
		return NotAvailable
	}
	if user.getCodeComparor(encryptWay)(beComparedCode) {
		return Success
	}
	return AuthenticateFailed
}

func (passCode PassCode) getCodeComparor(encryptWay string) func(beComparePassCode string) bool {
	switch string(passCode.LoginMode) {
	case "PASSWORD":
		{
			return passCode.Password.getComparor(encryptWay)
		}
	case "SMS_CODE":
		{
			return passCode.SmsCode.getComparor(encryptWay)
		}
	default:
		panic(fmt.Errorf("can not find LoginMode"))
	}

}

type CodeComparor interface {
	getComparor(encryptWay string) func(beCompareCode string) bool
}

func (password Password) getComparor(encryptWay string) func(beComparePassCode string) bool {
	return func(beComparePassCode string) bool {
		return tools.ChooseMatcher(encryptWay)(beComparePassCode, string(password))
	}
}

func (smsCode SmsCode) getComparor(ignore string) func(beCompareCode string) bool {
	return func(beCompareCode string) bool {
		return string(smsCode) == beCompareCode
	}

}
