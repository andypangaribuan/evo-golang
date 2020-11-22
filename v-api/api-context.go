package v_api

import (
	"github.com/andypangaribuan/evo-golang/v-ext"
	"github.com/andypangaribuan/evo-golang/vi"
	"github.com/pkg/errors"
	"reflect"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */


func (slf *Context) GetRequestBody() (string, error) {
	return vi.Utils.GetEchoRequestBody(slf.echo.c)
}


func (slf *Context) BindPars(destination interface{}) error {
	destValue := reflect.ValueOf(destination)
	if destValue.Kind() != reflect.Ptr {
		return errors.New("destination must pass a pointer")
	}
	if destValue.IsNil() {
		return errors.New("can not pass a nil pointer")
	}

	if slf.echo != nil {
		return vi.Utils.BindEchoBodyToPars(slf.echo.c, destination)
	}
	return errors.New("404")
}



func (slf *Context) GetRequestIp() (reqIp string) {
	if slf.echo != nil {
		reqIp = slf.echo.c.RealIP()
	}
	return
}


func (slf *Context) GetTraceId() string {
	return vi.Utils.GetTraceId(slf.echo.c)
}

func (slf *Context) SetRequestUserId(uid string) {
	if slf.echo != nil {
		slf.echo.c.Request().Header.Add(v_ext.Conf.RequestUidKey(), uid)
	}
}





func (slf *Context) ResponseStr(code int, response string) error {
	if slf.echo != nil {
		return slf.echo.c.String(code, response)
	}

	return nil
}


func (slf *Context) ResponseJson(code int, response interface{}) error {
	if slf.echo != nil {
		return slf.echo.c.JSON(code, response)
	}

	return nil
}


func (slf *Context) ResponseJsonBlob(code int, response []byte) error {
	if slf.echo != nil {
		return slf.echo.c.JSONBlob(code, response)
	}

	return nil
}
