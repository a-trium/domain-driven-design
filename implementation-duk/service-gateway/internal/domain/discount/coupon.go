package discount

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"time"
)

type Coupon struct {
	domain.BaseModel
	Name      string    `gorm:"column:name; type:varchar(20); not null; index;"`
	ExpiredAt time.Time `sql:"index"`
	// Type
	// meta
}

func (Coupon) TableName() string {
	return "coupon"
}

// TODO 전체쿠폰(ex:장바구니, 특정가격...etc) 상품쿠폰(ex:카테고리, 셀러...etc)
