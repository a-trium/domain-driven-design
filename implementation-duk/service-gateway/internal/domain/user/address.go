package user

type Address struct {
	Address1 string `gorm:"not null"`
	//Address2 string `gorm:"not null"`
	//Address3 string `gorm:"not null"`
	//Address4 string `gorm:"not null"`
	//Address5 string `gorm:"not null"`
	ZipCode  string `gorm:"not null"`
}
