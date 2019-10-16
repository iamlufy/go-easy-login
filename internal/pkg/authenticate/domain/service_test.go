package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/mocks"
	"oneday-infrastructure/tools"
	"testing"
)

var tt *testing.T

var mockRepo *mocks.LoginUserRepo

func TestLogin(t *testing.T) {
	tt = t
	mockRepo = &mocks.LoginUserRepo{}
	mockRepo.Test(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "authenticate Suite")
}

var _ = Describe("service", func() {
	var service LoginUserService
	BeforeSuite(func() {
		service = NewLoginUserService(mockRepo)
	})

	Context("Authenticate", func() {
		var cmd = &LoginCmd{
			Username:         "username",
			EffectiveSeconds: 10,
			PassCode:         "123",
			LoginMode:        "PASSWORD",
			EncryptWay:       "MD5",
		}

		user := LoginUser{
			IsLock: false,
		}
		Describe("Authenticate", func() {
			Context("login by Password", func() {
				BeforeEach(func() {
					user.PassCode = PassCode{
						LoginMode: LoginMode(cmd.LoginMode),
						Password:  Password(tools.ChooseEncrypter(cmd.EncryptWay)(cmd.PassCode))}
					mockRepo.On("FindOne", cmd.Username).Return(user, true).Once()
				})

				It("should return true ", func() {
					token, result := service.Authenticate(cmd)
					Expect(string(result)).To(Equal(Success))
					Expect(token).NotTo(Equal(""))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("login by sms code", func() {
				BeforeEach(func() {
					cmd.LoginMode = "SMS_CODE"
					cmd.EncryptWay = ""
					user.PassCode = PassCode{
						LoginMode: LoginMode(cmd.LoginMode),
						SmsCode:   SmsCode(cmd.PassCode),
					}
					mockRepo.On("FindOne", cmd.Username).Return(user, true).Once()
				})

				It("should return true ", func() {
					token, result := service.Authenticate(cmd)
					Expect(string(result)).To(Equal(Success))
					Expect(token).NotTo(Equal(""))
					mockRepo.AssertExpectations(tt)
				})
			})
		})

		Describe("GetUserStatus", func() {
			Context("user does not exist", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(LoginUser{}, false).Once()
				})

				It("should return false", func() {
					Expect(string(service.GetUserStatus(cmd.Username))).To(Equal(NotExist))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user is locked", func() {
				BeforeEach(func() {
					user.IsLock = true
					mockRepo.On("FindOne", cmd.Username).Return(user, true).Once()
				})

				It("should return false", func() {
					Expect(string(service.GetUserStatus(cmd.Username))).To(Equal(NotAvailable))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user are allowed to login", func() {
				BeforeEach(func() {
					user.IsLock = false
					mockRepo.On("FindOne", cmd.Username).Return(user, true).Once()
				})

				It("should return true", func() {
					Expect(string(service.GetUserStatus(cmd.Username))).To(Equal(ALLOWED))
					mockRepo.AssertExpectations(tt)

				})
			})
		})

		Describe("reset user Password", func() {
			cmd := &ResetPasswordCmd{
				Username:    "username",
				NewPassword: "newPassword",
				OldPassword: "oldPassword",
				EncryptWay:  "MD5",
			}
			When("oldPassword is correct", func() {

				BeforeEach(func() {
					loginUser := LoginUser{
						PassCode: NewPassCode(tools.ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword)),
					}
					mockRepo.On("GetOne", cmd.Username).
						Return(loginUser).Once()

					loginUser = LoginUser{
						PassCode: NewPassCode(tools.ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)),
					}
					mockRepo.On(
						"UpdatePasswordByUsername",
						loginUser).Return(LoginUser{}).Once()
				})

				It("should return success", func() {
					Expect(string(service.ReSetPassword(cmd))).To(Equal(ResetPasswordSuccess))
					mockRepo.AssertExpectations(tt)
				})
			})

			When("oldPassword is error", func() {
				BeforeEach(func() {
					mockRepo.On("GetOne", cmd.Username).
						Return(LoginUser{PassCode: PassCode{
							Password: Password(tools.ChooseEncrypter(cmd.EncryptWay)(""))}}).Once()
				})

				It("should return success", func() {
					Expect(string(service.ReSetPassword(cmd))).To(Equal(PasswordError))
					mockRepo.AssertExpectations(tt)
				})
			})

		})
	})

})
