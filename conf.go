
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var c conf
	conf:=c.getConf()
	fmt.Println(*conf)
	fmt.Println(conf.Root)
}

//profile variables
type conf struct {

	Root string  `yaml:"root"`
	Password string `yaml:"password"`
	Localhost string `yaml:"localhost"`
	//"mysqlServiceHost"//"172.17.0.1"////"172.17.0.1""host.docker.internal"
	Port string `yaml:"port"`
	Database string `yaml:"database"`
	Charset string `yaml:"charset"`
}
func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

