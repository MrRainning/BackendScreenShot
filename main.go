package main

import (
	"BackendScreenShot/utils"
	"context"
	"encoding/base64"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/chromedp/chromedp"
)

// TODO
// IP限流、频控
// 支持指定页面大小和质量
// 超时控制
// 异步api

var chromeContext context.Context

func main() {
	StartBrowser()
	http.HandleFunc("/screenshot", func(rw http.ResponseWriter, r *http.Request) {
		// get request params
		values := r.URL.Query()
		urlStr := values.Get("url")
		widthStr := values.Get("width")
		heightStr := values.Get("height")
		width, _ := strconv.Atoi(widthStr)
		if width == 0 {
			width = 1280
		}
		height, _ := strconv.Atoi(heightStr)
		if height == 0 {
			height = 720
		}

		// result store here
		var buf []byte

		// define action
		// set page size
		viewPortAction := chromedp.EmulateViewport(int64(width), int64(height))
		// open page
		nevigateAction := chromedp.Navigate(urlStr)
		// do screenShot
		shotAction := chromedp.CaptureScreenshot(&buf)

		// do actions
		err := chromedp.Run(chromeContext, viewPortAction, nevigateAction, shotAction)
		if err != nil {
			utils.Log().Error(err)
		}
		// write resp
		_, err = rw.Write(buf)
		if err != nil {
			utils.Log().Error("Err:", err)
		}
		rw.WriteHeader(http.StatusOK)
		utils.Log().Debug("this is base64")
		utils.Log().Debug(base64.StdEncoding.EncodeToString(buf))
	})
	utils.Log().Info("Server start ...")
	http.ListenAndServe(":8080", nil)

}

// 后台开一个浏览器进程
func StartBrowser() context.CancelFunc {
	var cancel context.CancelFunc
	chromeContext, cancel = chromedp.NewContext(context.Background())
	return cancel
}
