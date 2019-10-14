package domain_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/mocks"
	"testing"
)

var tt *testing.T

var mockRepo = &mocks.TenantRepo{}

func TestTenant(t *testing.T) {
	tt = t

	RegisterFailHandler(Fail)
	RunSpecs(t, "tenant Suite")

}

var _ = Describe("tenant service", func() {
	var service TenantService
	BeforeSuite(func() {
		service = InitTenantService(mockRepo)
	})

	Describe("Insert tenant", func() {
		cmd := &AddTenantCmd{TenantName: "TenantName"}
		newTenantCO := TenantCO{UniqueCode: "code", TenantName: "TenantName"}
		newTenantDO := ToTenantDO(cmd)

		genUniqueCode := func() string { return "code" }
		newTenantDO.UniqueCode = genUniqueCode()

		When("tenant does not exist", func() {
			BeforeEach(func() {
				mockRepo.On("Insert", newTenantDO).Return(*newTenantDO).Once()
				mockRepo.On("FindByName", newTenantDO.TenantName).Return(TenantDO{}, false).Once()
			})

			It("should return success", func() {
				co, result := service.Add(cmd, genUniqueCode)
				Expect(result).To(Equal(AddSuccess))
				Expect(co).To(Equal(newTenantCO))

				mockRepo.AssertExpectations(tt)
			})
		})

		When("tenant had existed", func() {
			BeforeEach(func() {
				cmd = &AddTenantCmd{TenantName: "TenantName"}
				mockRepo.On("FindByName", newTenantDO.TenantName).Return(TenantDO{
					Model: gorm.Model{ID: 1},
				}, true).Once()
			})

			It("should return existed", func() {
				_, result := service.Add(cmd, genUniqueCode)
				Expect(result == TenantExist).To(BeTrue())
			})
		})
	})

})
