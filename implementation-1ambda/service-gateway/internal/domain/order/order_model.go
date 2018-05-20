package order

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model

	State  string `gorm:"column:state; not null;"`
	Amount uint   `gorm:"column:amount; not null;"`

	ShippingCountry  string `gorm:"column:shipping_country; not null;"`
	ShippingCity     string `gorm:"column:shipping_city; not null;"`
	ShippingState    string `gorm:"column:shipping_state; not null;"`
	ShippingZipCode  string `gorm:"column:shipping_zipcode; not null;"`
	ShippingAddress1 string `gorm:"column:shipping_address1; not null;"`
	ShippingAddress2 string `gorm:"column:shipping_address2; not null;"`
	ShippingMessage  string `gorm:"column:shipping_message; not null;"`

	OrdererName  string `gorm:"column:orderer_name; not null;"`
	OrdererPhone string `gorm:"column:orderer_phone; not null;"`
	OrdererEmail string `gorm:"column:orderer_email; not null;"`

	RecipientName  string `gorm:"column:recipient_name; not null;"`
	RecipientPhone string `gorm:"column:recipient_phone; not null;"`
	RecipientEmail string `gorm:"column:recipient_email; not null;"`

	User    user.User `gorm:"foreignkey:UserRef"`
	UserRef uint

	OrderDetails []Order `gorm:"foreignkey:OrderRef"`
}

func (Order) TableName() string {
	return "Order"
}
