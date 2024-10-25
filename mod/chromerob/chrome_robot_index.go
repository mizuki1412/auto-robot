package chromerob

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

func Start() {
	browser := rod.New().MustConnect().NoDefaultDevice()

	page := browser.MustPage("https://www.baidu.com/")
	page.MustElement("#kw").MustInput("hello")
	page.MustElement("#su").MustClick()
	page.MustWaitStable()

	utils.Pause()
}
