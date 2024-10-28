package chromerob

import (
	"github.com/go-rod/rod"
	"github.com/mizuki1412/go-core-kit/v2/library/jsonkit"
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
	"github.com/spf13/cast"
)

// LogicFort36Login 35,36堡垒机登录
func LogicFort36Login(browser *rod.Browser) {
	page := browser.MustPage(configkit.GetString("fort.url"))

	ok := make(chan int)
	RequestHandleAdd("bhost/_getIcode", "POST", func(res ResponseEntity) {
		data := jsonkit.ParseMap(res.Body)
		page.MustElement("#vdcode").MustInput(cast.ToString(data["info"])).MustWaitLoad()
		ok <- 1
	})
	RequestHandleExec(page)
	page.MustElement("#uCode").MustInput(configkit.GetString("fort.uname")).MustWaitLoad()
	page.MustElement("#pW").MustInput(configkit.GetString("fort.pwd")).MustWaitLoad()
	<-ok
	page.MustElement(".login").MustClick()
	page.MustWaitStable()
}

func LogicFort38Login(browser *rod.Browser) {
	page := browser.MustPage(configkit.GetString("fort.url"))

	page.MustElement("#username").MustInput(configkit.GetString("fort.uname")).MustWaitLoad()
	page.MustElement("#tUserLock").MustInput(configkit.GetString("fort.pwd")).MustWaitLoad()
	page.MustElement("#do_login").MustClick()
	page.MustWaitStable()
}
