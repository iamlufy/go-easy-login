package mocks

import (
	"bou.ke/monkey"
	"reflect"
)

func InstanceMethod(target interface{}, methodName string, replacement interface{}) {

	monkey.PatchInstanceMethod(reflect.TypeOf(target), methodName, replacement)
}
func ResetMethod(target interface{}, method string) {
	monkey.UnpatchInstanceMethod(reflect.TypeOf(target), method)
}

func MockFunc(target interface{}, replaceFunc interface{}) {
	monkey.Patch(target, replaceFunc)
}
