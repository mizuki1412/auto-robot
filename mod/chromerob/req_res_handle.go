package chromerob

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"strings"
)

type ReqResHandle struct {
	// 定义参数
	Prefix string
	Method string
	Handle func(*proto.NetworkResponseReceived)
	// 动态参数
	ReqId proto.NetworkRequestID
}

// 处理定义. key=prefix
var handles = map[string]*ReqResHandle{}

// 处理单体. key=requestId
var entities = map[proto.NetworkRequestID]*ReqResHandle{}

// RequestHandleAdd 添加接口监听逻辑。
// prefix必须是唯一判断这个url的
func RequestHandleAdd(prefix string, method string, handle func(*proto.NetworkResponseReceived)) {
	handles[prefix] = &ReqResHandle{Handle: handle, Method: method, Prefix: prefix}
}

// RequestHandleExc 开始监听接口返回
func RequestHandleExc(page *rod.Page) {
	go page.EachEvent(
		func(e *proto.NetworkRequestWillBeSent) {
			for k, v := range handles {
				if strings.Index(e.Request.URL, k) >= 0 && v.Method == e.Request.Method {
					entities[e.RequestID] = &ReqResHandle{
						Prefix: v.Prefix,
						Method: v.Method,
						Handle: v.Handle,
						ReqId:  e.RequestID,
					}
				}
			}

		},
		func(e *proto.NetworkRequestWillBeSentExtraInfo) {
			//if e.RequestID == sessionRequestId {
			//	cookieStr = e.Headers["cookie"].String()
			//	done <- 1
			//}
		}, func(e *proto.NetworkResponseReceived) {
			if v, ok := entities[e.RequestID]; ok {
				v.Handle(e)
			}
		})()
}
