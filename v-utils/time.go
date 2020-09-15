package v_utils

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
const (
	layoutTimeDate = "2006-01-02"
	layoutTimeFull = "2006-01-02 15:04:05"
	layoutTimeMillis = "2006-01-02 15:04:05.000"
	layoutTimeMicros = "2006-01-02 15:04:05.000000"
)



func (slf *VS) RemoveUnIntTime(value string) string {
	return slf.RemoveAllUnusedStr(value, "-", " ", ":", ".")
}



func (slf *VS) TimeNowUtcMicros() (dt time.Time, dts string) {
	dt = time.Now().UTC()
	dts = slf.Time2StrMicros(dt)
	return
}



func (*VS) Time2Str(tm time.Time, format string) string {
	replacer := [][]string{
		{"yyyy", "2006"},
		{"MM", "01"},
		{"dd", "02"},
		{"HH", "15"},
		{"mm", "04"},
		{"ss", "05"},
		{"SSSSSS", "000000"},
		{"SSSSS", "00000"},
		{"SSSS", "0000"},
		{"SSS", "000"},
		{"SS", "00"},
		{"S", "0"},
	}

	for _, arr := range replacer  {
		format = strings.Replace(format, arr[0], arr[1], -1)
	}

	return tm.Format(format)
}

func (*VS) Time2StrDate(tm time.Time) string {
	return tm.Format(layoutTimeDate)
}

func (*VS) Time2StrFull(tm time.Time) string {
	return tm.Format(layoutTimeFull)
}

func (*VS) Time2StrMillis(tm time.Time) string {
	return tm.Format(layoutTimeMillis)
}

func (*VS) Time2StrMicros(tm time.Time) string {
	return tm.Format(layoutTimeMicros)
}



func str2Time(layout string, value string) (tm time.Time, err error) {
	tm, err = time.Parse(layout, value)
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

func (*VS) Str2TimeDate(value string) (tm time.Time, err error) {
	return str2Time(layoutTimeDate, value)
}

func (*VS) Str2TimeFull(value string) (tm time.Time, err error) {
	return str2Time(layoutTimeFull, value)
}

func (*VS) Str2TimeMillis(value string) (tm time.Time, err error) {
	return str2Time(layoutTimeMillis, value)
}

func (*VS) Str2TimeMicros(value string) (tm time.Time, err error) {
	return str2Time(layoutTimeMicros, value)
}

