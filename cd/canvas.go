// Copyright (C) 2013 yofu. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package cd

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liupcd -lcd -lcdpdf
#cgo linux CFLAGS : -I../../libs/cd/include
#cgo linux LDFLAGS: -L../../libs/cd -liupcd -lcd -lcdpdf
#include <stdlib.h>
#include <iup.h>
#include <cd.h>
*/
import "C"
import "unsafe"
import "github.com/yofu/go-iup/iup"

type Canvas struct {
    p     *C.cdCanvas
}

type State struct {
    p     *C.cdState
}

type Nativer interface {
    Native() uintptr
}

func CreateCanvas(context *Context, canv Nativer) *Canvas {
    ptr := C.cdCreateCanvas(context.p, unsafe.Pointer(canv.Native()))
	if ptr == nil { return nil }
    return NewCanvas(ptr)
}

func CreatePrinter(context *Context, data string) *Canvas {
    d := iup.NewCS(data)
    defer iup.FreeCS(d)
    ptr := C.cdCreateCanvas(context.p, unsafe.Pointer(d))
	if ptr == nil { return nil }
    return NewCanvas(ptr)
}

func NewCanvas(p *C.cdCanvas) *Canvas {
    c := new(Canvas)
    c.p = p
    return c
}

func (cvs *Canvas) Native() uintptr {
    return (uintptr)(unsafe.Pointer(cvs.p))
}

func (cvs *Canvas) Kill() {
    C.cdKillCanvas(cvs.p)
}

func NewState(p *C.cdState) *State {
    s := new(State)
    s.p = p
    return s
}
