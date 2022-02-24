package ini

// go-ini 的简单封装

import (
	"golib"
	"gopkg.in/ini.v1"
)

type ConfigStruct struct {
	File *ini.File
	path string
}

func New() ConfigStruct {
	return ConfigStruct{}
}

// Create 创建一个配置文件
func (i ConfigStruct) Create(filePath string) (ConfigStruct, error) {
	// 创建ini文件,并跳过无法识别的数据行
	cfg := ini.Empty(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	})

	filePathAbs, err1 := golib.GetAbsPath(filePath)
	if err1 != nil {
		return ConfigStruct{}, err1
	}

	err2 := cfg.SaveTo(filePathAbs)
	if err2 != nil {
		return ConfigStruct{}, err2
	}

	var configObject ConfigStruct
	configObject.File = cfg
	configObject.path = filePathAbs

	return configObject, nil
}

// LoadFile 加载一个配置文件
func (i ConfigStruct) LoadFile(filePath string) (ConfigStruct, error) {
	// 加载ini文件,并跳过无法识别的数据行
	cfg, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, filePath)

	if err != nil {
		return ConfigStruct{}, err
	}

	var configObject ConfigStruct
	configObject.File = cfg
	configObject.path = filePath

	return configObject, nil
}

// Save 保存
func (i ConfigStruct) Save() error {
	err := i.File.SaveTo(i.path)
	if err != nil {
		return err
	}
	return nil
}
