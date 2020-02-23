package config

import (
	"golibrary/config/loader"
)

type GeneratorFunc func() loader.ConfLoader

var confLoaderMap = make(map[string]GeneratorFunc)

func RegisteConfLoaderGenerator(confType string, generator GeneratorFunc) {
	if _, ok := confLoaderMap[confType]; !ok {
		confLoaderMap[confType] = generator
	}
}

func loadDefaultConfig() (loader.ConfLoader,error) {
	loader,err := LoadConfig("./conf.yml")
	if err != nil {
		return nil,err
	}
	return loader,nil
}

func init() {
	confLoaderMap[YamlConfType] = loader.NewYamlLoader
	loadDefaultConfig()
}
