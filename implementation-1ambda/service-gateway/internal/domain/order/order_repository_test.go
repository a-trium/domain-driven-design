package order_test

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/order"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
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

	Describe("TODO", func() {
		Context("TODO", func() {
			It("TODO", func() {
			})
		})
	})
})
