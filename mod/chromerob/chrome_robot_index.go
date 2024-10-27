package chromerob

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/mizuki1412/go-core-kit/v2/cli/configkey"
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
)

func Start() {
	// 自动寻找chrome
	path, _ := launcher.LookPath()
	u := launcher.New().
		Headless(false).
		Devtools(false).
		UserDataDir(configkit.GetString(configkey.ProjectDir) + "/chrome").
		Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	browser.MustIgnoreCertErrors(true)

	if configkit.GetString("mod") == "douyin" {
		LogicDouyin(browser)
	}
	//page := browser.MustPage("https://www.douyin.com/user/self?from_tab_name=main&showTab=favorite_collection")
	//page.MustElement("#kw").MustInput("hello")
	//page.MustElement("#su").MustClick()
	//page.MustWaitStable()

	defer browser.MustClose()
	select {}
}
