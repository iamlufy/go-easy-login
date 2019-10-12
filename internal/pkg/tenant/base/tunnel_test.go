package base

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"oneday-infrastructure/internal/pkg/tenant/domain"
	"strconv"
	"testing"
	"time"
)

var tt *testing.T

func TestTunnel(t *testing.T) {
	tt = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "tunnel Suite")
}

var _ = Describe("tunnel test", func() {
	rand.Seed(time.Now().UnixNano())
	var DB *gorm.DB
	BeforeEach(func() {
		tenantRepo = NewRepo()
		DB = tenantRepo.Begin()
	})

	AfterEach(func() {
		DB.Rollback()
	})

	Context("Add", func() {
		tenant := &tenant_domain.TenantDO{
			TenantName: strconv.Itoa(rand.Intn(10000000)),
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
		}

		It("should return with ID ", func() {
			Expect(tenantRepo.Add(tenant).ID).NotTo(Equal(0))
		})
	})

	Context("FindByName", func() {
		tenant := &tenant_domain.TenantDO{
			TenantName: strconv.Itoa(rand.Intn(10000000)),
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
		}

		BeforeEach(func() {
			tenantRepo.Add(tenant)
		})

		It("should return tenant", func() {
			Expect(tenantRepo.FindByName(tenant.TenantName).TenantName).To(Equal(tenant.TenantName))
		})
	})
})
