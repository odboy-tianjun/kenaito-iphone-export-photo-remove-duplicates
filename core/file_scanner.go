package core

import (
	"fmt"
	"kenaito-media-group-tool/util"
	"kenaito-media-group-tool/vars"
	"log"
	"os"
	"path"
	"path/filepath"
)

type FileScanner struct {
}

func (FileScanner) DoScan(rootDir string) {
	fmt.Printf("=== 开始扫描文件 \n")
	if err := filepath.Walk(rootDir, visitFile); err != nil {
		log.Fatal(err)
	}
}

// 定义walkFn回调函数visit
func visitFile(currentPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err // 如果有错误，直接返回
	}
	if !info.IsDir() {
		ext := path.Ext(currentPath)
		if ".exe" != ext {
			vars.GlobalFilePathList = append(vars.GlobalFilePathList, currentPath)
			// filename, include ext
			vars.GlobalFilePath2FileNameExtMap[currentPath] = filepath.Base(currentPath)
			// filename, exclude ext
			vars.GlobalFilePath2FileNameMap[currentPath] = util.GetFileNameWithoutExt(currentPath)
			// file ext
			vars.GlobalFilePath2FileExtMap[currentPath] = ext
		}
	}
	return nil
}
