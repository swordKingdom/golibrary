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

func init() {
	confLoaderMap[YamlConfType] = loader.NewYamlLoader
}
