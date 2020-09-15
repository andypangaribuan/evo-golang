package v_var

import (
	"evo-lib/v-ext"
	"strings"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (*VS) StrTrimSpace(value string) string {
	return strings.TrimSpace(value)
}


func (*VS) PtrStrTrimSpace(value *string) *string {
	var newValue *string

	if value != nil {
		v := *value
		v = strings.TrimSpace(v)
		newValue = &v
	}

	return newValue
}


func (*VS) StrTrimMaxLength(value string, maxLength int) string {
	if len(value) > maxLength {
		return value[:maxLength]
	}
	return value
}


func (*VS) PtrStrTrimMaxLength(value *string, maxLength int) *string {
	if value != nil && len(*value) > maxLength {
		v := *value
		v = v[:maxLength]
		return &v
	}
	return value
}


func (*VS) StrSplitMaxDbText(value string) (txt string, arr []string) {
	txt = value
	arr = make([]string, 0)

	if value == "" {
		return
	}

	switch v_ext.Conf.DbType {
	case "postgres": return
	case "mysql":
		val := value
		for {
			length := len(val)
			if length == 0 {
				break
			}

			split := val
			if length > 65000 {
				split = val[:65000]
			}
			arr = append(arr, split)
			val = val[len(split):]
		}
	}

	return
}


func (slf *VS) PtrStrSplitMaxDbText(value *string) (txt *string, arr []string) {
	if value == nil {
		return
	}

	v, arr := slf.StrSplitMaxDbText(*value)
	return &v, arr
}
