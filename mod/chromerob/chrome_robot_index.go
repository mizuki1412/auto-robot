package chromerob

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func Start() {
	// 自动寻找chrome
	path, _ := launcher.LookPath()
	u := launcher.New().
		Headless(false).
		Devtools(false).
		Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	browser.MustIgnoreCertErrors(true)

	page := browser.MustPage("https://www.baidu.com/")
	RequestHandleAdd("https://ug.baidu.com/mcp/pc/pcsearch", "POST", func(e *proto.NetworkResponseReceived) {
		println("res: " + e.Response.HeadersText)
		r, _ := proto.NetworkGetResponseBody{RequestID: e.RequestID}.Call(page)
		println(r.Body)
	})
	RequestHandleExc(page)

	page.MustElement("#kw").MustInput("hello")
	page.MustElement("#su").MustClick()

	page.MustWaitStable()

	defer browser.MustClose()
	select {}
}
