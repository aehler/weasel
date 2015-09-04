package form

import (
	"weasel/app/crypto"
	"reflect"
	"strconv"
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

		if f.skipFields[field.Name] != nil {

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

func (f *Form) unmarshal(s interface {}) error {

	v := reflect.ValueOf(s)

	st := reflect.TypeOf(s)

	if st.Kind() == reflect.Ptr {

		st = st.Elem()

	}

	if st.Kind() != reflect.Struct {

		return errors.New(fmt.Sprintf("Form unmarshal recieved %s, but needs Struct", st.Kind().String()))
	}

	for _, e := range f.Elements {

		switch v.Elem().FieldByName(e.Name).Kind() {

		case reflect.String :
			v.Elem().FieldByName(e.Name).SetString(e.GetValue())

		case reflect.Uint :

			val, err := strconv.ParseUint(e.GetValue(), 10, 64)

			if err != nil {
				val = 0
			}

			v.Elem().FieldByName(e.Name).Set(reflect.ValueOf(uint(val)))

		case reflect.Float64,
			reflect.Float32:

			val, err := strconv.ParseFloat(e.GetValue(), 64)

			if err != nil {
				val = 0
			}

			v.Elem().FieldByName(e.Name).Set(reflect.ValueOf(float64(val)))

		default:

			return errors.New(fmt.Sprintf("Cannot unmarshal form, unknown type %s, element %s", v.Elem().FieldByName(e.Name).Kind(), e.Name))

		}

	}


	return nil
}
