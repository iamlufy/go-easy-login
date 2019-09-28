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

func CanLogin(username string) LoginAbleStatus {
	userDO, existed := FindUser(username)
	if !existed {
		return NotExist
	}
	if userDO.IsLock {
		return LOCKED
	}
	return ALLOWED
}

func AddUser(cmd *AddLoginUserCmd, isExisted func(username string) bool) AddUserResult {
	if isExisted(cmd.Username) {
		return Existed
	}
	loginUserDO := ToLoginUserDO(cmd)
	loginUserDO.Password = ChooseEncrypter(cmd.EncryptWay)(loginUserDO.Password)
	Add(loginUserDO)
	return Success
}
