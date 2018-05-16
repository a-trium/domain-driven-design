package user_test

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/user"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserRepository", func() {

	var db *gorm.DB
	var repo user.UserRepository

	BeforeEach(func() {
		db = test.GetTestDatabase()
		repo = user.NewUserRepository(db)
	})

	AfterEach(func() {
	})

	Describe("create()", func() {
		Context("When invalid user", func() {
			It("should return nil exception", func() {
				u := &user.User{}
				ex := repo.Create(u)

				Expect(ex).To(BeNil())
				Expect(u.ID).Should(BeNumerically(">", 0))
			})
		})

	})
})
