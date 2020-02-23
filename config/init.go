package configloader

import (
	"errors"
	"fmt"
	"strings"

	"this_is_a_explame/lib/configloader/loader"
)

type generatorFunc func() loader.ConfLoader

var confLoaderMap = make(map[string]generatorFunc)

const (
	//YamlConfType yaml文件类型
	YamlConfType = "yml"
	//XMLConfType xml文件类型
	XMLConfType = "xml"
	//TomlConfType toml文件类型
	TomlConfType = "toml"

	//FileNameDefultStepSep 文件后缀分割符号
	FileNameDefultStepSep = "."
)

//LoadConfig 加载配置文件
func LoadConfig(file string) (loader.ConfLoader, error) {
	tmpStr := strings.Split(file, FileNameDefultStepSep)
	if len(tmpStr) < 2 {
		return nil, fmt.Errorf("conf file name:%v  error", file)
	}
	confType := tmpStr[len(tmpStr)-1]
	if generator, ok := confLoaderMap[confType]; ok {
		l := generator()
		err := l.LoadConfigFromFile(file)
		return l, err
	} else {
		return nil, errors.New("unsupport type conf file")
	}
}

func RegisteConfLoaderGenerator(confType string, f generatorFunc) {
	if _, ok := confLoaderMap[confType]; !ok {
		confLoaderMap[confType] = f
	}
}

func init() {
	confLoaderMap[YamlConfType] = loader.NewYamlLoader
	confLoaderMap[XMLConfType] = loader.NewXMLLoader
	confLoaderMap[TomlConfType] = loader.NewTomlLoader
}
