package domain

func Authenticate(cmd *LoginCmd) bool {
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
			return GetUser(username).Password
		}
	case "SMS":
		return func(username string) string {
			return FindSmsCode(GetUser(username).Mobile)
		}
	default:
		panic("unknown authenticate way")
	}
}

func GetUserStatus(username string) UserStatus {
	userDO, existed := FindUser(username)
	if !existed {
		return NotExist
	}
	if userDO.IsLock {
		return LOCKED
	}
	return ALLOWED
}

func AddUser(cmd *AddLoginUserCmd) AddUserResult {
	if isUserExist(cmd.Username) {
		return AddExistingUser
	}
	loginUserDO := ToLoginUserDO(cmd)
	loginUserDO.Password = ChooseEncrypter(cmd.EncryptWay)(loginUserDO.Password)
	Add(loginUserDO)
	return AddUserSuccess
}

func SetNewPassword(cmd *UpdatePasswordCmd) UpdatePasswordResult {
	user, existed := FindUser(cmd.Username)
	if !existed {
		return UpdateUserNotExisting
	}
	if ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword) != user.Password {
		return PasswordError
	}
	user.Password = ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)
	UpdatePassword(user)
	return UpdatePasswordSuccess
}

func isUserExist(username string) bool {
	_, exist := FindUser(username)
	return exist
}
