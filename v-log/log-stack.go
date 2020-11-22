package v_log

import (
	"bytes"
	"fmt"
	"github.com/andypangaribuan/evo-golang/v-ext"
	"github.com/andypangaribuan/evo-golang/vi"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (slf *VS) Stack(args ...interface{}) (stack *string) {
	v := slf.BaseStack(2, args...)
	stack = &v
	return
}


func (slf *VS) BaseStack(skip int, args ...interface{}) (stack string) {
	pc, filePath, lineNumber, _ := runtime.Caller(skip)
	funcName := runtime.FuncForPC(pc).Name()

	format := ":: %s \n:: %s:%d"
	data := []interface{}{funcName, filePath, lineNumber}

	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			format += "\n:: %v"
			data = append(data, v)
		}
	}

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if codes := readFile(filePath, lineNumber); codes != "" {
			format += "\n:: START CODE STACK\n"
			format += "%v"
			format += "\n:: END CODE STACK"

			data = append(data, codes)
		}
	}

	for _, v := range args {
		format += "\n\n>_\n"
		format += "%+v"

		data = append(data, getLogValue(v))
	}

	format += "\n\n|=|"
	stack = fmt.Sprintf(format, data...)
	stack = removeInvalidChar(stack)
	return
}


func getLogValue(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	objRef := reflect.ValueOf(obj)
	objKind := objRef.Kind()
	if objKind == reflect.Ptr {
		objRef = objRef.Elem()
		objKind = objRef.Kind()
	}

	if objKind == reflect.Invalid {
		return "nil"
	}
	obj = objRef.Interface()

	value := ""
	switch data := obj.(type) {
	case string: value = data
	case []byte: value = bytes.NewBuffer(data).String()
	case error: value = codeStack(data) + "\n\n" + fmt.Sprintf("%+v", data)
	case v_ext.DbTxError:
		if data.Msg != "" {
			value = data.Msg
		}
		if data.Err != nil {
			if value != "" {
				value += "\n\n"
			}
			value += codeStack(data.Err) + "\n\n" + fmt.Sprintf("%+v", data.Err)
		}
	default:
		if content, err := vi.Json.JsonEncode(data); err == nil && content != "" {
			value = content
		}
		if value == "" || value == "{}" {
			value = fmt.Sprintf("%+v", data)
		}
	}

	return value
}


func removeInvalidChar(value string) string {
	/* REMOVE INVALID UTF8 CHAR */
	if !utf8.ValidString(value) {
		v := make([]rune, 0, len(value))
		for i, r := range value {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(value[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		value = string(v)
	}

	/* REMOVE NON PRINTABLE CHAR */
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}

		sc := fmt.Sprintf("%q", r)
		scn := len(sc)

		// find hex character code
		// https://golang.org/pkg/regexp/syntax/
		if scn >= 6 && sc[:1] == "'" && sc[scn-1:] == "'" && sc[1:3] == "\\x" {
			return -1
		}

		return r
	}, value)

	return clean
}
