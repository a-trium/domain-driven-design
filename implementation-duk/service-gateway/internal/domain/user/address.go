package user

type Address struct {
	Address1 string `gorm:"column:address1; type:varchar(255); not null;"`
	//Address2 string `gorm:"column:address2; type:varchar(255); not null;"`
	//Address3 string `gorm:"column:address3; type:varchar(255); not null;"`
	ZipCode string `gorm:"column:zip_code; type:varchar(255); not null;"`
}
