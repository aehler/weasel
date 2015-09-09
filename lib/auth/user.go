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

func AuthUser(login, password string) (*auth.User, error) {

	u := auth.User{}

	if err := registry.Registry.Connect.Get(&u, `select user_lastname, user_firstname, user_middlename, user_id, is_active, user_login, user_email, is_admin, organization_id
	from weasel_auth.users where user_login=$1 and user_password=$2`,
		login,
		password,
	); err != nil {

		time.Sleep(2000 * time.Millisecond)

		return &u, err

	}

	return &u, nil

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
