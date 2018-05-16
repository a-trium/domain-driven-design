package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Password string `gorm:"column:password;"`
	Email    string `gorm:"column:email;"`
	Phone    string `gorm:"column:phone;"`
	Name     string `gorm:"column:name;"`
	Birthday string `gorm:"column:birthday;"`
	Address  string `gorm:"column:address;"`
}

func (User) TableName() string {
	return "User"
}

type Category struct {
	gorm.Model

	Name        string `gorm:"column:name;"`
	Description string `gorm:"column:description;"`
}

func (Category) TableName() string {
	return "Category"
}

type Image struct {
	gorm.Model

	Name string `gorm:"column:name;"`
	Type string `gorm:"column:type;"`
	Path string `gorm:"column:path;"`
}

func (Image) TableName() string {
	return "Image"
}

type Product struct {
	gorm.Model

	Name   string `gorm:"column:name;"`
	Price  string `gorm:"column:price;"`
	Detail string `gorm:"column:detail;"`

	Category    Category `gorm:"foreignkey:CategoryRef"`
	CategoryRef uint

	Image    Image `gorm:"foreignkey:ImageRef"`
	ImageRef uint
}

func (Product) TableName() string {
	return "Product"
}

type Order struct {
	gorm.Model

	State  string `gorm:"column:state;"`
	Amount uint   `gorm:"column:amount;"`

	ShippingCountry  string `gorm:"column:shipping_country;"`
	ShippingCity     string `gorm:"column:shipping_city;"`
	ShippingState    string `gorm:"column:shipping_state;"`
	ShippingZipCode  string `gorm:"column:shipping_zipcode;"`
	ShippingAddress1 string `gorm:"column:shipping_address1;"`
	ShippingAddress2 string `gorm:"column:shipping_address2;"`
	ShippingMessage  string `gorm:"column:shipping_message;"`

	OrdererName  string `gorm:"column:orderer_name;"`
	OrdererPhone string `gorm:"column:orderer_phone;"`
	OrdererEmail string `gorm:"column:orderer_email;"`

	RecipientName  string `gorm:"column:recipient_name;"`
	RecipientPhone string `gorm:"column:recipient_phone;"`
	RecipientEmail string `gorm:"column:recipient_email;"`

	User    User `gorm:"foreignkey:UserRef"`
	UserRef uint

	OrderDetails []Order `gorm:"foreignkey:OrderRef"`
}

func (Order) TableName() string {
	return "Order"
}

type OrderDetail struct {
	gorm.Model

	Index    uint `gorm:"column:index;"`
	Price    uint `gorm:"column:price;"`
	Quantity uint `gorm:"column:quantity;"`
	Amount   uint `gorm:"column:amount;"`

	Order      Order   `gorm:"foreignkey:OrderRef"`
	OrderRef   uint
	Product    Product `gorm:"foreignkey:ProductRef"`
	ProductRef uint
}

func (OrderDetail) TableName() string {
	return "OrderDetail"
}
