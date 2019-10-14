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
	SourceCode       string
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

const ResetPasswordSuccess = Success
const UserNotExisting = NotExisting
const PasswordError = "PASSWORD_ERROR"

type AuthenticateResult string

const NotAvailable = "NotAvailable"
const AuthenticateFailed = "AuthenticateFailed"
