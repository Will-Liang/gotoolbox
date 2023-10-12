package fileutils

import (
	"fmt"
	"github.com/Will-Liang/gotoolbox/logutils"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"path"
	"path/filepath"
)

// 检查路径 如果不存此路径下的文件，创建文件
// path 要检查的路径
// flag 文件是1，目录是0
// 如果返回""说明失败
func CheckPath(path string, flag int) string {
	if fileutil.IsExist(path) {
		return path
	}

	//该路径不存在，创建目录或者文件
	if flag == 1 {
		//判断文件的父级目录是否存在
		dir_path := filepath.Dir(path)
		if !fileutil.IsExist(dir_path) {

			if !CreateDir(dir_path) {
				//创建路径失败
				return ""
			}
		}

		if fileutil.CreateFile(path) {
			return path
		} else {
			path = ""
		}

	}
	if flag == 0 {
		if CreateDir(path) {
			return path
		} else {
			path = ""
		}

	}
	return path
}

// 创建目录，返回bool
func CreateDir(path string) bool {
	if fileutil.CreateDir(path) == nil {
		return true
	} else {
		return false
	}
}

// 删除给定目录下的所有文件，不删除目录
func DeleteDirFiles(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return
	}

	// 遍历目录中的每个文件
	for _, file := range files {
		// 构建完整的文件路径
		fullPath := path.Join(dirPath, file.Name())

		err := os.Remove(fullPath)
		if err != nil {
			fmt.Println(logutils.GetLog("ERROR", "Error deleting file: "+err.Error()))
		} else {
			fmt.Println("Deleted file:", fullPath)
		}

	}

}

// 遍历指定目录下的所有文件，返回文件路径和文件名称的切片
func ListFiles(dirPath string) ([]string, []string) {
	var paths []string
	var names []string

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("读取目录失败:", err)
		return paths, names
	}

	for _, file := range files {
		filePath := path.Join(dirPath, file.Name())

		if file.IsDir() {
			// 递归遍历子目录
			subPaths, subNames := ListFiles(filePath)
			paths = append(paths, subPaths...)
			names = append(names, subNames...)
		} else {
			// 添加文件路径和文件名称到切片
			paths = append(paths, filePath)
			names = append(names, file.Name())
		}
	}

	return paths, names
}
