package util

import (
	"os"
	"path/filepath"
)

// CreateDir 创建目录
func CreateDir(dirPath string) bool {
	err := os.Mkdir(dirPath, 0755)
	return err == nil
}

// CalcPercentage 计算percentage相对于total的百分比
func CalcPercentage(percentage int, total int) int {
	return int(float64(percentage) / float64(total) * 100)
}

func GetFileNameWithoutExt(path string) string {
	// 获取文件名（包含后缀）
	filename := filepath.Base(path)
	// 去除后缀
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	return name
}
