package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/authenticate/mocks"
	"testing"
)

var tt *testing.T

var mockRepo *mocks.LoginUserRepo

func TestLogin(t *testing.T) {
	tt = t
	mockRepo = &mocks.LoginUserRepo{}
	mockRepo.Test(t)
	domain.NewRepo(mockRepo)
	RegisterFailHandler(Fail)
	RunSpecs(t, "authenticate Suite")
}

var _ = Describe("service", func() {

	Context("Authenticate", func() {
		var cmd = &domain.LoginCmd{
			Username:         "username",
			EffectiveSeconds: 10,
			Mobile:           "12345678901",
			SourceCode:       "123",
			LoginMode:        "PASSWORD",
			EncryptWay:       "MD5",
		}

		var userDo = domain.LoginUserDO{
			Username:   cmd.Username,
			Password:   domain.ChooseEncrypter("MD5")(cmd.SourceCode),
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
					Expect(domain.Authenticate(cmd)).To(BeTrue())
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
					Expect(domain.Authenticate(cmd)).To(BeTrue())
					mockRepo.AssertExpectations(tt)
				})
			})
		})

		Describe("can login", func() {
			Context("user does not exist", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(domain.LoginUserDO{}, false).Once()
				})

				It("should return false", func() {
					Expect(domain.GetUserStatus(cmd.Username)).To(Equal(domain.NotExist))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user is locked", func() {
				BeforeEach(func() {
					userDo.IsLock = true
					mockRepo.On("FindOne", cmd.Username).Return(userDo, true).Once()
				})

				It("should return false", func() {
					Expect(domain.GetUserStatus(cmd.Username)).To(Equal(domain.LOCKED))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user are allowed to login", func() {
				BeforeEach(func() {
					userDo.IsLock = false
					mockRepo.On("FindOne", cmd.Username).Return(userDo, true).Once()
				})

				It("should return true", func() {
					Expect(domain.GetUserStatus(cmd.Username)).To(Equal(domain.ALLOWED))
					mockRepo.AssertExpectations(tt)

				})
			})
		})

		Describe("add login user", func() {
			cmd := &domain.AddLoginUserCmd{
				Username:   "username",
				Password:   "password",
				EncryptWay: "MD5",
				UniqueCode: "",
			}
			When(" user had existed", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(domain.LoginUserDO{}, true).Once()
				})

				It("should return AddExistingUser", func() {
					Expect(domain.AddUser(cmd)).To(Equal(domain.AddExistingUser))
					mockRepo.AssertExpectations(tt)
				})
			})

			When("user is not existing", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(domain.LoginUserDO{}, false).Once()
					userDo := domain.ToLoginUserDO(cmd)
					userDo.IsLock = false
					userDo.Password = domain.ChooseEncrypter(cmd.EncryptWay)(cmd.Password)

					mockRepo.On("Add", userDo).Return(domain.LoginUserDO{}).Once()
				})

				It("should return AddUserSuccess", func() {
					Expect(domain.AddUser(cmd)).To(Equal(domain.AddUserSuccess))
					mockRepo.AssertExpectations(tt)
				})

			})
		})

		Describe("reset user password", func() {
			cmd := &domain.ResetPasswordCmd{
				Username:    "username",
				NewPassword: "newPassword",
				OldPassword: "oldPassword",
				EncryptWay:  "MD5",
			}
			When("user is existing", func() {
				When("oldPassword is correct", func() {
					BeforeEach(func() {
						mockRepo.On("FindOne", cmd.Username).
							Return(domain.LoginUserDO{Password: domain.ChooseEncrypter(cmd.EncryptWay)(cmd.OldPassword)}, true).Once()

						mockRepo.On(
							"Update",
							domain.LoginUserDO{Password: domain.ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)},
							map[string]interface{}{"password": domain.ChooseEncrypter(cmd.EncryptWay)(cmd.NewPassword)}).
							Return(domain.LoginUserDO{}).Once()
					})

					It("should return success", func() {
						Expect(domain.ReSetPassword(cmd)).To(Equal(domain.ResetPasswordSuccess))
						mockRepo.AssertExpectations(tt)
					})
				})

				When("oldPassword is error", func() {
					BeforeEach(func() {
						mockRepo.On("FindOne", cmd.Username).
							Return(domain.LoginUserDO{Password: ""}, true).Once()
					})

					It("should return success", func() {
						Expect(domain.ReSetPassword(cmd)).To(Equal(domain.PasswordError))
						mockRepo.AssertExpectations(tt)
					})
				})

			})
			When("user is nonexistent", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).
						Return(domain.LoginUserDO{}, false).Once()
				})
				It("should return not exist", func() {
					domain.ReSetPassword(cmd)
					mockRepo.AssertExpectations(tt)
				})
			})
		})
	})

})
