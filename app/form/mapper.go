package form

import (
	"weasel/app/crypto"
	"reflect"
	"errors"
	"fmt"
)

//Takes any linear struct, searches for tag `weaselform:"elementType"` and tries to create appropriate elements in Form.Elements
func (f *Form) MapStruct(s interface {}) error {

	st := reflect.TypeOf(s)

	if st.Kind() != reflect.Struct {

		return errors.New(fmt.Sprintf("Form MapStruct recieved %s, but needs Struct", st.Kind().String()))
	}

	for i := 0; i < st.NumField(); i++ {

		field := st.Field(i)

		if field.Tag.Get("weaselform") == "" {

			continue
		}

		e := Element{
			Name : field.Name,
			HashName : crypto.Encrypt(field.Name, f.salt),
			Label : field.Tag.Get("formLabel"),
			Order : uint(i),
			Type : elementType[field.Tag.Get("weaselform")],
			TypeName : field.Tag.Get("weaselform"),
		}

		f.Elements = append(f.Elements, &e)

	}

	return nil
}
