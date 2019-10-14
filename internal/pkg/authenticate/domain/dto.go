package domain

type AddLoginUserCmd struct {
	Username   string
	Password   string
	EncryptWay string
	Mobile     string
	TenantCode string
}
type LoginCmd struct {
	Username         string
	EffectiveSeconds int
	Mobile           string
	SourceCode       string
	LoginMode        string
	EncryptWay       string
	UniqueCode       string
	TenantCode       string
}

type ResetPasswordCmd struct {
	Username    string
	NewPassword string
	OldPassword string
	EncryptWay  string
	TenantCode  string
}

func ToLoginUserDO(cmd *AddLoginUserCmd) *LoginUserDO {
	return &LoginUserDO{
		Username: cmd.Username,
		Password: cmd.Password,
		IsLock:   false,
		Mobile:   cmd.Mobile,
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

type ResetPasswordResult string

const ResetPasswordSuccess ResetPasswordResult = Success
const UserNotExisting ResetPasswordResult = NotExisting
const PasswordError ResetPasswordResult = "PASSWORD_ERROR"
