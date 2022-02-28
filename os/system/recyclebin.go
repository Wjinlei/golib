package system

import (
	"fmt"
	"golib"
	"golib/conf/ini"
	"path/filepath"
	"strings"
)

// 回收站结构体
type recycleBinStruct struct {
	Path string // 回收站路径
}

// RecycleBinInterface 回收站操作接口
type RecycleBinInterface interface {
	Remove(filePath string) (fileNameInRecycleBin string, err error) // 移除文件到回收站
	Restore(fileNameInRecycleBin string) error                       // 恢复文件
	Delete(fileNameInRecycleBin string) error                        // 彻底删除文件
	Empty() error                                                    // 清空回收站
}

// NewRecycleBin 新建回收站
func NewRecycleBin(location string) (RecycleBinInterface, error) {
	filePathAbs, err := golib.GetAbsPath(location)
	if err != nil {
		return nil, err
	}

	err = golib.MakeDir(fmt.Sprintf("%s/files", filePathAbs))
	if err != nil {
		return nil, err
	}
	err = golib.MakeDir(fmt.Sprintf("%s/info", filePathAbs))
	if err != nil {
		return nil, err
	}

	var bin recycleBinStruct
	bin.Path = filePathAbs
	return bin, nil
}

// Remove 移动文件到回收站
func (bin recycleBinStruct) Remove(filePath string) (fileNameInRecycleBin string, err error) {
	// 获取文件的绝对路径
	filePathAbs, err := golib.GetAbsPath(filePath)
	if err != nil {
		return "", err
	}

	// 随机字符串
	r1 := golib.CreateRandomString(8)
	r2 := golib.CreateRandomString(8)
	r3 := golib.CreateRandomString(8)
	r4 := golib.CreateRandomString(8)
	r5 := golib.CreateRandomString(8)

	// 文件在回收站中的名字
	fileNameInRecycleBin = fmt.Sprintf("%s-%s-%s-%s-%s", r1, r2, r3, r4, r5)
	// 文件信息在回收站中的名字
	infoNameInRecycleBin := fmt.Sprintf("%s-%s-%s-%s-%s.trashinfo", r1, r2, r3, r4, r5)
	// 文件的实际路径
	filePathInRecycleBin := fmt.Sprintf("%s/files/%s", bin.Path, fileNameInRecycleBin)
	// 文件信息的实际路径
	infoPathInRecycleBin := fmt.Sprintf("%s/info/%s", bin.Path, infoNameInRecycleBin)

	// 加载info文件
	configStruct := ini.New()
	configObject, err := configStruct.Create(infoPathInRecycleBin)
	if err != nil {
		return "", err
	}
	configObject.File.Section("Trash Info").Key("Path").SetValue(filePathAbs)
	configObject.File.Section("Trash Info").Key("Name").SetValue(filepath.Base(filePathAbs))
	configObject.File.Section("Trash Info").Key("DeletionDate").SetValue(golib.GetNowTime())

	// 保存
	err = configObject.Save()
	if err != nil {
		return "", err
	}

	// 移动文件到回收站
	err = golib.Move(filePathAbs, filePathInRecycleBin)
	if err != nil {
		return "", err
	}

	return fileNameInRecycleBin, nil
}

// Restore 从回收站恢复文件
func (bin recycleBinStruct) Restore(fileNameInRecycleBin string) error {
	// 文件的实际路径
	filePathInRecycleBin := fmt.Sprintf("%s/files/%s", bin.Path, fileNameInRecycleBin)
	// 文件信息的实际路径
	infoPathInRecycleBin := fmt.Sprintf("%s/info/%s.trashinfo", bin.Path, fileNameInRecycleBin)

	// 加载Info文件
	configStruct := ini.New()
	configObject, err := configStruct.LoadFile(infoPathInRecycleBin)
	if err != nil {
		return err
	}

	// 读取文件真实路径
	realPathObject, err := configObject.File.Section("Trash Info").GetKey("Path")
	if err != nil {
		return err
	}
	realPath := realPathObject.String()
	if strings.TrimSpace(realPath) == "" {
		return fmt.Errorf("path is empty")
	}

	// 还原到真实路径
	err = golib.Move(filePathInRecycleBin, realPath)
	if err != nil {
		return err
	}

	// 删除Info文件
	_ = golib.Delete(infoPathInRecycleBin)

	return nil
}

// Delete 彻底删除文件
func (bin recycleBinStruct) Delete(fileNameInRecycleBin string) error {
	// 文件的实际路径
	filePathInRecycleBin := fmt.Sprintf("%s/files/%s", bin.Path, fileNameInRecycleBin)
	// 文件信息的实际路径
	infoPathInRecycleBin := fmt.Sprintf("%s/info/%s.trashinfo", bin.Path, fileNameInRecycleBin)

	// 删除文件
	err := golib.Delete(filePathInRecycleBin)
	if err != nil {
		return err
	}
	_ = golib.Delete(infoPathInRecycleBin)

	return nil
}

// Empty 清空回收站
func (bin recycleBinStruct) Empty() error {
	filesDir := fmt.Sprintf("%s/files", bin.Path)
	infoDir := fmt.Sprintf("%s/info", bin.Path)

	// 清空回收站目录
	err := golib.DeleteAll(filesDir)
	if err != nil {
		return err
	}

	err = golib.DeleteAll(infoDir)
	if err != nil {
		return err
	}

	// 重建目录
	err = golib.MakeDir(filesDir)
	if err != nil {
		return err
	}

	err = golib.MakeDir(infoDir)
	if err != nil {
		return err
	}

	return nil
}
