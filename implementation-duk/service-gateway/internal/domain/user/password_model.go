package user

type Password struct {
	password string `gorm:"column:password; type:varchar(255); not null;"`
}

func (pw *Password) getPassword() string {
	return pw.password
}

func (pw *Password) changePassword(newPassword string) {
	// TODO : crypto
	pw.password = newPassword
}
