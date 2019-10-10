package domain

type AddLoginUserCmd struct {
	Username   string
	Password   string
	EncryptWay string
	TenantId   string
	UniqueCode string
	mobile     string
}
type LoginCmd struct {
	Username         string
	EffectiveSeconds int
	Mobile           string
	SourceCode       string
	LoginMode        string
	EncryptWay       string
}

type UpdatePasswordCmd struct {
	Username    string
	NewPassword string
	OldPassword string
	EncryptWay  string
}

func ToLoginUserDO(cmd *AddLoginUserCmd) *LoginUserDO {
	return &LoginUserDO{
		Username:   cmd.Username,
		Password:   cmd.Password,
		UniqueCode: cmd.UniqueCode,
		IsLock:     false,
		Mobile:     cmd.mobile,
	}
}

const Success = "SUCCESS"
const Existing = "EXISTING"
const NotExisting = "NOT_EXISTING"

type UserStatus string

const NotExist UserStatus = "Not_Exist"
const ALLOWED UserStatus = "ALLOWED"
const LOCKED UserStatus = "LOCKED"

type AddUserResult string

const AddUserSuccess AddUserResult = Success
const AddExistingUser AddUserResult = Existing

type UpdatePasswordResult string

const UpdatePasswordSuccess UpdatePasswordResult = Success
const UpdateUserNotExisting UpdatePasswordResult = NotExisting
const PasswordError UpdatePasswordResult = "PASSWORD_ERROR"
