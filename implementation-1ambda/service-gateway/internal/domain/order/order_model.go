package order

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
)

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

	User    user.User `gorm:"foreignkey:UserRef"`
	UserRef uint

	OrderDetails []Order `gorm:"foreignkey:OrderRef"`
}

func (Order) TableName() string {
	return "Order"
}
