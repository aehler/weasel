package auth

import (
	"weasel/app/registry"
	"weasel/app/crypto"
	"weasel/middleware/auth"
	"time"
	"fmt"
)

type RegisterForm struct {
	Login string `weaselform:"login" formLabel:"Email"`
	Email string `weaselform:"login" formLabel:"Email"`
	Password string `weaselform:"password" formLabel:"Пароль"`
	Password2 string `weaselform:"password" formLabel:"Повторите пароль"`
	UserLastName string `weaselform:"text" formLabel:"Фамилия"`
	UserFirstName  string `weaselform:"text" formLabel:"Имя"`
	UserMiddleName string `weaselform:"text" formLabel:"Отчество"`
}

type user struct {
	UserLastName   string `json:"ul" db:"user_lastname"`
	UserFirstName  string `json:"uf" db:"user_firstname"`
	UserMiddleName string `json:"um" db:"user_middlename"`
	OrganizationId uint   `json:"oi" db:"organization_id"`
	UserID         uint   `json:"i" db:"user_id"`
	IsActive       bool   `json:"a" db:"is_active"`
	Login          string `json:"l" db:"user_login"`
	Email          string `json:"e" db:"user_email"`
	IsAdmin        bool   `json:"adm" db:"is_admin"`
}

func AuthUser(login, password string) (*auth.User, error) {

	u := user{}

	if err := registry.Registry.Connect.Get(&u, `select user_lastname, user_firstname, user_middlename, user_id, is_active, user_login, user_email, is_admin, organization_id
	from weasel_auth.users where user_login=$1 and user_password=$2 and is_active = true`,
		login,
		password,
	); err != nil {

		time.Sleep(2000 * time.Millisecond)

		return &auth.User{}, err

	}

	return &auth.User{
		UserLastName : u.UserLastName,
		UserFirstName : u.UserFirstName,
		UserMiddleName : u.UserMiddleName,
		OrganizationId : u.OrganizationId,
		UserID : u.UserID,
		IsActive : u.IsActive,
		Login : u.Login,
		Email : u.Email,
		IsAdmin : u.IsAdmin,
	}, nil

	//return &r, nil

}

func AddUser(r RegisterForm) (uint, error) {

	res := 0

	fmt.Println(r)

	password := crypto.Encrypt(r.Password, "")

	if err := registry.Registry.Connect.Get(&res, `select * from weasel_auth.add_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		1,
		r.UserFirstName,
		r.UserLastName,
		r.UserMiddleName,
		"job_title",
		"",
		r.Login,
		password,
		1,
		1,
	); err != nil {

		return 0, err

	}

	return uint(res), nil

}
