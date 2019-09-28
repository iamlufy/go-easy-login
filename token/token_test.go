package token

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

var tt *testing.T

func TestToken(t *testing.T) {
	tt = t

	RegisterFailHandler(Fail)
	RunSpecs(t, "TestToken")
}

var _ = Describe("token", func() {
	uniqueCode := "123"
	var token string

	Context("Generate Token", func() {
		It("should return token", func() {
			Expect(Generate(uniqueCode, 100)).NotTo(BeNil())
		})
	})
	Context("Verify", func() {
		When("when token is legal", func() {
			Context("when token is valid", func() {
				BeforeEach(func() {
					token = Generate(uniqueCode, 100)
				})
				It("should return true", func() {
					result, e := Verify(token)
					Expect(result.Valid).To(BeTrue())
					Expect(e).To(BeNil())
				})
			})

			Context("when token is expired", func() {
				BeforeEach(func() {
					token = Generate(uniqueCode, -10)
				})
				It("token should be invalid  ", func() {
					result, e := Verify(token)
					Expect(result.Valid).To(BeFalse())
					Expect(e).To(BeNil())

				})
			})

		})
		When("when token is illegal", func() {
			BeforeEach(func() {
				token = "123"
			})
			It("should return error ", func() {
				result, e := Verify(token)
				Expect(result).To(BeNil())
				Expect(e).ToNot(BeNil())
			})

		})

	})

	Context("VerifyAndRefresh", func() {
		When("when token is valid", func() {
			BeforeEach(func() {
				token = Generate(uniqueCode, 1)
			})
			It("should cache token", func() {
				tokenString, result := VerifyAndRefresh(token)

				Expect(getCache(uniqueCode)).To(Equal(token))
				Expect(tokenString).To(Equal(token))
				Expect(result).To(BeTrue())
			})
		})
		When("when token is invalid", func() {
			BeforeEach(func() {
				token = Generate(uniqueCode, 1)
				VerifyAndRefresh(token)
				time.After(2 * 1e9)
			})
			It("should refresh token", func() {
				tokenString, result := VerifyAndRefresh(token)
				Expect(tokenString).To(Equal(token))
				Expect(result).To(BeTrue())
				Expect(getCache(uniqueCode)).To(Equal(tokenString))

			})
		})

	})
})
