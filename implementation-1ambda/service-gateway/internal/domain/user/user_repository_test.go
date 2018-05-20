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
	var repo user.Repository

	BeforeEach(func() {
		db = test.GetTestDatabase()
		repo = user.NewRepository(db)
	})

	AfterEach(func() {
	})

	Describe("Create()", func() {
		Context("When creating a new user", func() {
			It("should return nil exception", func() {
				u := &user.User{}
				created, ex := repo.CreateUser(u)

				Expect(ex).To(BeNil())
				Expect(created.ID).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("Delete()", func() {
		Context("When trying to delete non-existing user", func() {
			It("should return not found exception", func() {
				deleted, ex := repo.DeleteUser(0)

				Expect(deleted).To(BeFalse())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})

	Describe("Find()", func() {
		Context("When user dose not exsit", func() {
			It("should return NotFoundException", func() {
				invlaidUserId := uint(0)

				found, ex := repo.FindUserById(invlaidUserId)
				Expect(found).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())

			})
		})
	})
})
