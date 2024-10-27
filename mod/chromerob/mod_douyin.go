package chromerob

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/mizuki1412/go-core-kit/v2/cli/configkey"
	"github.com/mizuki1412/go-core-kit/v2/library/filekit"
	"github.com/mizuki1412/go-core-kit/v2/library/jsonkit"
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
	"github.com/mizuki1412/go-core-kit/v2/service/logkit"
	"github.com/spf13/cast"
	"strings"
)

// LogicDouyin 实现收藏夹列表生成html，需要手动点入下载
func LogicDouyin(browser *rod.Browser) {
	page := browser.MustPage("https://www.douyin.com/user/self?from_tab_name=main&showTab=favorite_collection")
	dist := configkit.GetString(configkey.ProjectDir) + "/douyin.html"
	if filekit.Exists(dist) {
		filekit.RemoveFile(dist)
	}
	RequestHandleAdd("https://www.douyin.com/aweme/v1/web/aweme/listcollection", "POST", func(res ResponseEntity) {
		data := jsonkit.ParseMap(res.Body)
		if len(data) > 0 {
			list := cast.ToSlice(data["aweme_list"])
			for _, l := range list {
				lv := cast.ToStringMap(l)
				//title := cast.ToString(lv["desc"])
				video := cast.ToStringMap(lv["video"])
				addr := cast.ToStringMap(video["play_addr"])
				key := cast.ToString(addr["uri"])
				var url string
				for _, a := range cast.ToStringSlice(addr["url_list"]) {
					if strings.Index(a, "https://v3-web") >= 0 {
						url = a
						break
					}
				}
				cover := cast.ToStringMap(video["origin_cover"])
				covers := cast.ToStringSlice(cover["url_list"])
				var img string
				if len(covers) > 0 {
					img = covers[0]
				}
				err := filekit.WriteFileAppend(dist, []byte(fmt.Sprintf("<img height=\"200\" src=\"%s\" /><a href=\"%s\">%s</a><br><br>", img, url, key)))
				if err != nil {
					logkit.Error(err.Error())
				}
			}
		}
	})
	RequestHandleExec(page)
	//page.MustElement("#kw").MustInput("hello")
	//page.MustElement("#su").MustClick()

	page.MustWaitStable()
}
