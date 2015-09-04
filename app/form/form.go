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
	Method   string
	skipFields map[string]interface {}
}

func New(t, n, s string) *Form {

	return &Form {
		Action  :  "",
		Name    :  n,
		Title   :  t,
		Elements:  []*Element{},
		Method  :  "POST",
		skipFields : map[string]interface {}{},
	}
}

func (f *Form) ParseForm(reciever interface {}, req *http.Request) error {

	req.ParseForm()

	for _, e := range f.Elements {

		e.Value = append(e.Value, req.PostFormValue(e.HashName))

	}

	return f.unmarshal(reciever)
}

func (f *Form) Json() string {

	m, err := json.Marshal(f)

	if err != nil {

		return "{}"
	}

	return string(m)
}

func (f *Form) Skip(fields ...string) {

	for _, nm := range fields {

		f.skipFields[nm] = "skip"
	}

}
