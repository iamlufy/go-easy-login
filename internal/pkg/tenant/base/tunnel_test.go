package base

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"oneday-infrastructure/tools"
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
	var tenantRepo TenantRepo
	BeforeEach(func() {
		tenantRepo = NewTenantRepo(tools.OpenDB)
		tenantRepo.DB = tenantRepo.Begin()
	})

	AfterEach(func() {
		tenantRepo.DB.Rollback()
	})

	Context("Insert", func() {
		tenant := &TenantDO{
			TenantName: strconv.Itoa(rand.Intn(10000000)),
			TenantCode: strconv.Itoa(rand.Intn(10000000)),
		}

		It("should return with ID ", func() {
			tenantRepo.PsqlTunnel.Insert(tenant)
			Expect(tenant.ID).NotTo(Equal(0))
		})
	})

	Context("FindByName", func() {
		tenant := &TenantDO{
			TenantName: strconv.Itoa(rand.Intn(10000000)),
			TenantCode: strconv.Itoa(rand.Intn(10000000)),
		}

		BeforeEach(func() {
			tenantRepo.PsqlTunnel.Insert(tenant)
		})

		It("should return tenant", func() {
			do, exist := tenantRepo.FindByName(tenant.TenantName)
			Expect(do.TenantName).To(Equal(tenant.TenantName))
			Expect(exist).To(BeTrue())
		})
	})
})
