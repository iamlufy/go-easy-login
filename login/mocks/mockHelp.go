package mocks

import (
	"github.com/bouk/monkey"
	"reflect"
)

func InstanceMethod(target interface{}, methodName string, replacement interface{}) {

	monkey.PatchInstanceMethod(reflect.TypeOf(target), methodName, replacement)
}
func ResetMethod(target interface{}, method string) {
	monkey.UnpatchInstanceMethod(reflect.TypeOf(target), method)
}
