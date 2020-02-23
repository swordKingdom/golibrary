package config

import (
	"fmt"
	"golibrary/config/loader"
	"strings"
)
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

var GlobalConfig loader.ConfLoader

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

