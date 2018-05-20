package order_test

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/order"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderRepository", func() {

	var db *gorm.DB
	var repo order.Repository

	BeforeEach(func() {
		db = test.GetTestDatabase(false)
		repo = order.NewRepository(db)
	})

	AfterEach(func() {
	})

	Describe("AddCategory()", func() {
		Context("When got valid record", func() {
			It("should return the created record", func() {
				r := &order.Order{}
				created, ex := repo.AddOrder(r)

				Expect(ex).To(BeNil())
				Expect(created.ID).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("FindOrder()", func() {
		Context("When record does not exist", func() {
			It("should return NotFoundException", func() {
				invalidId := uint(0)
				record, ex := repo.FindOrderById(invalidId)

				Expect(record).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})

	Describe("AddImage()", func() {
		Context("When got valid record", func() {
			It("should return the created record", func() {
				r := &order.OrderDetail{}
				record, ex := repo.AddOrderDetail(r)

				Expect(ex).To(BeNil())
				Expect(record.ID).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("FindOrderDetail()", func() {
		Context("When record does not exist", func() {
			It("should return NotFoundException", func() {
				invalidId := uint(0)
				record, ex := repo.FindOrderDetailById(invalidId)

				Expect(record).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})
})
