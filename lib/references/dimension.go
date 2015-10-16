package references

import (
	"weasel/app/form"
	"encoding/json"
	"strconv"
	"errors"
)

type Dimensions []*Dimension

type Dimension struct {
	ReferenceAlias string
	ReferenceLabel string
	Value uint
	Label string
	Options form.Options
}

func DimOptions(dest *Dimensions, oid uint) error {

	var als []string

	for _, dim := range *dest {

		als = append(als, dim.ReferenceAlias)

	}

	refs, err := ByAlias(oid, als...)

	if err != nil {

		return err

	}

	for _, dim := range *dest {

		i, err := refs[dim.ReferenceAlias].ItemsList()

		if err != nil {

			return err

		}

		dim.Options = i.Opts()

		dim.ReferenceLabel = refs[dim.ReferenceAlias].Name

	}

	return nil

}

func (d *Dimensions) MapValues(vals map[string]string) {

	for _, dim := range *d {

		for k, vs := range vals {

			v, _ := strconv.ParseUint(vs, 10, 64)

			if k == dim.ReferenceAlias {

				dim.Value = uint(v)

				for _, opt := range dim.Options {

					if opt.Value == uint(v) {

						dim.Label = opt.Label

					}

				}

			}

		}

	}

}

func (d *Dimensions) Scan(src interface {}) error {

	var source []byte

	switch src.(type) {

	case string:

		source = []byte(src.(string))

	case []byte:

		source = src.([]byte)

	default:

		return errors.New("Incompatible type for Dimensions")
	}

	if err := json.Unmarshal(source, &d); err != nil {

		return err
	}

	return nil
}
