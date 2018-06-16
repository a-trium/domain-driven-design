package product

type Repository interface {
	FindById(id int) (*Product, error)
	FindByTagName(tagName string) []Product
}
