package vi

import (
	"github.com/labstack/echo/v4"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2020/06/17
	Copyright BoltIdea. All rights reserved.
   ============================================ */
var Log ivLog
var Json ivJson
var Utils ivUtils


type ivLog interface {
	Stack(args ...interface{}) (stack *string)
	BaseStack(skip int, args ...interface{}) (stack string)
}


type ivJson interface {
	JsonMarshal(obj interface{}) ([]byte, error)
	JsonUnMarshal(data []byte, out interface{}) error
	JsonEncode(obj interface{}) (string, error)
}


type ivUtils interface {
	GetEchoRequestBody(c echo.Context) (body string, err error)
	BindEchoBodyToPars(c echo.Context, destination interface{}) error
	GetTraceId(c echo.Context) (traceId string)
	HttpPostData(url string, header map[string]string, data []byte, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error)
}
