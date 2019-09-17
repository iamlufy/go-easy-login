package base_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"oneday-infrastructure/login/base"
	"oneday-infrastructure/login/base/cache"
	"oneday-infrastructure/login/domain/service"
	"testing"
)

var tt *testing.T

func TestLoginUserE_DoVerify(t *testing.T) {
	tt = t

	RegisterFailHandler(Fail)
	RunSpecs(t, "TestLoginUserE")
}

var _ = Describe("token", func() {
	uniqueCode := "123"
	var tokenService service.TokenService
	BeforeEach(func() {
		tokenService = &base.TokenServiceImpl{}
	})

	Context("generate token", func() {
		It("should return token", func() {
			token := tokenService.Generate(uniqueCode, 3600)
			Expect(cache.Get("token:" + uniqueCode)).To(Equal(token))
		})
	})

	Context("token invalid and refresh token", func() {
		var token string
		BeforeEach(func() {
			token = tokenService.Generate(uniqueCode, -1)
		})

		It("should refresh token", func() {
			result, newToken := tokenService.Verify(token)
			Expect(result).To(BeTrue())
			Expect(newToken).ToNot(Equal(token))
		})
	})
})
