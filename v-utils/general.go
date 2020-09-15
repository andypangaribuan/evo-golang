package v_utils

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
const charsetLowerCase = "abcdefghijklmnopqrstuvwxyz"
const charsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))



func (*VS) RemoveAllUnusedStr(value string, args ...string) string {
	for _, arg := range args {
		value = strings.Replace(value, arg, "", -1)
	}
	return value
}


func (slf *VS) GetUUID() string {
	id := uuid.NewV4().String()
	id = slf.RemoveAllUnusedStr(id, "-")
	return id // length: 32
}


func (slf *VS) GetUniqueId() string {
	tm := slf.Time2StrMicros(time.Now().UTC())
	tm = slf.RemoveAllUnusedStr(tm, "-", " ", ":", ".")

	y1 := tm[:2]
	y2 := tm[2:4]
	M := tm[4:6]
	d := tm[6:8]
	h := tm[8:10]
	m := tm[10:12]
	s := tm[12:14]
	mc1 := tm[14:17]
	mc2 := tm[17:]

	id := mc1 + s + y1 + h + M + m + y2 + d + mc2
	return id
}


func (slf *VS) GetId100() (id string) {
	id = slf.GetUUID() + slf.GetUniqueId() + slf.RandomAlphabetNumeric(30) + slf.RandomNumeric(18)
	if len(id) > 100 {
		id = id[:100]
	}
	return
}


func (slf *VS) GetId10() (id string) {
	id = slf.GetUUID()
	if len(id) > 10 {
		id = id[:10]
	}
	return
}

func (slf *VS) GetId20() (id string) {
	id = slf.GetUUID()
	if len(id) > 20 {
		id = id[:20]
	}
	return
}

func (slf *VS) GetId30() (id string) {
	id = slf.GetUUID()
	if len(id) > 30 {
		id = id[:30]
	}
	return
}



func (*VS) RandomAlphabet(length int) string {
	return randomCharset(length, charsetLowerCase + charsetUpperCase)
}

func (*VS) RandomLowerCaseAlphabet(length int) string {
	return randomCharset(length, charsetLowerCase)
}

func (*VS) RandomUpperCaseAlphabet(length int) string {
	return randomCharset(length, charsetUpperCase)
}

func (*VS) RandomNumeric(length int) string {
	return randomCharset(length, charsetNumeric)
}

func (*VS) RandomAlphabetNumeric(length int) string {
	return randomCharset(length, charsetLowerCase + charsetUpperCase + charsetNumeric)
}

func randomCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}



func (*VS) Base64EncodeString(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func (*VS) Base64DecodeString(value string) (decode string, err error) {
	if blob, _err := base64.StdEncoding.DecodeString(value); _err == nil {
		decode = bytes.NewBuffer(blob).String()
	} else {
		err = errors.WithStack(_err)
	}
	return
}

func (*VS) MD5Encryption(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	return hex.EncodeToString(hash.Sum(nil))
}



func (*VS) ParsValidationNullOrEmpty(names []string, args ...interface{}) (msg *string) {
	pars := ""
	argsLength := len(args)

	for i, name := range names {
		isEmptyOrNull := true

		if i < argsLength {
			objRef := reflect.ValueOf(args[i])
			objKind := objRef.Kind()
			if objKind == reflect.Ptr {
				objRef = objRef.Elem()
				objKind = objRef.Kind()
			}

			if objKind != reflect.Invalid {
				switch objKind {
				case reflect.String:
					v := objRef.String()
					v = strings.TrimSpace(v)
					if v != "" {
						isEmptyOrNull = false
					}
					break
				case reflect.Slice, reflect.Array, reflect.Map:
					length := objRef.Len()
					if length > 0 {
						isEmptyOrNull = false
					}
					break
				default:
					isEmptyOrNull = false
					break
				}
			}
		}

		if isEmptyOrNull {
			if pars != "" {
				pars += ", "
			}
			pars += name
		}
	}

	if pars != "" {
		pars = "empty parameters: " + pars
		msg = &pars
	}
	return
}

var strIsDigitAndLetterRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
func (*VS) StrIsDigitAndLetter(value string) bool {
	return strIsDigitAndLetterRegex(value)
}
