package exception

type ProductNotFound struct {
}

func NewProductNotFound() error {
	return &ProductNotFound{}
}

func (*ProductNotFound) Error() string {
	return "Product not found"
}