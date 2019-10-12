package tenant_domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"oneday-infrastructure/internal/pkg/authenticate/mocks"
	"testing"
)

var tt *testing.T

var mockRepo string

func TestTenant(t *testing.T) {
	tt = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "tenant Suite")

}

var _ = Describe("tenant service", func() {

	Describe("Add tenant", func() {
		var cmd *AddTenantCmd

		newTenantCO := TenantCO{UniqueCode: "code", TenantName: "name"}
		newTenantDO := &TenantDO{UniqueCode: "code", TenantName: "name"}
		When("tenant does not exist", func() {
			BeforeEach(func() {
				cmd = &AddTenantCmd{TenantName: "TenantName"}
				mocks.MockFunc(find, func(string) (TenantDO, bool) { return TenantDO{}, false })
				mocks.MockFunc(add, func(do *TenantDO) TenantDO { return *newTenantDO })
				add(newTenantDO)
			})

			It("should return success", func() {
				co, result := AddTenant(cmd)
				Expect(result).To(Equal(AddSuccess))
				Expect(co).To(Equal(newTenantCO))
			})
		})

		When("tenant had existed", func() {
			BeforeEach(func() {
				cmd = &AddTenantCmd{TenantName: "TenantName"}
				mocks.MockFunc(find, func(string) (TenantDO, bool) { return TenantDO{}, true })
			})

			It("should return existed", func() {
				_, result := AddTenant(cmd)
				Expect(result == TenantExist).To(BeTrue())
			})
		})
	})

})
