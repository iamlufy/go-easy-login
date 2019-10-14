package domain

import "oneday-infrastructure/internal/pkg/authenticate/facade"

type LoginUserService struct {
	LoginUserRepo
}

var service *LoginUserService

func NewLoginUserService(userRepo LoginUserRepo) LoginUserService {
	if service == nil {
		return LoginUserService{userRepo}
	} else {
		service.LoginUserRepo = userRepo
	}
	return *service
}

func (service LoginUserService) Authenticate(cmd *LoginCmd) (string, AuthenticateResult) {
	user, exist := service.FindOne(cmd.Username)
	if !exist {
		return "", NotExisting
	}
	if !user.isAvailable() {
		return "", NotAvailable
	}
	if !service.compareSourceCode(cmd, user) {
		return "", AuthenticateFailed
	}
	return facade.GenerateToken(cmd.UniqueCode, cmd.EffectiveSeconds), Success

}

func (service LoginUserService) compareSourceCode(cmd *LoginCmd, user LoginUser) bool {
	switch cmd.LoginMode {
	case "PASSWORD":
		return user.comparePassword(cmd.SourceCode, cmd.EncryptWay)
	case "SMS":
		return cmd.SourceCode == service.FindSmsCode(user.Mobile)
	default:
		panic("unknown authenticate way")
	}
}

func (service LoginUserService) GetUserStatus(username string) UserStatus {
	userDO, existed := service.FindOne(username)
	if !existed {
		return NotExist
	}
	if userDO.IsLock {
		return LOCKED
	}
	return ALLOWED
}

func (service LoginUserService) ReSetPassword(cmd *ResetPasswordCmd) ResetPasswordResult {
	user, existed := service.FindOne(cmd.Username)
	if !existed {
		return UserNotExisting
	}
	if ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword) != user.Password {
		return PasswordError
	}
	user.resetPassword(cmd.NewPassword, cmd.EncryptWay)
	service.UpdateByUsername(user)
	return ResetPasswordSuccess
}
func (service LoginUserService) Encrypt(encryptWay, s string) string {
	return ChooseEncrypter(encryptWay)(s)
}
