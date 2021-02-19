// Created by cgo - DO NOT EDIT

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:5
package cd
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:18

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:17
import "unsafe"
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:20

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:19
type Context struct {
	p *[0]byte
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:22
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:27

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:26
func NewContext(p *[0]byte,) *Context {
											c := new(Context)
											c.p = p
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:33

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:32
	return c
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:36

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:35
func (con *Context) Native() uintptr {
	return (uintptr)(unsafe.Pointer(con.p))
}
//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:40

//line C:\Users\fukushima\golang\src\github.com\yofu\go-iup\cd\context.go:39
func (con *Context) toNative() *[0]byte {
	return (*[0]byte)(unsafe.Pointer(con.Native()))
}
