package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/mocks"
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
			SourceCode:       "123",
			LoginMode:        "PASSWORD",
			EncryptWay:       "MD5",
			//TenantCode:       "tenantCode",
		}

		user := LoginUser{
			Username: cmd.Username,
			Password: ChooseEncrypter("MD5")(cmd.SourceCode),
			IsLock:   false,
			Mobile:   "12345678901",
		}
		Describe("Authenticate", func() {
			Context("login by password", func() {
				BeforeEach(func() {
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
					cmd.LoginMode = "SMS"
					cmd.EncryptWay = ""
					mockRepo.On("FindOne", cmd.Username).Return(user, true).Once()
					mockRepo.On("FindSmsCode", user.Mobile).Return(cmd.SourceCode).Once()
				})

				It("should return true ", func() {
					token, result := service.Authenticate(cmd)
					Expect(string(result)).To(Equal(Success))
					Expect(token).NotTo(Equal(""))
					mockRepo.AssertExpectations(tt)
				})
			})
		})

		Describe("can login", func() {
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
					Expect(string(service.GetUserStatus(cmd.Username))).To(Equal(LOCKED))
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

		Describe("reset user password", func() {
			cmd := &ResetPasswordCmd{
				Username:    "username",
				NewPassword: "newPassword",
				OldPassword: "oldPassword",
				EncryptWay:  "MD5",
			}
			When("user is existing", func() {
				When("oldPassword is correct", func() {
					loginUserDO := LoginUser{
						Password: ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword),
						Username: cmd.Username,
					}

					BeforeEach(func() {
						mockRepo.On("FindOne", cmd.Username).
							Return(loginUserDO, true).Once()

						loginUserDO.Password = ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)

						mockRepo.On(
							"UpdateByUsername",
							loginUserDO).Return(LoginUser{}).Once()
					})

					It("should return success", func() {
						Expect(string(service.ReSetPassword(cmd))).To(Equal(ResetPasswordSuccess))
						mockRepo.AssertExpectations(tt)
					})
				})

				When("oldPassword is error", func() {
					BeforeEach(func() {
						mockRepo.On("FindOne", cmd.Username).
							Return(LoginUser{Password: ""}, true).Once()
					})

					It("should return success", func() {
						Expect(string(service.ReSetPassword(cmd))).To(Equal(PasswordError))
						mockRepo.AssertExpectations(tt)
					})
				})

			})
			When("user is nonexistent", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).
						Return(LoginUser{Password: ""}, false).Once()
				})
				It("should return not exist", func() {
					service.ReSetPassword(cmd)
					mockRepo.AssertExpectations(tt)
				})
			})
		})
	})

})
