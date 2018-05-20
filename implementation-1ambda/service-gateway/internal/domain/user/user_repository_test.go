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
		db = test.GetTestDatabase(false)
		repo = user.NewRepository(db)
	})

	AfterEach(func() {
	})

	Describe("Create()", func() {
		Context("When creating a new record", func() {
			It("should return nil exception", func() {
				u := &user.User{}
				record, ex := repo.AddUser(u)

				Expect(ex).To(BeNil())
				Expect(record.ID).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("Delete()", func() {
		Context("When trying to delete non-existing record", func() {
			It("should return not found exception", func() {
				record, ex := repo.DeleteUser(0)

				Expect(record).To(BeFalse())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})

	Describe("Find()", func() {
		Context("When the record dose not exist", func() {
			It("should return NotFoundException", func() {
				invalidId := uint(0)

				record, ex := repo.FindUserById(invalidId)
				Expect(record).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())

			})
		})
	})
})
