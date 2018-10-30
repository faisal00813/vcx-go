// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 29 Oct 2018 22:41:05 IST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package vcx

/*
#cgo LDFLAGS: -lvcx
#include "vcx.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

func (x Callback_string) PassRef() (ref *C.cb_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_stringEBC8845DFunc == nil {
		callback_stringEBC8845DFunc = x
	}
	return (*C.cb_string)(C.cb_string_ebc8845d), nil
}

func (x Callback_string) PassValue() (ref C.cb_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_stringEBC8845DFunc == nil {
		callback_stringEBC8845DFunc = x
	}
	return (C.cb_string)(C.cb_string_ebc8845d), nil
}

func NewCallback_stringRef(ref unsafe.Pointer) *Callback_string {
	return (*Callback_string)(ref)
}

//export callback_stringEBC8845D
func callback_stringEBC8845D(cCmd_handle C.vcx_command_handle_t, cError_code C.vcx_error_t, cStr_value *C.char) {
	if callback_stringEBC8845DFunc != nil {
		Cmd_handleebc8845d := (Command_handle_t)(cCmd_handle)
		Error_codeebc8845d := (Error_t)(cError_code)
		Str_valueebc8845d := packPCharString(cStr_value)
		callback_stringEBC8845DFunc(Cmd_handleebc8845d, Error_codeebc8845d, Str_valueebc8845d)
		return
	}
	panic("callback func has not been set (race?)")
}

var callback_stringEBC8845DFunc Callback_string

func (x Callback_int) PassRef() (ref *C.cb_int, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_intA42EE42DFunc == nil {
		callback_intA42EE42DFunc = x
	}
	return (*C.cb_int)(C.cb_int_a42ee42d), nil
}

func (x Callback_int) PassValue() (ref C.cb_int, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_intA42EE42DFunc == nil {
		callback_intA42EE42DFunc = x
	}
	return (C.cb_int)(C.cb_int_a42ee42d), nil
}

func NewCallback_intRef(ref unsafe.Pointer) *Callback_int {
	return (*Callback_int)(ref)
}

//export callback_intA42EE42D
func callback_intA42EE42D(cCmd_handle C.vcx_command_handle_t, cError_code C.vcx_error_t, cInt_value C.uint) {
	if callback_intA42EE42DFunc != nil {
		Cmd_handlea42ee42d := (Command_handle_t)(cCmd_handle)
		Error_codea42ee42d := (Error_t)(cError_code)
		Int_valuea42ee42d := (uint32)(cInt_value)
		callback_intA42EE42DFunc(Cmd_handlea42ee42d, Error_codea42ee42d, Int_valuea42ee42d)
		return
	}
	panic("callback func has not been set (race?)")
}

var callback_intA42EE42DFunc Callback_int

func (x Callback_int_string) PassRef() (ref *C.cb_int_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_int_stringF8CA4BD5Func == nil {
		callback_int_stringF8CA4BD5Func = x
	}
	return (*C.cb_int_string)(C.cb_int_string_f8ca4bd5), nil
}

func (x Callback_int_string) PassValue() (ref C.cb_int_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_int_stringF8CA4BD5Func == nil {
		callback_int_stringF8CA4BD5Func = x
	}
	return (C.cb_int_string)(C.cb_int_string_f8ca4bd5), nil
}

func NewCallback_int_stringRef(ref unsafe.Pointer) *Callback_int_string {
	return (*Callback_int_string)(ref)
}

//export callback_int_stringF8CA4BD5
func callback_int_stringF8CA4BD5(cCmd_handle C.vcx_command_handle_t, cError_code C.vcx_error_t, cInt_value C.uint, cStr_value *C.char) {
	if callback_int_stringF8CA4BD5Func != nil {
		Cmd_handlef8ca4bd5 := (Command_handle_t)(cCmd_handle)
		Error_codef8ca4bd5 := (Error_t)(cError_code)
		Int_valuef8ca4bd5 := (uint32)(cInt_value)
		Str_valuef8ca4bd5 := packPCharString(cStr_value)
		callback_int_stringF8CA4BD5Func(Cmd_handlef8ca4bd5, Error_codef8ca4bd5, Int_valuef8ca4bd5, Str_valuef8ca4bd5)
		return
	}
	panic("callback func has not been set (race?)")
}

var callback_int_stringF8CA4BD5Func Callback_int_string

func (x Callback_enum) PassRef() (ref *C.cb_enum, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_enumF7B70E00Func == nil {
		callback_enumF7B70E00Func = x
	}
	return (*C.cb_enum)(C.cb_enum_f7b70e00), nil
}

func (x Callback_enum) PassValue() (ref C.cb_enum, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_enumF7B70E00Func == nil {
		callback_enumF7B70E00Func = x
	}
	return (C.cb_enum)(C.cb_enum_f7b70e00), nil
}

func NewCallback_enumRef(ref unsafe.Pointer) *Callback_enum {
	return (*Callback_enum)(ref)
}

//export callback_enumF7B70E00
func callback_enumF7B70E00(cCmd_handle C.vcx_command_handle_t, cError_code C.vcx_error_t, cEnum_value C.vcx_state_t) {
	if callback_enumF7B70E00Func != nil {
		Cmd_handlef7b70e00 := (Command_handle_t)(cCmd_handle)
		Error_codef7b70e00 := (Error_t)(cError_code)
		Enum_valuef7b70e00 := (State_t)(cEnum_value)
		callback_enumF7B70E00Func(Cmd_handlef7b70e00, Error_codef7b70e00, Enum_valuef7b70e00)
		return
	}
	panic("callback func has not been set (race?)")
}

var callback_enumF7B70E00Func Callback_enum

func (x Callback_error_code) PassRef() (ref *C.cb_error_code, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_error_code438ABEB5Func == nil {
		callback_error_code438ABEB5Func = x
	}
	return (*C.cb_error_code)(C.cb_error_code_438abeb5), nil
}

func (x Callback_error_code) PassValue() (ref C.cb_error_code, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if callback_error_code438ABEB5Func == nil {
		callback_error_code438ABEB5Func = x
	}
	return (C.cb_error_code)(C.cb_error_code_438abeb5), nil
}

func NewCallback_error_codeRef(ref unsafe.Pointer) *Callback_error_code {
	return (*Callback_error_code)(ref)
}

//export callback_error_code438ABEB5
func callback_error_code438ABEB5(cCmd_handle C.vcx_command_handle_t, cError_code C.vcx_error_t) {
	if callback_error_code438ABEB5Func != nil {
		Cmd_handle438abeb5 := (Command_handle_t)(cCmd_handle)
		Error_code438abeb5 := (Error_t)(cError_code)
		callback_error_code438ABEB5Func(Cmd_handle438abeb5, Error_code438abeb5)
		return
	}
	panic("callback func has not been set (race?)")
}

var callback_error_code438ABEB5Func Callback_error_code

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	str = safeString(str)
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}
