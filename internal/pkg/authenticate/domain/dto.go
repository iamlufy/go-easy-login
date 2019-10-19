package domain

type AddLoginUserCmd struct {
	Username   string
	Password   string
	EncryptWay string
	Mobile     string
}

type LoginCmd struct {
	Username         string `json:"username" binding:"required"`
	EffectiveSeconds int    `json:"effectiveSeconds" binding:"required"`
	PassCode         string `json:"passCode" binding:"required"`
	LoginMode        string `json:"loginMode" binding:"required"`
	EncryptWay       string `json:"encryptWay" binding:"required"`
	UniqueCode       string `json:"uniqueCode" binding:"required"`
}

type ResetPasswordCmd struct {
	Username    string `json:"username" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"OldPassword" binding:"required"`
	EncryptWay  string `json:"encryptWay" binding:"required"`
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
