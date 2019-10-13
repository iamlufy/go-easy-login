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
	InitLoginUserRepo(mockRepo)
	RegisterFailHandler(Fail)
	RunSpecs(t, "authenticate Suite")
}

var _ = Describe("service", func() {
	var service LoginUserService
	BeforeSuite(func() {
		service = InitLoginUserService(mockRepo)
	})

	Context("Authenticate", func() {
		var cmd = &LoginCmd{
			Username:         "username",
			EffectiveSeconds: 10,
			Mobile:           "12345678901",
			SourceCode:       "123",
			LoginMode:        "PASSWORD",
			EncryptWay:       "MD5",
		}

		var userDo = LoginUserDO{
			Username:   cmd.Username,
			Password:   ChooseEncrypter("MD5")(cmd.SourceCode),
			IsLock:     false,
			UniqueCode: "code",
			Mobile:     "12345678901",
		}
		Describe("Authenticate", func() {
			Context("login by password", func() {
				BeforeEach(func() {
					mockRepo.On("GetOne", cmd.Username).Return(userDo).Once()
				})

				It("should return true ", func() {
					Expect(service.Authenticate(cmd)).To(BeTrue())
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("login by sms code", func() {
				BeforeEach(func() {
					cmd.LoginMode = "SMS"
					cmd.EncryptWay = ""
					mockRepo.On("GetOne", cmd.Username).Return(userDo).Once()
					mockRepo.On("FindSmsCode", userDo.Mobile).Return(cmd.SourceCode).Once()
				})

				It("should return true ", func() {
					Expect(service.Authenticate(cmd)).To(BeTrue())
					mockRepo.AssertExpectations(tt)
				})
			})
		})

		Describe("can login", func() {
			Context("user does not exist", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(LoginUserDO{}, false).Once()
				})

				It("should return false", func() {
					Expect(service.GetUserStatus(cmd.Username)).To(Equal(NotExist))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user is locked", func() {
				BeforeEach(func() {
					userDo.IsLock = true
					mockRepo.On("FindOne", cmd.Username).Return(userDo, true).Once()
				})

				It("should return false", func() {
					Expect(service.GetUserStatus(cmd.Username)).To(Equal(LOCKED))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user are allowed to login", func() {
				BeforeEach(func() {
					userDo.IsLock = false
					mockRepo.On("FindOne", cmd.Username).Return(userDo, true).Once()
				})

				It("should return true", func() {
					Expect(service.GetUserStatus(cmd.Username)).To(Equal(ALLOWED))
					mockRepo.AssertExpectations(tt)

				})
			})
		})

		Describe("add login user", func() {
			cmd := &AddLoginUserCmd{
				Username:   "username",
				Password:   "password",
				EncryptWay: "MD5",
				UniqueCode: "",
			}
			When(" user had existed", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(LoginUserDO{}, true).Once()
				})

				It("should return AddExistingUser", func() {
					Expect(service.AddUser(cmd)).To(Equal(AddExistingUser))
					mockRepo.AssertExpectations(tt)
				})
			})

			When("user is not existing", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(LoginUserDO{}, false).Once()
					userDo := ToLoginUserDO(cmd)
					userDo.IsLock = false
					userDo.Password = ChooseEncrypter(cmd.EncryptWay)(cmd.Password)

					mockRepo.On("Add", userDo).Return(LoginUserDO{}).Once()
				})

				It("should return AddUserSuccess", func() {
					Expect(service.AddUser(cmd)).To(Equal(AddUserSuccess))
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
					BeforeEach(func() {
						mockRepo.On("FindOne", cmd.Username).
							Return(LoginUserDO{Password: ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword)}, true).Once()

						mockRepo.On(
							"Update",
							LoginUserDO{Password: ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)},
							map[string]interface{}{"password": ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)}).
							Return(LoginUserDO{}).Once()
					})

					It("should return success", func() {
						Expect(service.ReSetPassword(cmd)).To(Equal(ResetPasswordSuccess))
						mockRepo.AssertExpectations(tt)
					})
				})

				When("oldPassword is error", func() {
					BeforeEach(func() {
						mockRepo.On("FindOne", cmd.Username).
							Return(LoginUserDO{Password: ""}, true).Once()
					})

					It("should return success", func() {
						Expect(service.ReSetPassword(cmd)).To(Equal(PasswordError))
						mockRepo.AssertExpectations(tt)
					})
				})

			})
			When("user is nonexistent", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).
						Return(LoginUserDO{}, false).Once()
				})
				It("should return not exist", func() {
					service.ReSetPassword(cmd)
					mockRepo.AssertExpectations(tt)
				})
			})
		})
	})

})
