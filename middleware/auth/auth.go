package auth

import (
	"weasel/app"
	"weasel/app/session"
	"weasel/app/registry"
	"encoding/json"
	"errors"
	"fmt"
	"weasel/middleware/guest"
)

type Auth struct {
	User *User
	SSID string
}

type User struct {
	UserLastName   string `json:"ul" db:"user_lastname"`
	UserFirstName  string `json:"uf" db:"user_firstname"`
	UserMiddleName string `json:"um" db:"user_middlename"`
	OrganizationId uint   `json:"oi" db:"organization_id"`
	UserID         uint   `json:"i" db:"user_id"`
	IsActive       bool   `json:"a" db:"is_active"`
	Login          string `json:"l" db:"user_login"`
	Email          string `json:"e" db:"user_email"`
	IsAdmin        bool   `json:"adm" db:"is_admin"`
	SessionID      string `json:"-" db:"-"`
}

func GetAuthUser(c *app.Context) {

	var sd string

	if err := session.Get(c.Request, &sd, &session.Config{Keys : registry.Registry.SessionKeys}); err != nil {

		fmt.Println(err)

		guest.GuestSettings(c)

		return
	}

	v, err := registry.Registry.Session.Get(sd)

	if err != nil {

		fmt.Println(err)

		guest.GuestSettings(c)

		return
	}

	u := User{}

	if err := json.Unmarshal(v, &u); err != nil {

		fmt.Println(err)

		guest.GuestSettings(c)

		return
	}

	u.SessionID = sd

	c.Set("user", u)

}

func Check(c *app.Context) {

	var sd string

	if err := session.Get(c.Request, &sd, &session.Config{Keys : registry.Registry.SessionKeys}); err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	v, err := registry.Registry.Session.Get(sd)

	if err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	u := User{}

	if err := json.Unmarshal(v, &u); err != nil {

		fmt.Println(err)

		app.Redirect("/login/", c, 302)

		return
	}

	u.SessionID = sd

	c.Set("user", u)
}

func (u *User) Scan(src interface {}) error {

	var source []byte

	switch src.(type) {

	case string:

		source = []byte(src.(string))

	case []byte:

		source = src.([]byte)

	default:

		return errors.New("Incompatible type for auth.User")
	}

	if err := json.Unmarshal(source, &u); err != nil {

		return err
	}

	return nil
}
