package domain

type LoginUserService struct{}

func InitLoginUserService(userRepo LoginUserRepo) LoginUserService {
	InitLoginUserRepo(userRepo)
	return LoginUserService{}
}

func (service LoginUserService) Authenticate(cmd *LoginCmd) bool {
	getCode := encryptCode(cmd.LoginMode)
	return matcher(cmd.EncryptWay)(cmd.SourceCode, getCode(cmd.Username))
}

func matcher(encryptWay string) Matcher {
	return ChooseMatcher(encryptWay)
}

func encryptCode(loginMode string) func(string) string {
	switch loginMode {
	case "PASSWORD":
		return func(username string) string {
			return getUser(username).Password
		}
	case "SMS":
		return func(username string) string {
			return findSmsCode(getUser(username).Mobile)
		}
	default:
		panic("unknown authenticate way")
	}
}

func (service LoginUserService) GetUserStatus(username string) UserStatus {
	userDO, existed := findUser(username)
	if !existed {
		return NotExist
	}
	if userDO.IsLock {
		return LOCKED
	}
	return ALLOWED
}

func (service LoginUserService) AddUser(cmd *AddLoginUserCmd) AddUserResult {
	if isUserExist(cmd.Username) {
		return AddExistingUser
	}
	loginUserDO := ToLoginUserDO(cmd)
	loginUserDO.Password = ChooseEncrypter(cmd.EncryptWay)(loginUserDO.Password)
	add(loginUserDO)
	return AddUserSuccess
}

func (service LoginUserService) ReSetPassword(cmd *ResetPasswordCmd) ResetPasswordResult {
	user, existed := findUser(cmd.Username)
	if !existed {
		return UserNotExisting
	}
	if ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword) != user.Password {
		return PasswordError
	}
	user.Password = ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)
	updatePassword(user)
	return ResetPasswordSuccess
}

func (service LoginUserService) GetUniqueCode(username string) string {
	return getUniqueCode(username)
}

func isUserExist(username string) bool {
	_, exist := findUser(username)
	return exist
}
