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
		db = test.GetTestDatabase(true)
		repo = user.NewRepository(db)
	})

	AfterEach(func() {
	})

	//Describe("Create()", func() {
	//	Context("When creating a new record", func() {
	//		It("should return nil exception", func() {
	//			u := &user.User{}
	//			record, ex := repo.AddUser(u)
	//
	//			Expect(ex).To(BeNil())
	//			Expect(record.ID).Should(BeNumerically(">", 0))
	//		})
	//	})
	//})
	//
	//Describe("Delete()", func() {
	//	Context("When trying to delete non-existing record", func() {
	//		It("should return not found exception", func() {
	//			record, ex := repo.DeleteUser(0)
	//
	//			Expect(record).To(BeFalse())
	//			Expect(ex).NotTo(BeNil())
	//			Expect(ex.IsNotFoundException()).To(BeTrue())
	//		})
	//	})
	//})
	//
	//Describe("Find()", func() {
	//	Context("When the record dose not exist", func() {
	//		It("should return NotFoundException", func() {
	//			invalidId := uint(0)
	//
	//			record, ex := repo.FindUserById(invalidId)
	//			Expect(record).To(BeNil())
	//			Expect(ex).NotTo(BeNil())
	//			Expect(ex.IsNotFoundException()).To(BeTrue())
	//
	//		})
	//	})
	//})

	Describe("Register()", func() {
		//When("got empty uid or empty password", func() {
		//	It("should return BadRequestException", func() {
		//		u1, ex1 := repo.Register("  ", "password")
		//
		//		Expect(u1).Should(BeNil())
		//		Expect(ex1).ShouldNot(BeNil())
		//		Expect(ex1.IsBadRequestException()).Should(BeTrue())
		//
		//		u2, ex2 := repo.Register("uid", "  ")
		//
		//		Expect(u2).Should(BeNil())
		//		Expect(ex2).ShouldNot(BeNil())
		//		Expect(ex2.IsBadRequestException()).Should(BeTrue())
		//	})
		//})

		When("got valid uid and password", func() {
			It("should create AuthIdentity and User", func() {
				aid, ex := repo.Register("uid", "password")

				Expect(ex).Should(BeNil())
				Expect(aid.ID).Should(BeNumerically(">", 0))

			})
		})
	})
})
