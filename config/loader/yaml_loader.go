package loader

import (
	"bufio"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

//YamlLoader yaml配置文件加载对象
type YamlLoader struct {
	baseConfLoader
	confPath string
}



func (y *YamlLoader) saveConf(metaData interface{}) error {
	setDataFunc := func (name string ,value interface{}){
		y.confMeta[name] = value
	}
	analysisStruct("",metaData,setDataFunc)
	return nil
}

//LoadConfigFromFile 读取配置文件
func (y *YamlLoader) LoadConfigFromFile(fileName string) error {
	y.baseConfLoader.init()
	y.baseConfLoader.lock.Lock()
	defer y.baseConfLoader.lock.Unlock()
	if fileName == "" {
		fileName = os.Getenv(EnvConfBasePath)
		if fileName == "" {
			fileName = EnvConfBasePath
		}
	}
	y.confPath = fileName
	if info, _ := os.Stat(fileName); info == nil {
		return errors.New("load conf error")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return err
	}
	var conf interface{}
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		return err
	}
	return y.saveConf(conf)
}

//LoadConfigFromFileReader 从fileReader对象中读 取配置
func (y *YamlLoader) LoadConfigFromFileReader(file *os.File) error {
	y.baseConfLoader.lock.Lock()
	defer y.baseConfLoader.lock.Unlock()
	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bs, &y.baseConfLoader.confMeta)
	if err != nil {
		return err
	}
	y.confPath = file.Name()
	return nil
}

func (y *YamlLoader) ReLoadConf() {
	y.LoadConfigFromFile(y.confPath)
}

func (y *YamlLoader) GetFileName() string {
	return y.confPath
}

func NewYamlLoader() ConfLoader {
	return &YamlLoader{}
}
