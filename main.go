package main

import (
	"bytes"
	"fmt"

	"encoding/binary"

	"encoding/hex"

	"strings"

	"github.com/gocolly/colly/v2"
)

type Result struct {
	Name  string `json:"name"`
	Img   string `json:"img"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

var _YahooResult = make([]*Result, 0)

func main() {
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		// fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.OnHTML(".BaseGridItem__grid___2wuJ7", func(e *colly.HTMLElement) {
		// fmt.Println(e.Attr("Response")) // 抓此Tag中的name屬性 來找出此Tag，再印此Tag中的content屬性
		// fmt.Println(e.ChildText(".BaseGridItem__title___2HWui")) // 取出商品名稱
		// fmt.Println(e.ChildAttr("a", "href"))                    // 取出商品連結
		_img := strings.Split(e.ChildAttr(".SquareFence_wrap_3jTo2 > img.SquareImg_img_2gAcq", "srcset"), " ")
		// fmt.Println(_img[0]) // 取出商品圖片

		_price := e.ChildText(".BaseGridItem__itemInfo___3E5Bx > em")
		if _price == "" {
			_price = e.ChildText(".BaseGridItem__price___31jkj > em")
		}
		// fmt.Println(_price) // 取出商品售價
		// fmt.Printf("===================================\n\n")
		_YahooResult = append(_YahooResult, &Result{
			Name:  e.ChildText(".BaseGridItem__title___2HWui"),
			Img:   _img[0],
			Link:  e.ChildAttr("a", "href"),
			Price: _price,
		})
	})

	c.Visit("https://tw.buy.yahoo.com/search/product?p=iphone14") // Visit 要放最後

	for _, e := range _YahooResult {
		fmt.Println(e.Name, "|", e.Img, "|", e.Link, "|", e.Price)
	}
}

func u2s(form string) (to string, err error) {

	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))

	if err != nil {
		return
	}

	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {

		binary.Read(br, binary.BigEndian, &r)

		to += string(r)
	}
	return
}
