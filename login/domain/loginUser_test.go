package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"oneday-infrastructure/login/domain"
	"oneday-infrastructure/login/mocks"
	"testing"
)

var helper *mocks.EncryptHelper
var tt *testing.T

func TestLoginUserE_DoVerify(t *testing.T) {
	helper = &mocks.EncryptHelper{}
	tt = t

	RegisterFailHandler(Fail)
	RunSpecs(t, "TestLoginUserE")
}

var _ = Describe("verify", func() {
	user := &domain.LoginUserE{
		Username:   "username",
		IsLock:     false,
		UniqueCode: "code",
		Mobile:     "12345678901",
		EncryptWay: "MD5",
	}
	sourceCode := "123"
	encryptCode := "123"


	Context("verify successfully", func() {
		BeforeEach(func() {
			helper.On("Match", sourceCode, encryptCode).Return(true).Once()
			domain.AddEncryptHelper("MD5", helper)
		})
		It("do verify ", func() {
			result, err := user.DoVerify(sourceCode, encryptCode)
			Expect(err).To(BeNil())
			Expect(result).To(BeTrue())
			helper.AssertExpectations(tt)
		})
	})

	Context("verify fail", func() {
		BeforeEach(func() {
			helper.On("Match", sourceCode, encryptCode).Return(false).Once()
			domain.AddEncryptHelper("MD5", helper)
		})

		It("do verify", func() {
			result, err := user.DoVerify(sourceCode, encryptCode)

			Expect(err).To(BeNil())
			Expect(result).To(BeFalse())
			helper.AssertExpectations(tt)

		})
	})

	Context("can not login", func() {
		BeforeEach(func() {
			user.Lock()
		})

		It("do verify", func() {
			result, err := user.DoVerify(sourceCode, encryptCode)

			Expect(err).NotTo(BeNil())
			Expect(result).To(BeFalse())
		})
	})
})

