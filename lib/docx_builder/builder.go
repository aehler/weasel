package docx_builder

import (
	"github.com/flosch/pongo2"
	"regexp"
	"fmt"
	"encoding/xml"
)

type Body struct {
	Paragraph []string `xml:"p>r>t"`
}

type Document struct {
	XMLName xml.Name `xml:"document"`
	Body    Body     `xml:"body"`
}

func (d *Document) Extract(xmlContent string) error {
	/*
	   Extracts the xml elements into their respective struct fields
	*/
	return xml.Unmarshal([]byte(xmlContent), d)
}


var currentLangs = []string{"ru", "en"}

type builder struct {
	Context      pongo2.Context
	TemplateDir  string
	Template     string
	TemplateFile string
	TemplateSrc  string
	BuildLang    string
	BuiltTpl     *pongo2.Template
	BuiltPage    string
	TemplateArch []byte
	Protocol     []byte
}

func (builder *builder) registerFilters() {
	//log.Print("Регистрация фильтров")
}

func (builder *builder) buildTemplate() error {

	tpl, err := pongo2.FromString(builder.TemplateSrc)

	if err != nil {
		return err
	}

	builder.BuiltTpl = tpl

	return nil
}

func (builder *builder) exec() (string, error) {

	out, err := builder.BuiltTpl.Execute(builder.Context)
	if err != nil {
		return "", err
	}

	builder.BuiltPage = out

	return builder.BuiltPage, nil
}

func (builder *builder) pack() error {

	prot, err := builder.packByte()
	if(err != nil) {
		return err
	}
	builder.Protocol = prot

	return nil

}

func (builder *builder) clearStrangeTags() []string {

	res := []string{}

	errstr := []string{}

	re := regexp.MustCompile(`{{(.+?)}}`)

	rerep := regexp.MustCompile(`{{(<.+?>)+(.+?)(<.+?>)+}}`)

	re3 := regexp.MustCompile(`<|>`)

	res = re.FindAllString(builder.TemplateSrc, -1)

	for _, s := range res {

		rep := (rerep.ReplaceAllString(s, "{{$2}}"))

		if len(rep) != len(s) {

			errstr = append(errstr, rep)
			continue
		}

		fs := re3.FindAllString(s, -1)

		if len(fs) > 0 {

			errstr = append(errstr, rep)
		}
	}

	return errstr
}

func (builder *builder) Run() error {

	builder.registerFilters()

	err := builder.buildTemplate()
	if(err != nil) {
		fmt.Println("Template build error:", err)
		return err
	}

	out, err := builder.exec()
	if(err != nil) {
		fmt.Println("Exec error:", err)
		return err
	}

	builder.BuiltPage = out

	err = builder.pack()
	if(err != nil) {
		fmt.Println("Pack error:", err)
		return err
	}

	return nil

}

func (builder *builder) RunTest() []string {

	return builder.clearStrangeTags()

}
