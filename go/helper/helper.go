package helper

import (
	"gopkg.in/yaml.v3"
)

type Conf struct {
	AccountID string `yaml:"account_id"`
	Region    string `yaml:"region"`
	CIDR      string `yaml:"cidr"`
}

func Config(yamlFile []byte) Conf {
	var c Conf
	yaml.Unmarshal(yamlFile, &c)
	/*****
	DEBUG
	*****/
	//fmt.Printf("Result: %v\n", c)

	return c
}
