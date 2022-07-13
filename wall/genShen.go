package wall

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/EasyGolang/goTools/mStr"
	"github.com/gocolly/colly"
)

// 原神壁纸下载器
// wall 页面地址
var WallUrl = "https://wall.alphacoders.com/by_sub_category.php?id=333944&name=%E5%8E%9F%E7%A5%9E+%E5%A3%81%E7%BA%B8&lang=Chinese"

var SavePath = "./cache/"

func Start() {
	var imgSrc []string

	for i := 1; i < 20; i++ {
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
		SaveFile := SavePath + fileName

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
