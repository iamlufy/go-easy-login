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

	Context("Insert", func() {
		userDO := &LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			TenantCode: tenantCode,
			Mobile:     "23456789011",
		}
		It("should return with ID", func() {
			repo.PsqlTunnel.Insert(userDO)
			Expect(userDO.ID).NotTo(BeNil())
		})
	})

	Describe("GetOne", func() {
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
				repo.PsqlTunnel.Insert(userDO)
			})

			It("should get user by username", func() {
				dbUser := repo.GetOne(userDO.Username)
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

	Describe("UpdateFields", func() {

		userDO := &LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			TenantCode: tenantCode,
			Mobile:     "23456789011",
		}

		BeforeEach(func() {
			repo.PsqlTunnel.Insert(userDO)
		})

		It("should return new user", func() {
			repo.PsqlTunnel.Update(userDO, map[string]interface{}{
				"password": "123",
			})
			Expect(userDO.Password).To(Equal("123"))
		})
	})

	Describe("Update And FindOne", func() {

		userDO := &LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			TenantCode: tenantCode,
			Mobile:     "23456789011",
		}

		BeforeEach(func() {
			repo.PsqlTunnel.Insert(userDO)
		})

		It("should return new user", func() {
			repo.PsqlTunnel.UpdateFields(userDO.Username, tenantCode, map[string]interface{}{
				"password": "123",
			})
			Expect(
				(repo.PsqlTunnel.UpdateFields(
					userDO.Username,
					tenantCode,
					map[string]interface{}{"password": "123"})).Password).
				To(Equal("123"))
			loginUserDO, exist := repo.PsqlTunnel.FindOne(userDO.Username, tenantCode)
			Expect(exist).Should(BeTrue())
			Expect(loginUserDO.Password).Should(Equal("123"))
		})
	})

})
