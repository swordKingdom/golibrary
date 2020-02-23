package loader

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

//TomlLoader toml配置文件加载对象
type TomlLoader struct {
	baseConfLoader
	confPath string
}

//LoadConfigFromFile 读取配置
//TODO：toml读取
func (t *TomlLoader) LoadConfigFromFile(fileName string) error {
	t.baseConfLoader.init()
	t.baseConfLoader.lock.Lock()
	defer t.baseConfLoader.lock.Unlock()
	if fileName == "" {
		fileName = os.Getenv(EnvConfBasePath)
		if fileName == "" {
			fileName = EnvConfBasePath
		}
	}
	if info, _ := os.Stat(fileName); info == nil {
		return errors.New("load conf error")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := toml.DecodeReader(file, &t.baseConfLoader.confMap); err != nil {
		// handle error
	}
	t.confPath = fileName
	return nil
}

//LoadConfigFromFileReader 通过fileReader对象读取对象
func (t *TomlLoader) LoadConfigFromFileReader(file *os.File) error {
	t.baseConfLoader.lock.Lock()
	defer t.baseConfLoader.lock.Unlock()
	if _, err := toml.DecodeReader(file, &t.baseConfLoader.confMap); err != nil {
		// handle error
		return err
	}
	t.confPath = file.Name()
	return nil
}

func (t *TomlLoader) ReLoadConf() {
	t.LoadConfigFromFile(t.confPath)
}

func (t *TomlLoader) GetFileName() string {
	return t.confPath
}

func NewTomlLoader() ConfLoader {
	return &TomlLoader{}
}
