package model

type User struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty" validate:"required,alpha"`
	Email    string  `json:"email,omitempty" validate:"required,email"`
	Phone    string  `json:"phone,omitempty" validate:"required,numeric"`
	Password string  `json:"password,omitempty" validate:"required,len=6"`
	Token    *string `json:"token,omitempty"`
}
type Address struct {
	Id        int    `json:"id,omitempty"`
	UserId    int    `json:"userId,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Address   string `json:"address,omitempty"`
	IsDefault bool   `json:"isDefault"`
}

type Login struct {
	EmailOrPhone string `json:"emailOrPhone,omitempty" validate:"required"`
	Password     string `json:"password,omitempty" validate:"required"`
}
