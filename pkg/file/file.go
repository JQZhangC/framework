package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

// GetExt 获取文件后缀名
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if notExist := CheckExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
