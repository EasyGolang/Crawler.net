package main

import (
	_ "embed"

	"Crawler.net/server/global"
	"Crawler.net/server/global/config"
	"Crawler.net/server/router"
	"Crawler.net/wall"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)

	// 初始化系统参数
	global.Start()
	// 下载 原神 壁纸
	wall.Genshin()

	// 启动 http 监听服务
	router.Start()
}
