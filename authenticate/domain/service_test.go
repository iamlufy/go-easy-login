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
					mockRepo.On("GetOne", cmd.Username).Return(&userDo).Once()
				})

				It("should return true ", func() {
					Expect(domain.Authenticate(cmd)).To(BeTrue())
					mockRepo.AssertExpectations(tt)
				})
			})

			FContext("login by sms code", func() {
				BeforeEach(func() {
					cmd.LoginMode = "SMS"
					cmd.EncryptWay = ""
					mockRepo.On("GetOne", cmd.Username).Return(&userDo).Once()
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
					mockRepo.On("FindOne", cmd.Username).Return(nil, false).Once()
				})

				It("should return false", func() {
					Expect(domain.CanLogin(cmd.Username)).To(Equal(domain.NotExist))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user is locked", func() {
				BeforeEach(func() {
					userDo.IsLock = true
					mockRepo.On("FindOne", cmd.Username).Return(userDo, true).Once()
				})

				It("should return false", func() {
					Expect(domain.CanLogin(cmd.Username)).To(Equal(domain.LOCKED))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("user are allowed to login", func() {
				BeforeEach(func() {
					userDo.IsLock = false
					mockRepo.On("FindOne", cmd.Username).Return(userDo, true).Once()
				})

				It("should return true", func() {
					Expect(domain.CanLogin(cmd.Username)).To(Equal(domain.ALLOWED))
					mockRepo.AssertExpectations(tt)

				})
			})
		})

		Describe("add login user", func() {
			cmd := &domain.AddLoginUserCmd{
				Username:   "username",
				Password:   "password",
				EncryptWay: "",
				TenantId:   "",
				UniqueCode: "",
			}
			Context("when user had existed", func() {
				BeforeEach(func() {
					mockRepo.On("FindOne", cmd.Username).Return(nil, true).Once()
				})

				It("should return Existed", func() {
					Expect(domain.AddUser(cmd, func(username string) bool {
						return true
					})).To(Equal(domain.Existed))
					mockRepo.AssertExpectations(tt)
				})
			})

			Context("where user is not existing", func() {
				It("should return Success", func() {
					Expect(domain.AddUser(cmd, func(username string) bool {
						return false
					})).To(Equal(domain.Success))
					mockRepo.AssertExpectations(tt)
				})

			})
		})
	})

})
