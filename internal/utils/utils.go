// vim: noexpandtab

package utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
)

func IgnoreFuncFields() cmp.Option {
	return cmp.FilterValues(func(x, y interface{}) bool {
		return reflect.TypeOf(x).Kind() == reflect.Func && reflect.TypeOf(y).Kind() == reflect.Func
	}, cmp.Ignore())
}

func HourMinuteStringToTime(str string) (*time.Time, error) {
	now := time.Now()
	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), str))
	if err != nil {
		return nil, err
	}
	return &t, nil
}
