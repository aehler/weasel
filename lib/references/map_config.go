package references

import (
	"weasel/app/form"
	"weasel/app/registry"
	"errors"
)

func (r *Reference) mapConfig() error {

	if registry.Registry.ReferenceConf[r.Alias] == nil {

		return errors.New("Config not found!")
	}

	r.Meta = referenceMeta{
		Type : registry.Registry.ReferenceConf[r.Alias].RefType,
		Fields : []*form.Element{},
	}

	for _, f := range registry.Registry.ReferenceConf[r.Alias].Fields {

		el := form.Element{
			Name : f.Name,
			Label : f.Label,
			Order : f.Ord,
			Type : form.MapType(f.Type),
			TypeName : f.Type,
		}

		r.Meta.Fields = append(r.Meta.Fields, &el)

	}

	return nil
}
