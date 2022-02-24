package system

import (
	"fmt"
	"golib"
	"golib/config/ini"
	"path/filepath"
)

// 回收站结构体
type recycleBinStruct struct {
	Path string // 回收站路径
}

// RecycleBinInterface 回收站操作接口
type RecycleBinInterface interface {
	Remove(filePath string) error  // 移除文件到回收站
	Restore(fileName string) error // 恢复文件
	Delete(fileName string) error  // 彻底删除文件
	Empty() error                  // 清空回收站
}

// NewRecycleBin 新建回收站
func NewRecycleBin(location string) (RecycleBinInterface, error) {
	filePathAbs, err := golib.GetAbsPath(location)
	if err != nil {
		return nil, err
	}

	if err := golib.MakeDir(fmt.Sprintf("%s/files", filePathAbs)); err != nil {
		return nil, err
	}
	if err := golib.MakeDir(fmt.Sprintf("%s/info", filePathAbs)); err != nil {
		return nil, err
	}

	var bin recycleBinStruct
	bin.Path = filePathAbs
	return bin, nil
}

func (bin recycleBinStruct) Remove(filePath string) error {
	// 获取文件的绝对路径
	filePathAbs, err := golib.GetAbsPath(filePath)
	if err != nil {
		return err
	}

	// 随机字符串
	r1 := golib.CreateRandomString(8)
	r2 := golib.CreateRandomString(8)
	r3 := golib.CreateRandomString(8)
	r4 := golib.CreateRandomString(8)
	r5 := golib.CreateRandomString(8)

	// 文件在回收站中的名字
	fileNameInRecycleBin := fmt.Sprintf("%s-%s-%s-%s-%s", r1, r2, r3, r4, r5)
	// 文件信息在回收站中的名字
	infoNameInRecycleBin := fmt.Sprintf("%s-%s-%s-%s-%s.trashinfo", r1, r2, r3, r4, r5)
	// 文件的实际路径
	filePathInRecycleBin := fmt.Sprintf("%s/files/%s", bin.Path, fileNameInRecycleBin)
	// 文件信息的实际路径
	infoPathInRecycleBin := fmt.Sprintf("%s/info/%s", bin.Path, infoNameInRecycleBin)

	// 移动文件到回收站
	if err := golib.Move(filePathAbs, filePathInRecycleBin); err != nil {
		return err
	}

	// 加载info文件
	configStruct := ini.New()
	configObject, err := configStruct.Create(infoPathInRecycleBin)
	if err != nil {
		return err
	}
	configObject.File.Section("Trash Info").Key("Path").SetValue(filePathAbs)
	configObject.File.Section("Trash Info").Key("Name").SetValue(filepath.Base(filePathAbs))
	configObject.File.Section("Trash Info").Key("DeletionDate").SetValue(golib.GetNowTime())

	// 保存
	if err := configObject.Save(); err != nil {
		return err
	}

	return nil
}

func (bin recycleBinStruct) Restore(fileName string) error {
	return nil
}

func (bin recycleBinStruct) Delete(fileName string) error {
	return nil
}

func (bin recycleBinStruct) Empty() error {
	return nil
}
