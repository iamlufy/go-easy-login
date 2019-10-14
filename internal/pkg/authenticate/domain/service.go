package domain

type LoginUserService struct {
	LoginUserRepo
}

func InitLoginUserService(userRepo LoginUserRepo) LoginUserService {
	return LoginUserService{userRepo}
}

func (service LoginUserService) Authenticate(cmd *LoginCmd) bool {
	getCode := service.encryptCode(cmd.LoginMode)
	return matcher(cmd.EncryptWay)(cmd.SourceCode, getCode(cmd.Username, cmd.TenantCode))
}

func matcher(encryptWay string) Matcher {
	return ChooseMatcher(encryptWay)
}

func (service LoginUserService) encryptCode(loginMode string) func(string, string) string {
	switch loginMode {
	case "PASSWORD":
		return func(username, tenantCode string) string {
			return service.GetOne(username, tenantCode).Password
		}
	case "SMS":
		return func(username, tenantCode string) string {
			return service.FindSmsCode(service.GetOne(username, tenantCode).Mobile)
		}
	default:
		panic("unknown authenticate way")
	}
}

func (service LoginUserService) GetUserStatus(username, tenantCode string) UserStatus {
	userDO, existed := service.FindOne(username, tenantCode)
	if !existed {
		return NotExist
	}
	if userDO.IsLock {
		return LOCKED
	}
	return ALLOWED
}

func (service LoginUserService) AddUser(cmd *AddLoginUserCmd) AddUserResult {
	if _, exist := service.FindOne(cmd.Username, cmd.TenantCode); exist {
		return AddExistingUser
	}
	loginUserDO := ToLoginUserDO(cmd)
	loginUserDO.Password = ChooseEncrypter(cmd.EncryptWay)(loginUserDO.Password)
	service.Add(loginUserDO)
	return AddUserSuccess
}

func (service LoginUserService) ReSetPassword(cmd *ResetPasswordCmd) ResetPasswordResult {
	user, existed := findUser(service.LoginUserRepo, cmd.Username, cmd.TenantCode)
	if !existed {
		return UserNotExisting
	}
	if ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword) != user.Password {
		return PasswordError
	}
	user.Password = ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)
	updatePassword(service.LoginUserRepo, user)
	return ResetPasswordSuccess
}
