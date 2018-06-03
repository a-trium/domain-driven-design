package domain

import "time"

type BaseModel struct {
	ID        uint `gorm:"primary_key; type:unsigned big int; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
