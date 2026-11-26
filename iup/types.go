// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <iup.h>
#include <iupkey.h>

static int iupKeyBase(int c) {
#if IUP_VERSION_NUMBER >= 309000
	return iup_XkeyBase(c);
#else
	return c % 256;
#endif
}

static int iupKeyIsShift(int c) {
	return iup_isShiftXkey(c) ? 1 : 0;
}

static int iupKeyIsCtrl(int c) {
	return iup_isCtrlXkey(c) ? 1 : 0;
}

static int iupKeyIsAlt(int c) {
	return iup_isAltXkey(c) ? 1 : 0;
}

static int iupKeyIsSys(int c) {
	return iup_isSysXkey(c) ? 1 : 0;
}

static int iupKeyIsPrint(int c) {
	int b = iupKeyBase(c);
	return iup_isprint(b) ? 1 : 0;
}
*/
import "C"

/************************************************************************/
/*                   GO-IUP Version		                                */
/************************************************************************/
const (
	Name        = "GO-IUP - Iup Golang Binding"
	CopyRight   = "Copyright (C) 2011-2012 visualfc"
	Version     = "1.1"
	VersionDate = "2012.9.26"
)

/************************************************************************/
/*                   IUP Version		                                */
/************************************************************************/
const (
	IupName          = C.IUP_NAME
	IupCopyRight     = C.IUP_COPYRIGHT
	IupDescription   = C.IUP_DESCRIPTION
	IupVersion       = C.IUP_VERSION
	IupVersionNumber = C.IUP_VERSION_NUMBER
	IupVersionDate   = C.IUP_VERSION_DATE
)

/************************************************************************/
/*                   Common Return Values                               */
/************************************************************************/
const (
	ERROR   int = C.IUP_ERROR
	NOERROR     = C.IUP_NOERROR
	OPENDE      = C.IUP_OPENED
	INVALID     = C.IUP_INVALID
)

/************************************************************************/
/*                   Callback Return Values                             */
/************************************************************************/
const (
	IGNORE   int = C.IUP_IGNORE
	DEFAULT      = C.IUP_DEFAULT
	CLOSE        = C.IUP_CLOSE
	CONTINUE     = C.IUP_CONTINUE
)

/************************************************************************/
/*           IupPopup and IupShowXY Parameter Values                    */
/************************************************************************/
const (
	CENTER       int = C.IUP_CENTER       //0xFFFF  /* 65535 */
	LEFT             = C.IUP_LEFT         // 0xFFFE  /* 65534 */
	RIGHT            = C.IUP_RIGHT        // 0xFFFD  /* 65533 */
	MOUSEPOS         = C.IUP_MOUSEPOS     // 0xFFFC  /* 65532 */
	CURIUP           = C.IUP_CURRENT      // 0xFFFB  /* 65531 */
	CENTERPARENT     = C.IUP_CENTERPARENT // 0xFFFA  /* 65530 */
)

/************************************************************************/
/*               SHOW_CB Callback Values                                */
/************************************************************************/
const (
	SHOW     = C.IUP_SHOW
	RESTORE  = C.IUP_RESTORE
	MINIMIZE = C.IUP_MINIMIZE
	MAXIMIZE = C.IUP_MAXIMIZE
	HIDE     = C.IUP_HIDE
)

/************************************************************************/
/*               SCROLL_CB Callback Values                              */
/************************************************************************/
const (
	SBUP      = C.IUP_SBUP
	SBDN      = C.IUP_SBUP
	SBPGUP    = C.IUP_SBPGUP
	SBPGDN    = C.IUP_SBPGDN
	SBPOSV    = C.IUP_SBPOSV
	SBDRAGV   = C.IUP_SBDRAGV
	SBLEFT    = C.IUP_SBLEFT
	SBRIGHT   = C.IUP_SBRIGHT
	SBPGLEFT  = C.IUP_SBPGLEFT
	SBPGRIGHT = C.IUP_SBPGRIGHT
	SBPOSH    = C.IUP_SBPOSH
	SBDRAGH   = C.IUP_SBDRAGH
)

