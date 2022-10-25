package main

import (
	"bytes"
	"fmt"

	"encoding/binary"

	"encoding/hex"

	"strings"

	"github.com/gocolly/colly/v2"
)

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
		fmt.Println(e.Text)
	})

	c.Visit("https://tw.buy.yahoo.com/search/product?p=iphone14") // Visit 要放最後
	// fmt.Println(u2s("\u8a02\u95b1\u65b9\u6848"))
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
