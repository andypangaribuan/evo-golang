package clog

import (
	"bytes"
	"github.com/andypangaribuan/evo-golang/v-ext"
	"github.com/andypangaribuan/evo-golang/vi"
	"log"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func sendLogMiddlewareV1(model LogMiddlewareV1) {
	send(model, "/save/log-middleware/v1")
}

func SendLogTraceV1(model LogTraceV1) {
	send(model, "/save/log-trace/v1")
}

func SendLogErrorV1(model LogErrorV1) {
	send(model, "/save/log-error/v1")
}

func SendLogQuery(model LogQueryV1) {
	send(model, "/save/log-query/v1")
}


func send(model interface{}, subUrl string) {
	var stack *string
	url := v_ext.Store.CLogBaseUrl + subUrl

	if blob, err := vi.Json.JsonMarshal(model); err != nil {
		stack = vi.Log.Stack(url, err)
	} else {
		httpBlob, httpCode, err := vi.Utils.HttpPostData(url, nil, blob, true, nil)
		if err != nil {
			stack = vi.Log.Stack(url, blob, err)
			log.Println(stack)
		} else {
			if httpCode != 200 {
				log.Printf("url: %v\nhttp-code: %v\nresponse: %v\n", url, httpCode, bytes.NewBuffer(httpBlob).String())
			}
		}
	}

	if stack != nil {
		log.Println(stack)
	}
}
