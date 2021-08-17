package model

func init() {
	Register(&User{})
}

type User struct {
	Auditable
	Email string  `json:"email" gorm:"unique;not null"`
	Name  *string `json:"name"`
}
