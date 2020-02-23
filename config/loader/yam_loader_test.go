package loader

import (
	"fmt"
	"testing"
)

func TestNewYamlLoader(t *testing.T) {
	loader := NewYamlLoader()
	err := loader.LoadConfigFromFile("./conf.yml")
	if err != nil {
		print("load file error")
		return
	}
	intConf := loader.GetInt("database.data",0)
	fmt.Println(intConf)
	str := loader.GetString("stringTest","")
	fmt.Println(str)
}