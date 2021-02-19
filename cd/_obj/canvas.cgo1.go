// Created by cgo - DO NOT EDIT

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:5
package cd
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:18

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:17
import "unsafe"
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:20

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:19
type Canvas struct {
	p *[0]byte
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:24

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:23
func (con *Context) CreateCanvas(value uintptr, a ...interface{}) *Canvas {
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:27

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:26
	ptr := _Cfunc_cdCreateCanvas(con.p, iupcanv)
											return NewCanvas(ptr)
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:31

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:30
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:33

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:32
func NewCanvas(p *[0]byte,) *Canvas {
											c := new(Canvas)
											c.p = p
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:39

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:38
	return c
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:42

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:41
func (cvs *Canvas) Native() uintptr {
	return (uintptr)(unsafe.Pointer(cvs.p))
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:46

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\canvas.go:45
func (cvs *Canvas) toNative() *[0]byte {
	return (*[0]byte)(unsafe.Pointer(cvs.Native()))
}
