package models

import (
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

var nilTime = (time.Time{}).UnixNano()

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(dateLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dateLayout))), nil
}

type Filter struct {
	SortColumn string
	SortAsc    string
}

func NewFilter(filter string) *Filter {
	newFilter := &Filter{}
	runes := []rune(filter)
	if string(runes[0:1]) != "-" {
		newFilter.SortAsc = "ASC"
		newFilter.SortColumn = filter
	} else {
		newFilter.SortAsc = "DESC"
		newFilter.SortColumn = string(runes[1:])
	}

	return newFilter
}
