package docx_builder

import (
	"archive/zip"
	"io/ioutil"
	"bytes"
)

func (builder *builder) extractByte(template []byte) (string, error) {

	r, err:= zip.NewReader(bytes.NewReader(template), int64(len(template)))
	if err != nil {
		return "", err
	}

	content, err := builder.finder(r.File)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (builder *builder) Extract(fileName string) (string, error) {

	return builder.extract(fileName)

}

func (builder *builder) extract(fileName string) (string, error) {

	// Open a zip archive for reading.
	r, err := zip.OpenReader(fileName)
	if err != nil {
		return "", err
	}
	defer r.Close()

	content, err := builder.finder(r.File)
	if err != nil {
		return "", err
	}

	return content, nil
}

//*bytes.Buffer
func (builder *builder) packByte() ([]byte, error) {

	r, err:= zip.NewReader(bytes.NewReader(builder.TemplateArch), int64(len(builder.TemplateArch)))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	defer buf.Reset()

	w := zip.NewWriter(buf)

	for _, file := range r.File {

		var contents []byte

		if "word/document.xml" != file.Name {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}

			contents, err = ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			rc.Close()
		} else {
			contents = []byte(builder.BuiltPage)
		}

		f, err := w.Create(file.Name)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(contents)
		if err != nil {
			return nil, err
		}

	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (builder *builder) finder(r []*zip.File) (string, error) {

	var content []byte

	// Iterate through the files in the archive,
	// searching for "word/document.xml"
	for _, f := range r {
		if "word/document.xml" != f.Name {
			continue;
		}
		rc, err := f.Open()
		if err != nil {
			return "", err
		}

		content, err = ioutil.ReadAll(rc)
		if err != nil {
			return "", err
		}

		rc.Close()
	}

	return string(content), nil
}
