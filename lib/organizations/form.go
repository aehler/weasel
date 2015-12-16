package organizations

import (
	"weasel/app/form"
	"weasel/middleware/auth"
)

func (o *Organization) Form(u auth.User) (*form.Form, error) {

	f := form.New("organizations", "Организации", u.SessionID)

	if err := f.MapStruct(o); err != nil {

		return nil, err

	}

	return f, nil
}
