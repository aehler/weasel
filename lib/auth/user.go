package auth

import (
	"weasel/app/registry"
	"fmt"
)

type User struct {
	UserLastName   string `json:"ul" db:"user_lastname"`
	UserFirstName  string `json:"uf" db:"user_firstname"`
	UserMiddleName string `json:"um" db:"user_middlename"`
	UserID         uint   `json:"i" db:"id"`
	IsActive       bool   `json:"a" db:"is_active"`
	Login          string `json:"l" db:"login"`
	Email          string `json:"e" db:"email"`
}

type LoginForm struct {
	Login string `weaselform:"login" formLabel:"Логин"`
	Password string `weaselform:"password" formLabel:"Пароль"`
}

func AuthUser(login, password string) (*User, error) {

	fmt.Println()

	u := User{}

	if err := registry.Registry.Connect.Get(&u, `select user_lastname, user_firstname, user_middlename, id, is_active, login, email from users where login=$1 and password=$2`,
		login,
		password,
	); err != nil {

		return &u, err

	}

	return &u, nil

}

