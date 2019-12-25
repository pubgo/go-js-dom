// +build js,wasm

package godom

import "syscall/js"

var _ Value = (*jsObject)(nil)

type Value interface {
	JSValue() js.Value
	String() string
	Type() Type
	New(args ...interface{}) Value
	Call(m string, args ...interface{}) Value
	Invoke(args ...interface{}) Value
	Bool() bool
	Float() float64
	Get(p string) Value
	Set(p string, x interface{})
	Truthy() bool
	SetIndex(i int, x interface{})
	Length() int
	Int() int
	InstanceOf(t Value) bool
	Index(i int) Value
	Delete(key string)
}

type jsObject struct {
	o js.Value
}

func (w *jsObject) Delete(key string) {
	w.o.Call("delete", key)
}

func (w *jsObject) String() string {
	return w.o.String()
}

func (w *jsObject) Type() Type {
	return w.o.Type()
}

func (w *jsObject) New(args ...interface{}) Value {
	return ValueOf(w.o.New(args...))
}

func (w *jsObject) Call(m string, args ...interface{}) Value {
	return ValueOf(w.o.Call(m, args...))
}

func (w *jsObject) Invoke(args ...interface{}) Value {
	return ValueOf(w.o.Invoke(args...))
}

func (w *jsObject) Bool() bool {
	return w.o.Bool()
}

func (w *jsObject) Float() float64 {
	return w.o.Float()
}

func (w *jsObject) Get(p string) Value {
	return ValueOf(w.o.Get(p))
}

func (w *jsObject) Set(p string, x interface{}) {
	w.o.Set(p, x)
}

func (w *jsObject) Truthy() bool {
	return w.o.Truthy()
}

func (w *jsObject) SetIndex(i int, x interface{}) {
	w.o.SetIndex(i, x)
}

func (w *jsObject) Length() int {
	return w.o.Length()
}

func (w *jsObject) Int() int {
	return w.o.Int()
}

func (w *jsObject) InstanceOf(t Value) bool {
	return w.o.InstanceOf(t.JSValue())
}

func (w *jsObject) Index(i int) Value {
	return ValueOf(w.o.Index(i))
}

func (w *jsObject) JSValue() js.Value {
	return w.o.JSValue()
}

type (
	Func = js.Func
	Error = js.Error
	Type = js.Type
	ValueError = js.ValueError
	Wrapper = js.Wrapper
)

const (
	TypeUndefined = js.TypeUndefined
	TypeNull      = js.TypeNull
	TypeBoolean   = js.TypeBoolean
	TypeNumber    = js.TypeNumber
	TypeString    = js.TypeString
	TypeSymbol    = js.TypeSymbol
	TypeObject    = js.TypeObject
	TypeFunction  = js.TypeFunction
)

func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var _args []Value
		for _, arg := range args {
			_args = append(_args, ValueOf(arg))
		}
		return fn(ValueOf(this), _args)
	})
}

var (
	_undefined = &jsObject{o: js.Undefined()}
	_global    = ValueOf(js.Global())
)

func Undefined() Value {
	return _undefined
}

func Null() Value {
	return nil
}

func Global() Value {
	return _global
}

func ValueOf(x interface{}) Value {
	if x == js.Null() {
		return nil
	}

	if x == js.Undefined() {
		return nil
	}

	return &jsObject{o: js.ValueOf(x)}
}

func CopyBytesToGo(dst []byte, src Value) int {
	return js.CopyBytesToGo(dst, src.JSValue())
}

func CopyBytesToJS(dst Value, src []byte) int {
	return js.CopyBytesToJS(dst.JSValue(), src)
}

func Eval(args ...interface{}) Value {
	return Global().Call("eval", args ...)
}

// M is a simple map type. It is intended as a shorthand for JavaScript *jsObjects (before conversion).
type M map[string]interface{}

// S is a simple slice type. It is intended as a shorthand for JavaScript arrays (before conversion).
type S []interface{}
