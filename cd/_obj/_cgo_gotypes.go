// Created by cgo - DO NOT EDIT

package cd

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int32) { *dst = syscall.Errno(x) }
type _Ctype_void [0]byte

