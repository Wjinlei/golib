package golib

import (
	"bufio"
	"fmt"
	"golib/validator"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	FileCreate = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	FileAppend = os.O_WRONLY | os.O_CREATE | os.O_APPEND
)

// FileWrite 写入字符串到文件
func FileWrite(path string, content string, flag int) error {
	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// MakeDir 创建目录,不创建父级目录
func MakeDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}

// MakeDirParent 创建指定目录的父级目录
func MakeDirParent(dirPath string) error {
	if err := os.MkdirAll(filepath.Dir(dirPath), os.ModePerm); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}

// Delete 删除指定路径的空目录或文件
func Delete(path string) error {
	err := os.Remove(path)
	if !os.IsNotExist(err) {
		return err
	}
	return nil
}

// DeleteAll 删除指定路径的目录或文件,如果是目录,那么包括其子目录
func DeleteAll(path string) error {
	err := os.RemoveAll(path)
	if !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Move 移动/重命名文件或目录
func Move(oldPath string, newPath string) error {
	if validator.Exists(newPath) {
		return fmt.Errorf("newpath already exists")
	}
	return os.Rename(oldPath, newPath)
}

// ReadFile 读取文件内容
func ReadFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadLines 读所有行, ReadLinesOffsetN简单封装
func ReadLines(filePath string, lineFeed string) ([]string, error) {
	return ReadLinesOffsetN(filePath, 0, -1, lineFeed)
}

// ReadLinesOffsetN 读几行, offset表示从第几行开始读0开始, n表示读几行, 返回读取到的行的Slice
func ReadLinesOffsetN(filename string, offset uint, n int, lineFeed string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	var ret []string
	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, lineFeed))
	}
	return ret, nil
}

// Copy 复制文件或目录
func Copy(oldPath string, newPath string) error {
	lstat, err := os.Lstat(oldPath)
	if err != nil {
		return err
	}
	if lstat.IsDir() {
		return dirCopy(oldPath, newPath)
	}
	return fileCopy(oldPath, newPath)
}

// fileCopy 复制文件
func fileCopy(oldPath string, newPath string) error {
	// 先看看是不是软链接
	symLink, err := os.Readlink(oldPath)
	// 如果有错,证明不是
	if err != nil {
		// 读取文件内容并写入新文件
		fileData, err := ioutil.ReadFile(oldPath)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(newPath, fileData, 0644); err != nil {
			return err
		}
	} else {
		// 如果没错,证明是一个链接文件,则尝试创建新的软连接
		if err := os.Symlink(symLink, newPath); err != nil {
			return err
		}
	}
	return nil
}

// dirCopy 复制目录
func dirCopy(oldPath string, newPath string) error {
	// 创建目标目录
	if err := MakeDir(newPath); err != nil {
		return err
	}

	// 打开源目录
	oldDir, err := os.Open(oldPath)
	if err != nil {
		return err
	}
	defer func(oldDir *os.File) {
		_ = oldDir.Close()
	}(oldDir)

	// 读取目录中的文件信息
	fileStats, err := oldDir.Readdir(-1)
	if err != nil {
		return err
	}

	// 处理目录下的内容
	for _, fileStat := range fileStats {
		srcPath := fmt.Sprintf("%s/%s", oldPath, fileStat.Name())
		dstPath := fmt.Sprintf("%s/%s", newPath, fileStat.Name())
		if fileStat.IsDir() {
			// 递归创建子目录
			if err := dirCopy(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// 复制文件
			if err := fileCopy(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// FileDownload 下载文件
func FileDownload(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err := MakeDirParent(path); err != nil {
		return err
	}
	file, err := os.OpenFile(path, FileCreate, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
