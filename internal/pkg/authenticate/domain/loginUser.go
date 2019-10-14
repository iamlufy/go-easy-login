package domain

type LoginUser struct {
	Username string
	Mobile   string
	Password string
	IsLock   bool
}

type LoginUserRepo interface {
	GetOne(username string) LoginUser
	UpdateByUsername(user LoginUser) LoginUser
	FindOne(username string) (LoginUser, bool)
	FindSmsCode(mobile string) string
}

func (user *LoginUser) resetPassword(newPassword, encryptWay string) {
	user.Password = ChooseEncrypter(encryptWay)(newPassword)
}

func (user LoginUser) comparePassword(password, encryptWay string) bool {
	return user.Password == ChooseEncrypter(encryptWay)(password)
}

func (user LoginUser) isAvailable() bool {
	return !user.IsLock
}
