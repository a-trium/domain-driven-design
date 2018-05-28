package discount

import "github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"

type BaseCoupon struct {
	domain.BaseModel
	Name  string `gorm:"type:varchar(20); not null"`
	Value uint
}

func (BaseCoupon) TableName() string {
	return "coupon"
}

// TODO 전체쿠폰(ex:장바구니, 특정가격...etc) 상품쿠폰(ex:카테고리, 셀러...etc)
