package chromerob

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/mizuki1412/go-core-kit/v2/library/c"
	"github.com/mizuki1412/go-core-kit/v2/service/logkit"
	"strings"
)

type ReqResHandle struct {
	// 定义参数
	Prefix string
	Method string
	Handle func(res ResponseEntity)
	// 动态参数
	ReqId proto.NetworkRequestID
}

type ResponseEntity struct {
	Body string
}

// 处理定义. key=prefix
var handles = map[string]*ReqResHandle{}

// 处理单体. key=requestId
var entities = map[proto.NetworkRequestID]*ReqResHandle{}

// RequestHandleAdd 添加接口监听逻辑。
// prefix必须是唯一判断这个url的
func RequestHandleAdd(prefix string, method string, handle func(res ResponseEntity)) {
	handles[prefix] = &ReqResHandle{
		Handle: func(res ResponseEntity) {
			c.RecoverFuncWrapper(func() {
				handle(res)
			})
		},
		Method: method,
		Prefix: prefix}
}

// RequestHandleExec 开始监听接口返回
func RequestHandleExec(page *rod.Page) {
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
		//func(e *proto.NetworkRequestWillBeSentExtraInfo) {
		//if e.RequestID == sessionRequestId {
		//	cookieStr = e.Headers["cookie"].String()
		//	done <- 1
		//}
		//},
		// NetworkResponseReceived：容易拿不到body
		func(e *proto.NetworkLoadingFinished) {
			if v, ok := entities[e.RequestID]; ok {
				r, err := proto.NetworkGetResponseBody{RequestID: e.RequestID}.Call(page)
				if err != nil {
					logkit.Error(err.Error())
					return
				}
				v.Handle(ResponseEntity{
					Body: r.Body,
				})
			}
		})()
}
