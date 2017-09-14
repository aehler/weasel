package organizations

import (
	"weasel/app/form"
	"weasel/lib/auth"
)

func (o *Organization) Form(u auth.User) (*form.Form, error) {

	f := form.New("organizations", "Организации", u.Login)

	if err := f.MapStruct(o); err != nil {

		return nil, err

	}

	return f, nil
}
