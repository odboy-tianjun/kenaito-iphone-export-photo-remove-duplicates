package main

import (
	_ "embed"
	"fmt"
	"kenaito-media-group-tool/core"
	"os"
	"time"
)

func main() {
	fmt.Println("=== 警告：该程序会改变目录结构，将在5秒后开始执行")
	time.Sleep(time.Second * 5)
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("=== 获取当前路径异常", err)
		return
	}
	scanner := core.FileScanner{}
	scanner.DoScan(rootDir)
	core.DoDeleteLiveVideo(rootDir)
	fmt.Println("=== 3s后自动退出")
	time.Sleep(time.Second * 3)
}
