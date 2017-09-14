package registry

import (
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Path struct {
	Templates string
	Static string
	HTTPStatic string `yaml:"HTTPStatic"`
}

func ReadPathConf(config string) *Path {

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/config.yml", config))

	if err != nil {

		log.Fatal(err.Error())
	}

	rr := map[string]*Path{}

	if err := yaml.Unmarshal(data, &rr); err != nil {

		log.Fatal(err.Error())
	}

	return rr["path"]
}
