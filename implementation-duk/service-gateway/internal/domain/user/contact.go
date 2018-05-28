package user

type Contact struct {
	email string `gorm:"type:varchar(50); not null; unique"`
	phone string `gorm:"type:varchar(50); not null; unique"`
}
