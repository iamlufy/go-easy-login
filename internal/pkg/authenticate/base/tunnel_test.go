package base_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	. "oneday-infrastructure/internal/pkg/authenticate/base"
	"oneday-infrastructure/tools"
	"strconv"
	"testing"
	"time"
)

var tt *testing.T

func TestRepo(t *testing.T) {
	tt = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "tunnel Suite")
}

var _ = Describe("tunnelTest", func() {
	tenantCode := "tenantCode"
	rand.Seed(time.Now().UnixNano())
	var repo LoginUserRepo
	BeforeEach(func() {
		repo = NewLoginUserRepo(tools.OpenDB, tenantCode)
		repo.DB = repo.DB.Begin()
	})
	AfterEach(func() {
		repo.DB.Rollback()
	})

	Context("Add", func() {
		userDO := &LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			TenantCode: tenantCode,
			Mobile:     "23456789011",
		}
		It("should return with ID", func() {
			Expect(repo.PsqlTunnel.Add(userDO).ID).NotTo(BeNil())
		})
	})

	Describe("GetOne", func() {
		var addUserDO LoginUserDO
		userDO := &LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			TenantCode: tenantCode,
			Mobile:     "23456789011",
		}
		Context("when user exits", func() {
			BeforeEach(func() {
				addUserDO = repo.PsqlTunnel.Add(userDO)
			})

			It("should get user by username", func() {
				dbUser := repo.GetOne(addUserDO.Username)
				Expect(dbUser).NotTo(BeNil())
			})
		})

		Context("when user does not exist", func() {
			It("should panic", func() {
				Expect(func() {
					repo.GetOne(strconv.Itoa(rand.Intn(100000)))
				}).To(Panic())
			})

		})
	})

	Describe("update", func() {

		var addUserDO LoginUserDO
		userDO := &LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			TenantCode: tenantCode,
			Mobile:     "23456789011",
		}

		BeforeEach(func() {
			addUserDO = repo.PsqlTunnel.Add(userDO)
		})

		It("should return new user", func() {
			repo.PsqlTunnel.Update(addUserDO, map[string]interface{}{"password": "123"})
			Expect(repo.PsqlTunnel.GetOne(addUserDO.Username, addUserDO.TenantCode).Password).To(Equal("123"))
		})
	})

})
