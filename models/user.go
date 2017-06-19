package models

type User struct {
	FirstName string `json:"username"`
	LastName  string `json:"lastname"`
}

func NewUser() *User {
	return &User{}
}

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this UserModel) GetUsers() []User {
	users := []User{
		User{FirstName: "Ahmed", LastName: "ZENAKHI"},
		User{FirstName: "Zahia", LastName: "CHALA"},
	}
	return users
}
