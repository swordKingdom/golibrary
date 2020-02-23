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
	GetMetaData() map[string]interface{}
	GetInt(key string, defultValue int) int
	GetBool(key string, defultValue bool) bool
	GetString(key string, defultValue string) string
	GetFloat(key string, defultValue float32) float32
	GetValue(key string ,defultValue interface{})interface{}
}

//BaseConfLoader 配置文件对象操作的实现对象\
type baseConfLoader struct {
	confMeta map[string]interface{}
	lock    *sync.Mutex
}

func (b *baseConfLoader)GetMetaData() map[string]interface{} {
	tmpMetaData := make(map[string]interface{})
	for k,v :=range b.confMeta {
		tmpMetaData[k] = v
	}
	return tmpMetaData
}

func (b *baseConfLoader) init() {
	b.lock = new(sync.Mutex)
	b.confMeta = make(map[string]interface{})
}

//GetInt 获取Int类型的配置参数
func (b *baseConfLoader) GetInt(key string, defultValue int) int {
	if value, ok := b.confMeta[key]; ok {
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
	if value, ok := b.confMeta[key]; ok {
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
	if value, ok := b.confMeta[key]; ok {
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
	if value, ok := b.confMeta[key]; ok {
		fValue, ok := value.(float32)
		if ok {
			return fValue
		}
		return defultValue
	}
	return defultValue
}

//GetFloat 获取结构体，数组
func (b *baseConfLoader)GetValue(key string ,defultValue interface{})interface{}{
	if value, ok := b.confMeta[key]; ok {
		return value
	}
	return defultValue
}


func analysisStruct (baseStr string,d interface{},setDataFunc func(string,interface{})){
	if m,ok := d.(map[interface{}]interface{});ok {
		for k,v := range m {
			val := v
			name := ""
			if baseStr == ""{
				name = k.(string)
			}else{
				name = baseStr + "." +k.(string)
			}
			setDataFunc(name,val)
			analysisStruct(name,val,setDataFunc)
		}
	}
	return
}