package base_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"oneday-infrastructure/authenticate/base"
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/helper"
	"strconv"
	"testing"
	"time"
)

var tt *testing.T

var repo domain.LoginUserRepo
var db *gorm.DB

func TestRepo(t *testing.T) {
	tt = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "repo Suite")
}

var _ = Describe("repoTest", func() {
	rand.Seed(time.Now().UnixNano())

	BeforeEach(func() {
		repo = base.NewRepo(func(name string) *gorm.DB {
			db = helper.GetDb(name)
			db = db.Begin()
			return db
		})
	})
	AfterEach(func() {
		db.Rollback()
	})

	Context("Add", func() {
		userDO := &domain.LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
			Mobile:     "23456789011",
		}
		It("should return with ID", func() {
			Expect(repo.Add(userDO).ID).NotTo(BeNil())
		})
	})

	Describe("GenOne", func() {
		var addUserDO domain.LoginUserDO
		userDO := &domain.LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
			Mobile:     "23456789011",
		}
		Context("when user exits", func() {
			BeforeEach(func() {
				addUserDO = repo.Add(userDO)
			})

			It("should get user by username", func() {
				dbUser := repo.GetOne(addUserDO.Username)
				Expect(dbUser).NotTo(BeNil())
				Expect(dbUser.UniqueCode).NotTo(BeNil())
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

})
