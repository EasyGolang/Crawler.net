package wall

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gocolly/colly"
)

// 原神 壁纸下载器

// wall 页面地址
var WallUrl = "https://wall.alphacoders.com/by_sub_category.php?id=333944&name=%E5%8E%9F%E7%A5%9E+%E5%A3%81%E7%BA%B8&lang=Chinese"

// 保存目录
var SavePath = "./cache"

// 下载页数
var PageSize = 1

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

	DownLoadImg(imgSrc)
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

func SaveFile(Url string) {
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		nameArr := strings.Split(r.Request.URL.String(), "/")
		if len(nameArr) < 2 {
			return
		}
		fileName := nameArr[len(nameArr)-1]
		SaveFile := SavePath + "/" + fileName

		fmt.Println("下载第一张壁纸:", r.Request.URL.String())

		f, err := os.Create(SaveFile)
		if err != nil {
			fmt.Println("保存失败:" + SaveFile)
		}
		io.Copy(f, bytes.NewReader(r.Body))
		fmt.Println("保存成功:" + SaveFile)
	})

	c.Visit(Url)
}

func DownLoadImg(srcArr []string) {
	for _, v := range srcArr {
		SaveFile(v)
	}
}
