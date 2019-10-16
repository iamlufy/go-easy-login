package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/mocks"
	"oneday-infrastructure/tools"
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

	Describe("Add tenant", func() {
		cmd := &AddTenantCmd{TenantName: "TenantName"}
		newTenantCO := TenantCO{TenantCode: "code", TenantName: "TenantName"}
		tenant := Tenant{
			TenantName: "TenantName",
			TenantCode: "code",
		}
		genUniqueCode := func() string { return "code" }

		When("tenantName does not exist", func() {
			BeforeEach(func() {
				mockRepo.On("InsertTenant", tenant).Return(tenant).Once()
				mockRepo.On("FindByName", tenant.TenantName).Return(tenant, false).Once()
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
				mockRepo.On("FindByName", tenant.TenantName).Return(tenant, true).Once()
			})

			It("should return existed", func() {
				_, result := service.Add(cmd, genUniqueCode)
				Expect(result == TenantNameExist).To(BeTrue())
			})
		})
	})

	FDescribe("Add tenantUser", func() {
		cmd := &AddUserCmd{
			Username:   "Username",
			Password:   "Password",
			EncryptWay: "MD5",
			Mobile:     "23456789011",
			TenantCode: "123",
		}
		args := User{
			Username:   "Username",
			Password:   tools.ChooseEncrypter(cmd.EncryptWay)(cmd.Password),
			Mobile:     "23456789011",
			TenantCode: "123",
		}
		BeforeEach(func() {
			mockRepo.On("InsertUser", args).Once()
			mockRepo.On("GetByCode", cmd.TenantCode).Return(Tenant{
				TenantCode: cmd.TenantCode,
			}).Once()
		})

		It("should create successfully", func() {
			service.AddUser(cmd)

			mockRepo.AssertExpectations(tt)

		})

	})

})
