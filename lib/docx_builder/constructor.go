package docx_builder

import (
	"encoding/json"
	"github.com/flosch/pongo2"
)

func Simple() (*builder, error) {


	builder := &builder{}

	return builder, nil

}

func NewByte(jsonContext pongo2.Context, template []byte) (*builder, error) {

	builder := &builder{
		Context:     jsonContext,
	}

	if content, err := builder.extractByte(template); err == nil {
		builder.TemplateSrc = content
	} else {
		return builder, err
	}

	builder.TemplateArch = template

	return builder, nil

}

func New(jsonContext []byte, templateDir, template string) (*builder, error) {

	templateFile := templateDir + template

	var (
		context   pongo2.Context
	)

	if err := json.Unmarshal(jsonContext, &context); err == nil {
	} else {
		return nil, err
	}

	builder := &builder{
		Context:     context,
		TemplateDir: templateDir,
		Template: template,
		TemplateFile: templateFile,
	}

	if content, err := builder.extract(builder.TemplateFile); err == nil {
		builder.TemplateSrc = content
	} else {
		return builder, err
	}

	return builder, nil
}
