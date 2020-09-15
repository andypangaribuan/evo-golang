package v_utils

import (
	"bytes"
	"evo-lib/v-ext"
	"evo-lib/vi"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"io/ioutil"
	"reflect"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */

func (*VS) GetEchoRequestBody(c echo.Context) (body string, err error) {
	blob, _err := getEchoRequestBodyBlob(c)
	if _err != nil {
		err = _err
		return
	}

	body = bytes.NewBuffer(blob).String()
	return
}


func (slf *VS) BindEchoBodyToPars(c echo.Context, destination interface{}) error {
	destValue := reflect.ValueOf(destination)
	if destValue.Kind() != reflect.Ptr {
		return errors.New("destination must pass a pointer")
	}
	if destValue.IsNil() {
		return errors.New("can not pass a nil pointer")
	}

	blob, err := getEchoRequestBodyBlob(c)
	if err != nil {
		return err
	}

	err = vi.Json.JsonUnMarshal(blob, destination)
	if err != nil {
		err = errors.WithStack(err)
	}
	return err
}


func getEchoRequestBodyBlob(c echo.Context) (blob []byte, err error) {
	req := c.Request()
	reader := req.Body
	if reader == nil {
		err = errors.New("the request body reader are nil")
		return
	}

	blob, err = ioutil.ReadAll(reader)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(blob))
	return
}





func (slf *VS) GetTraceId(c echo.Context) (traceId string) {
	req := c.Request()
	traceId = req.Header.Get(v_ext.Conf.TraceIdKey())
	if traceId == "" {
		traceId = slf.GetId100()
		req.Header.Add(v_ext.Conf.TraceIdKey(), traceId)
	}
	return
}
