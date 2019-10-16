package domain

type LoginUserRepo interface {
	GetOne(username string) LoginUser
	UpdatePasswordByUsername(user LoginUser) LoginUser
	FindOne(username string) (LoginUser, bool)
	FindSmsCode(mobile string) string
}

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
	authenticate := user.authenticate(cmd.LoginMode, cmd.EncryptWay, cmd.PassCode)
	if !authenticate.IsSuccess() {
		return "", authenticate
	}
	//TODO to interface
	return GenerateToken(cmd.UniqueCode, cmd.EffectiveSeconds), Success

}

func (service LoginUserService) GetUserStatus(username string) UserStatus {
	user, existed := service.FindOne(username)
	if !existed {
		return NotExist
	}
	if !user.isAvailable() {
		return NotAvailable
	}
	return ALLOWED
}

func (service LoginUserService) ReSetPassword(cmd *ResetPasswordCmd) ResetPasswordResult {
	user := service.GetOne(cmd.Username)

	if !user.resetPassword(cmd.EncryptWay, cmd.OldPassword, cmd.NewPassword) {
		return PasswordError
	}
	service.UpdatePasswordByUsername(user)
	return ResetPasswordSuccess
}
