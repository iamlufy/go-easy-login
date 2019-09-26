package domain

import "errors"

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

func AddUser(cmd *AddLoginUserCmd) error {
	_, existed := FindUser(cmd.Username)
	if existed {
		return errors.New("user had existed")
	}
	loginUserDO := ToLoginUserDO(cmd)
	loginUserDO.Password = ChooseEncrypter(cmd.EncryptWay)(loginUserDO.Password)
	Add(loginUserDO)
	return nil
}
