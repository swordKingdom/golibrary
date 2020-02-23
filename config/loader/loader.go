package loader

import (
	"os"
	"sync"
)

const (
	//EnvConfBasePath 默认配置文件路径，配置在系统环境变量中
	EnvConfBasePath = "ENV_CONF_FILE"
)

//ConfLoader 配置加载器接口
type ConfLoader interface {
	ConfLoaderOp
	LoadConfigFromFile(file string) error
	LoadConfigFromFileReader(file *os.File) error
	ReLoadConf()
	GetFileName() string
}

//ConfLoaderOp 获取配置对象操作的接口
type ConfLoaderOp interface {
	GetInt(key string, defultValue int) int
	GetBool(key string, defultValue bool) bool
	GetString(key string, defultValue string) string
	GetFloat(key string, defultValue float32) float32
}

//BaseConfLoader 配置文件对象操作的实现对象
type baseConfLoader struct {
	confMap map[interface{}]interface{}
	lock    *sync.Mutex
}

func (b *baseConfLoader) init() {
	b.lock = new(sync.Mutex)
	b.confMap = make(map[interface{}]interface{})
}

//GetInt 获取Int类型的配置参数
func (b *baseConfLoader) GetInt(key string, defultValue int) int {
	if value, ok := b.confMap[key]; ok {
		intValue, ok := value.(int)
		if ok {
			return intValue
		}
		return defultValue
	}
	return defultValue
}

//GetBool 获取bool类型的配置参数
func (b *baseConfLoader) GetBool(key string, defultValue bool) bool {
	if value, ok := b.confMap[key]; ok {
		bValue, ok := value.(bool)
		if ok {
			return bValue
		}
		return defultValue
	}
	return defultValue
}

//GetString 获取string类型的配置参数
func (b *baseConfLoader) GetString(key string, defultValue string) string {
	if value, ok := b.confMap[key]; ok {
		bValue, ok := value.(string)
		if ok {
			return bValue
		}
		return defultValue
	}
	return defultValue
}

//GetFloat 获取float32类型的配置参数
func (b *baseConfLoader) GetFloat(key string, defultValue float32) float32 {
	if value, ok := b.confMap[key]; ok {
		fValue, ok := value.(float32)
		if ok {
			return fValue
		}
		return defultValue
	}
	return defultValue
}
