package form

import (
	"net/http"
	"encoding/json"
)

type Form struct {
	Action   string
	Name     string
	Title    string
	Elements []*Element `json:"e"`
	salt     string
}

func New(t, n, s string) *Form {

	return &Form {
		Action  :  "",
		Name    :  n,
		Title   :  t,
		Elements:  []*Element{},
	}
}

func (f *Form) ParseForm(reciever *interface {}, req http.Request) error {

	return nil

}

func (f *Form) Json() string {

	m, err := json.Marshal(f)

	if err != nil {

		return "{}"
	}

	return string(m)
}
