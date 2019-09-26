package token_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"oneday-infrastructure/authenticate/base/cache"
	. "oneday-infrastructure/token"
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

	Context("generate token", func() {
		It("should return token", func() {
			token := Generate(uniqueCode, 3600)
			Expect(cache.Get("token:" + uniqueCode)).To(Equal(token))
		})
	})

	Context("token invalid and refresh token", func() {
		var token string

		It("should refresh token", func() {
			result, newToken := Verify(token)
			Expect(result).To(BeTrue())
			Expect(newToken).ToNot(Equal(token))
		})
	})
})
