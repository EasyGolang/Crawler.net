package wall

import (
	"bytes"
	"io"
	"os"
	"strings"

	"Crawler.net/server/global"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gocolly/colly"
)

// 原神 壁纸下载器

// wall 页面地址
var WallUrl = "https://wall.alphacoders.com/by_sub_category.php?id=333944&name=%E5%8E%9F%E7%A5%9E+%E5%A3%81%E7%BA%B8&lang=Chinese"

// 保存目录
var SavePath = "./www/wall"

// 下载页数
var PageSize = 20

var ErrDownload = []string{}

func Genshin() {
	isSavePath := mPath.Exists(SavePath)
	if !isSavePath {
		os.MkdirAll(SavePath, 0o777)
	}
	var imgSrc []string
	for i := 0; i < PageSize; i++ {
		Url := WallUrl + "&quickload=7500+&page=" + mStr.ToStr(i)
		GetImgUrl(Url, func(s string) {
			imgSrc = append(imgSrc, s)
		})
	}
	for _, v := range imgSrc {
		SaveFile(v, func(i int) {
			if i < 0 {
				ErrDownload = append(ErrDownload, v)
			}
		})
	}

	errJsonStr := mJson.ToJson(ErrDownload)

	global.Log.Println(mJson.JsonFormat(errJsonStr))
}

func GetImgUrl(Url string, callBack func(string)) {
	c := colly.NewCollector()
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		class := e.Attr("class")
		isClass := strings.Contains(class, "thumb-desktop")
		if isClass {
			src := e.Attr("src")
			callBack(src)
		}
	})
	c.Visit(Url)
}

func SaveFile(Url string, callBack func(int)) {
	global.Log.Println("新建下载:", Url)
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		nameArr := strings.Split(r.Request.URL.String(), "/")
		if len(nameArr) < 2 {
			return
		}
		fileName := nameArr[len(nameArr)-1]
		SaveFile := SavePath + "/" + fileName
		f, err := os.Create(SaveFile)
		if err != nil {
			global.Log.Println("保存失败:", SaveFile)
		}
		io.Copy(f, bytes.NewReader(r.Body))
		global.Log.Println("保存成功:", SaveFile)
		callBack(1)
	})
	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			global.Log.Println("请求错误:", err)
			callBack(-1)
		}
	})
	c.Visit(Url)
}
