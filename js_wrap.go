// +build js,wasm

package godom

import "syscall/js"

type Func = js.Func
type Error = js.Error
type Type = js.Type
type Value = js.Value
type ValueError = js.ValueError
type Wrapper = js.Wrapper

func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	return js.FuncOf(fn)
}

func Undefined() Value {
	return js.Undefined()
}

func Null() Value {
	return js.Null()
}

func Global() Value {
	return js.Global()
}

func ValueOf(x interface{}) Value {
	return js.ValueOf(x)
}

func CopyBytesToGo(dst []byte, src Value) int {
	return js.CopyBytesToGo(dst, src)
}

func CopyBytesToJS(dst Value, src []byte) int {
	return js.CopyBytesToJS(dst, src)
}

func Eval(args ...interface{}) Value {
	return Global().Call("eval", args ...)
}
