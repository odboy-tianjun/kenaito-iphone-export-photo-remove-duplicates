package core

import (
	"fmt"
	"github.com/redmask-hb/GoSimplePrint/goPrint"
	uuid "github.com/satori/go.uuid"
	"kenaito-media-group-tool/util"
	"kenaito-media-group-tool/vars"
	"os"
	"strings"
)

func DoDeleteLiveVideo(rootDir string) {
	pathSep := string(os.PathSeparator)
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	total := len(vars.GlobalFilePathList)
	fmt.Printf("=== 文件总数: %d \n", total)

	for _, filePath := range vars.GlobalFilePathList {
		fileExt := vars.GlobalFilePath2FileExtMap[filePath]
		if ".HEIC" == fileExt {
			vars.GlobalHEICFilePathNameMap[filePath] = vars.GlobalFilePath2FileNameMap[filePath]
			continue
		}
		if ".MOV" == fileExt {
			vars.GlobalMOVFilePathPathNameMap[filePath] = vars.GlobalFilePath2FileNameMap[filePath]
		}
	}

	var movFilePath []string
	for _, heicName := range vars.GlobalHEICFilePathNameMap {
		for movPath, movName := range vars.GlobalMOVFilePathPathNameMap {
			if heicName == movName {
				println(movPath)
				movFilePath = append(movFilePath, movPath)
			}
		}
	}
	fmt.Printf("=== 发现实况视频文件数量: %d\n", len(movFilePath))

	duplicatesDir := rootDir + pathSep + "duplicates-" + uid
	util.CreateDir(duplicatesDir)

	fmt.Printf("=== 开始移动实况视频文件")
	doMoveFileToDir(movFilePath, duplicatesDir)
	fmt.Printf("=== 实况视频已移动到: %v 中\n", duplicatesDir)

}

// 批量移动文件到目录
func doMoveFileToDir(filePathList []string, videoDirPath string) {
	total := len(filePathList)
	var count = 0
	bar := goPrint.NewBar(100)
	bar.SetNotice("=== 移动文件到目录：")
	bar.SetGraph(">")
	pathSeparator := string(os.PathSeparator)
	for _, videoFilePath := range filePathList {
		moveFileToDir(videoFilePath, videoDirPath+pathSeparator)
		count = count + 1
		bar.PrintBar(util.CalcPercentage(count, total))
	}
	bar.PrintEnd("=== Finish")
}

// 移动文件到目录
func moveFileToDir(sourceFilePath string, targetDirectory string) bool {
	splits := strings.Split(sourceFilePath, string(os.PathSeparator))
	fileName := splits[len(splits)-1]
	targetFilePath := targetDirectory + fileName
	err := os.Rename(sourceFilePath, targetFilePath)
	//fmt.Printf("=== 移动文件, 源: %s, 目标: %s \n", sourceFilePath, targetFilePath)
	return err == nil
}
