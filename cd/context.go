// Copyright (C) 2013 yofu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package cd

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liupcd -lcd
#cgo linux CFLAGS : -I../../libs/cd/include
#cgo linux LDFLAGS: -L../../libs/cd -liupcd -lcd
#include <stdlib.h>
#include <iup.h>
#include <cd.h>
#include <cdiup.h>
#include <cdprint.h>
#include <cddxf.h>
#include <cddbuf.h>
#include <cdirgb.h>
#include <cdnative.h>
#include <cdgdiplus.h>
*/
import "C"
import "unsafe"

var (
	CD_IUP          = NewContext(C.cdContextIup())
	CD_DXF          = NewContext(C.cdContextDXF())
	CD_PRINTER      = NewContext(C.cdContextPrinter())
	CD_DBUFFER      = NewContext(C.cdContextDBuffer())
	CD_DBUFFERRGB   = NewContext(C.cdContextDBufferRGB())
	CD_NATIVEWINDOW = NewContext(C.cdContextNativeWindow())
)

type Context struct {
	p *C.cdContext
}

func NewContext(p *C.cdContext) *Context {
	c := new(Context)
	c.p = p
	return c
}

func (con *Context) Native() uintptr {
	return (uintptr)(unsafe.Pointer(con.p))
}

// func UseContextPlus(use bool) {
// 	C.cdInitContextPlus()
// 	if use {
// 		fmt.Println("use context plus")
// 		C.cdUseContextPlus(1)
// 	} else {
// 		C.cdUseContextPlus(0)
// 	}
// }
