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

	switch configkit.GetString("mod") {
	case "douyin":
		LogicDouyin(browser)
	case "fort36":
		LogicFort36Login(browser)
	case "fort38":
		LogicFort38Login(browser)
	}

	defer browser.MustClose()
	select {}
}
