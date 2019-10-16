package domain

type AddLoginUserCmd struct {
	Username   string
	Password   string
	EncryptWay string
	Mobile     string
}

type LoginCmd struct {
	Username         string
	EffectiveSeconds int
	PassCode         string
	LoginMode        string
	EncryptWay       string
	UniqueCode       string
}

type ResetPasswordCmd struct {
	Username    string
	NewPassword string
	OldPassword string
	EncryptWay  string
}

const Success = "SUCCESS"
const Existing = "EXISTING"
const NotExisting = "NOT_EXISTING"

type UserStatus string

const NotExist = "Not_Exist"
const ALLOWED = "ALLOWED"
const LOCKED = "LOCKED"

type AddUserResult string

const AddUserSuccess AddUserResult = Success
const AddExistingUser AddUserResult = Existing

type ResetPasswordResult string

func (r ResetPasswordResult) IsSuccess() bool { return r == Success }

const ResetPasswordSuccess = Success
const PasswordError = "PASSWORD_ERROR"

type AuthenticateResult string

func (r AuthenticateResult) IsSuccess() bool { return r == Success }

const NotAvailable = "NotAvailable"
const AuthenticateFailed = "AuthenticateFailed"
