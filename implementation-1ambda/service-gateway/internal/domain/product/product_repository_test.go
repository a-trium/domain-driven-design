package product_test

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductRepository", func() {

	var db *gorm.DB
	var repo product.Repository

	BeforeEach(func() {
		db = test.GetTestDatabase()
		repo = product.NewRepository(db)
	})

	AfterEach(func() {
	})

	Describe("AddCategory()", func() {
		Context("When got valid category", func() {
			It("should return the created category", func() {
				r := &product.Category{}
				created, ex := repo.AddCategory(r)

				Expect(ex).To(BeNil())
				Expect(created.ID).Should(BeNumerically(">", 0))
			})
		})
	})


})
