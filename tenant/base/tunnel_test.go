package base

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	tenant_domain "oneday-infrastructure/tenant/domain"
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

	BeforeEach(func() {
		repo = NewRepo()
		repo.DB = repo.DB.Begin()
	})

	AfterEach(func() {
		repo.DB.Rollback()
	})

	Context("Add", func() {
		tenant := &tenant_domain.TenantDO{
			TenantName: strconv.Itoa(rand.Intn(10000000)),
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
		}

		It("should return with ID ", func() {
			Expect(repo.Add(tenant).ID).NotTo(Equal(0))
		})
	})

	Context("FindByName", func() {
		tenant := &tenant_domain.TenantDO{
			TenantName: strconv.Itoa(rand.Intn(10000000)),
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
		}

		BeforeEach(func() {
			repo.Add(tenant)
		})

		It("should return tenant", func() {
			Expect(repo.FindByName(tenant.TenantName).TenantName).To(Equal(tenant.TenantName))
		})
	})
})
