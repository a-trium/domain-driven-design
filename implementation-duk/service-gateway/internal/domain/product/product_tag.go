package product

type ProductTag struct {
	ProductId uint `gorm:"primary_key; column:product_id; type:unsigned big int; not null; index"`
	TagId     uint `gorm:"primary_key; column:tag_id; type:unsigned big int; not null; index"`
}

func (ProductTag) TableName() string {
	return "product_tag"
}
