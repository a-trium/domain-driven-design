package user

type Contact struct {
	email string `gorm:"column:email; type:varchar(50); not null; unique; index;"`
	phone string `gorm:"column:phone; type:varchar(50); not null; unique; index;"`
}
