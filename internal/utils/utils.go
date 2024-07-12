// vim: noexpandtab

package utils

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

func IgnoreFuncFields() cmp.Option {
	return cmp.FilterValues(func(x, y interface{}) bool {
		return reflect.TypeOf(x).Kind() == reflect.Func && reflect.TypeOf(y).Kind() == reflect.Func
	}, cmp.Ignore())
}
