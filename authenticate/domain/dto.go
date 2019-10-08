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

func ToLoginUserDO(cmd *AddLoginUserCmd) *LoginUserDO {
	return &LoginUserDO{
		Username:   cmd.Username,
		Password:   cmd.Password,
		UniqueCode: cmd.UniqueCode,
		IsLock:     true,
		Mobile:     cmd.mobile,
	}
}

type UserStatus string

const NotExist UserStatus = "Not_Exist"
const ALLOWED UserStatus = "ALLOWED"
const LOCKED UserStatus = "LOCKED"

type AddUserResult string

const Success AddUserResult = "SUCCESS"
const Existed AddUserResult = "EXISTED"
