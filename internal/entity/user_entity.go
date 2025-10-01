package entity

type User struct {
	Base
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Nim      string `gorm:"type:varchar(255);not null" json:"nim"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Role     string `gorm:"type:enum('registrant', 'admin') not null" json:"role"`
}
