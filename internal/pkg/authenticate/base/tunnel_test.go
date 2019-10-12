package base_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	domain2 "oneday-infrastructure/internal/pkg/authenticate/domain"
	"strconv"
	"testing"
	"time"
)

var tt *testing.T

var repo Repo

func TestRepo(t *testing.T) {
	tt = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "repo Suite")
}

var _ = Describe("repoTest", func() {
	rand.Seed(time.Now().UnixNano())

	BeforeEach(func() {
		repo = NewRepo()
		repo.DB = repo.DB.Begin()
	})
	AfterEach(func() {
		repo.DB.Rollback()
	})

	Context("Add", func() {
		userDO := &domain2.LoginUserDO{
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
		var addUserDO domain2.LoginUserDO
		userDO := &domain2.LoginUserDO{
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

	Describe("update", func() {

		var addUserDO domain2.LoginUserDO
		userDO := &domain2.LoginUserDO{
			Model:      gorm.Model{},
			Username:   strconv.Itoa(rand.Intn(10000000)),
			Password:   strconv.Itoa(rand.Intn(10000000)),
			IsLock:     false,
			UniqueCode: strconv.Itoa(rand.Intn(10000000)),
			Mobile:     "23456789011",
		}

		BeforeEach(func() {
			addUserDO = repo.Add(userDO)
		})

		It("should return new user", func() {
			repo.Update(addUserDO, map[string]interface{}{"password": "123"})
			Expect(repo.GetOne(addUserDO.Username).Password).To(Equal("123"))
		})
	})

})