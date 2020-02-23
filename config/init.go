package config

import (
	"fmt"
	"strings"

	"golibrary/config/loader"
)

type GeneratorFunc func() loader.ConfLoader

var confLoaderMap = make(map[string]GeneratorFunc)

const (
	//FileNameDefultStepSep 文件后缀分割符号
	FileNameDefultStepSep = "."
	//YamlConfType yaml文件类型
	YamlConfType = "yml"
	//todo:未支持
	//XMLConfType xml文件类型
	XMLConfType = "xml"
	//TomlConfType toml文件类型
	TomlConfType = "toml"
)

//LoadConfig 加载配置文件
func LoadConfig(file string) (loader.ConfLoader, error) {
	tmpStrArr := strings.Split(file, FileNameDefultStepSep)
	if len(tmpStrArr) < 2 {
		return nil, fmt.Errorf("conf file name:%v  error", file)
	}
	confType := tmpStrArr[len(tmpStrArr)-1]
	if generator, ok := confLoaderMap[confType]; ok {
		l := generator()
		err := l.LoadConfigFromFile(file)
		return l, err
	} else {
		return nil, fmt.Errorf("unsupport %v type conf file",confType)
	}
}

func RegisteConfLoaderGenerator(confType string, generator GeneratorFunc) {
	if _, ok := confLoaderMap[confType]; !ok {
		confLoaderMap[confType] = generator
	}
}

func init() {
	confLoaderMap[YamlConfType] = loader.NewYamlLoader
}
