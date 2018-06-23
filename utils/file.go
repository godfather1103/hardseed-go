package utils

import (
	"os"
	"github.com/Tang-RoseChild/mahonia"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PathMkdir(path string)  error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUTF8StrFromGB2312(info string) string {
	if len(info)>0 {
		dec := mahonia.NewDecoder("gb2312")
		enc := mahonia.NewEncoder("UTF-8")
		return enc.ConvertString(dec.ConvertString(info))
	}else {
		return ""
	}
}