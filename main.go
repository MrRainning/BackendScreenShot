package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/chromedp/chromedp"
)

var chromeContext context.Context

func main() {
	StartBrowser()
	http.HandleFunc("/screenshot", func(rw http.ResponseWriter, r *http.Request) {
		urlStr := r.URL.Query().Get("url")
		fmt.Println("url==", urlStr)
		var buf []byte
		nevigateAction := chromedp.Navigate(urlStr)
		shotAction := chromedp.FullScreenshot(&buf, 80)
		//	closeAction := chromedp.
		chromedp.Run(chromeContext, nevigateAction, shotAction)
		_, err := rw.Write(buf)
		if err != nil {
			fmt.Println("Err:", err)
		}
		rw.WriteHeader(http.StatusOK)
		fmt.Println("this is base64")
		fmt.Println(base64.StdEncoding.EncodeToString(buf))
	})
	fmt.Println("Server start ...")
	http.ListenAndServe(":8080", nil)

}

// 后台开一个浏览器进程
func StartBrowser() context.CancelFunc {
	var cancel context.CancelFunc
	chromeContext, cancel = chromedp.NewContext(context.Background())
	return cancel
}
