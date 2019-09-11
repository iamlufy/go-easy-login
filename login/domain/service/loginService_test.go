package service

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"oneday-infrastructure/login/domain"
	"oneday-infrastructure/login/domain/common"
	"oneday-infrastructure/login/mocks"
	"testing"
)

var repo *mocks.LoginUserRepo
var token *mocks.TokenService
var tt *testing.T

func TestLogin(t *testing.T) {

	repo = &mocks.LoginUserRepo{}
	token = &mocks.TokenService{}
	repo.Test(t)
	loginService = NewLoginService(repo, token)
	tt = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "login Suite")
}

var _ = Describe("loginService", func() {
	var (
		loginCmd common.LoginCmd
		userDO   *domain.LoginUserDO
	)

	Describe("login", func() {
		BeforeEach(func() {
			loginCmd = common.LoginCmd{
				Username:         "username",
				TenantId:         "tenantId",
				EffectiveSeconds: 60,
				Mobile:           "12345678901",
				SourceCode:       "code",
				LoginWay:         "PASSWORD",
			}
			userDO = &domain.LoginUserDO{
				Username:   loginCmd.Username,
				TenantId:   loginCmd.TenantId,
				Mobile:     loginCmd.Mobile,
				Password:   "123123",
				UniqueCode: "uniqueCode",
			}
			loginService.TokenService.(*mocks.TokenService).On("Generate",userDO.UniqueCode,mock.Anything ).Return("token")
			loginService.LoginUserRepo.(*mocks.LoginUserRepo).On("GetOne", loginCmd.Username, loginCmd.TenantId).Return(userDO).Once()
		})

		Context("login successfully", func() {
			DoVerifyCalled := false
			BeforeEach(func() {
				mocks.InstanceMethod(&domain.LoginUserE{}, "DoVerify", func(e *domain.LoginUserE, s1, s2 string) (bool, error) { DoVerifyCalled = true; return true, nil })
			})

			It("do login", func() {
				token, err := loginService.Login(loginCmd)
				Expect(token).To(Equal("token"))
				Expect(DoVerifyCalled).To(BeTrue())
				Expect(err).To(BeNil())
				loginService.LoginUserRepo.(*mocks.LoginUserRepo).AssertExpectations(tt)
			})
		})

		Context("login fail", func() {
			BeforeEach(func() {
				mocks.InstanceMethod(&domain.LoginUserE{}, "DoVerify", func(e *domain.LoginUserE, s1, s2 string) (bool, error) { return false, errors.New("login fail") })
			})

			It("do login", func() {
				token, err := loginService.Login(loginCmd)
				Expect(token).To(BeEmpty())
				Expect(err).NotTo(BeNil())
			})
		})

	})

	Describe("get encryptCode", func() {
		code := "code"
		var loginWay string

		It("login way is sms code", func() {
			loginService.LoginUserRepo.(*mocks.LoginUserRepo).On("FindSmsCode", loginCmd.Mobile).Return(code).Once()
			loginWay = "SMS"
			Expect(loginService.encryptCode(loginWay, userDO)).To(Equal(code))
			loginService.LoginUserRepo.(*mocks.LoginUserRepo).AssertExpectations(tt)
		})

		It("login way is password", func() {
			loginWay = "PASSWORD"
			Expect(loginService.encryptCode(loginWay, userDO)).To(Equal(userDO.Password))
		})

	})

})