/************************************************************************/
/*               Keyboard Key Constants                                 */
/************************************************************************/
const (
	K_BS    = C.K_BS
	K_TAB   = C.K_TAB
	K_CR    = C.K_CR
	K_SP    = C.K_SP
	K_DEL   = C.K_DEL
	K_ESC   = C.K_ESC
	K_HOME  = C.K_HOME
	K_END   = C.K_END
	K_PGUP  = C.K_PGUP
	K_PGDN  = C.K_PGDN
	K_INS   = C.K_INS
	K_UP    = C.K_UP
	K_LEFT  = C.K_LEFT
	K_RIGHT = C.K_RIGHT
	K_DOWN  = C.K_DOWN
	K_F1    = C.K_F1
	K_F2    = C.K_F2
	K_F3    = C.K_F3
	K_F4    = C.K_F4
	K_F5    = C.K_F5
	K_F6    = C.K_F6
	K_F7    = C.K_F7
	K_F8    = C.K_F8
	K_F9    = C.K_F9
	K_F10   = C.K_F10
	K_F11   = C.K_F11
	K_F12   = C.K_F12
)

type KeyState int

func (k KeyState) Key() int {
	return int(C.iupKeyBase(C.int(k)))
}

func (k KeyState) Xkey() int {
	return int(k % 128)
}

func (k KeyState) IsShift() bool {
	return C.iupKeyIsShift(C.int(k)) != 0
}

func (k KeyState) IsCtrl() bool {
	return C.iupKeyIsCtrl(C.int(k)) != 0
}

func (k KeyState) IsAlt() bool {
	return C.iupKeyIsAlt(C.int(k)) != 0
}

func (k KeyState) IsSys() bool {
	return C.iupKeyIsSys(C.int(k)) != 0
}

func (k KeyState) IsPrint() bool {
	return C.iupKeyIsPrint(C.int(k)) != 0
}

/************************************************************************/
/*               Mouse Button Values and State                         */
/************************************************************************/
const (
	BUTTON1 = '1'
	BUTTON2 = '2'
	BUTTON3 = '3'
	BUTTON4 = '4'
	BUTTON5 = '5'
)

type MouseState struct {
	S string
}

func (s *MouseState) IsShift() bool {
	return s.S[0] == 'S'
}

func (s *MouseState) IsControl() bool {
	return s.S[1] == 'C'
}

func (s *MouseState) IsButton1() bool {
	return s.S[2] == '1'
}

func (s *MouseState) IsButton2() bool {
	return s.S[3] == '2'
}

func (s *MouseState) IsButton3() bool {
	return s.S[4] == '3'
}

func (s *MouseState) IsDouble() bool {
	return s.S[5] == 'D'
}

func (s *MouseState) IsAlt() bool {
	return s.S[6] == 'A'
}

func (s *MouseState) IsSys() bool {
	return s.S[7] == 'Y'
}

func (s *MouseState) IsButton4() bool {
	return s.S[8] == '4'
}

func (s *MouseState) IsButton5() bool {
	return s.S[9] == '5'
}

/************************************************************************/
/*                      Pre-Defined Masks                               */
/************************************************************************/
const (
	MASK_FLOAT  = "[+/-]?(/d+/.?/d*|/./d+)"
	MASK_UFLOAT = "(/d+/.?/d*|/./d+)"
	MASK_EFLOAT = "[+/-]?(/d+/.?/d*|/./d+)([eE][+/-]?/d+)?"
	MASK_INT    = "[+/-]?/d+"
	MASK_UINT   = "/d+"
)

/************************************************************************/
/*                   Record Input Modes                                 */
/************************************************************************/
const (
	RECBINARY   = C.IUP_RECBINARY
	IUP_RECTEXT = C.IUP_RECTEXT
)
