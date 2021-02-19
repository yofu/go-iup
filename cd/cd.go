// Copyright (C) 2013 yofu. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package cd

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liupcd -lcd -lgdiplus
#cgo linux CFLAGS : -I../../libs/cd/include
#cgo linux LDFLAGS: -L../../libs/cd -liupcd -lcd -lgdiplus
#include <stdlib.h>
#include <iup.h>
#include <cd.h>
#include <cdiup.h>
*/
import "C"
import "fmt"
import "unsafe"
import "github.com/yofu/go-iup/iup"

const CD_QUERY = -1

const (                           /* bitmap type */
    CD_RGB = iota                 /* these definitions are compatible with the IM library */
    CD_MAP
    CD_RGBA = 0x100
)

const (                         /* bitmap data */
    CD_IRED = iota
    CD_IGREEN
    CD_IBLUE
    CD_IALPHA
    CD_INDEX
    CD_COLORS
)

const (                          /* status report */
    CD_ERROR = -1
    CD_OK    =  0
)

const (                          /* clip mode */
    CD_CLIPOFF = iota
    CD_CLIPAREA
    CD_CLIPPOLYGON
    CD_CLIPREGION
)

const (                          /* region combine mode */
    CD_UNION = iota
    CD_INTERSECT
    CD_DIFFERENCE
    CD_NOTINTERSECT
)

const (                          /* polygon mode (begin...end) */
    CD_FILL = iota
    CD_OPEN_LINES
    CD_CLOSED_LINES
    CD_CLIP
    CD_BEZIER
    CD_REGION
    CD_PATH
)

const (                          /* path actions */
    CD_PATH_NEW = iota
    CD_PATH_MOVETO
    CD_PATH_LINETO
    CD_PATH_ARC
    CD_PATH_CURVETO
    CD_PATH_CLOSE
    CD_PATH_FILL
    CD_PATH_STROKE
    CD_PATH_FILLSTROKE
    CD_PATH_CLIP
)

const (                          /* fill mode */
    CD_EVENODD = iota
    CD_WINDING
)

const (                          /* line join  */
    CD_MITER = iota
    CD_BEVEL
    CD_ROUND
)

const (                          /* line cap  */
    CD_CAPFLAT   = iota
    CD_CAPSQUARE
    CD_CAPROUND
)

const (                          /* background opacity mode */
    CD_OPAQUE = iota
    CD_TRANSPARENT
)

const (                          /* write mode */
    CD_REPLACE = iota
    CD_XOR
    CD_NOT_XOR
)

const (                          /* color allocation mode (palette) */
    CD_POLITE = iota
    CD_FORCE
)

const (                          /* line style */
    CD_CONTINUOUS = iota
    CD_DASHED
    CD_DOTTED
    CD_DASH_DOT
    CD_DASH_DOT_DOT
    CD_CUSTOM
)

const (                          /* hatch type */
    CD_HORIZONTAL = iota
    CD_VERTICAL
    CD_FDIAGONAL
    CD_BDIAGONAL
    CD_CROSS
    CD_DIAGCROSS
)

const (                          /* interior style */
    CD_SOLID = iota
    CD_HATCH
    CD_STIPPLE
    CD_PATTERN
    CD_HOLLOW
)

const (                          /* text alignment */
    CD_NORTH = iota
    CD_SOUTH
    CD_EAST
    CD_WEST
    CD_NORTH_EAST
    CD_NORTH_WEST
    CD_SOUTH_EAST
    CD_SOUTH_WEST
    CD_CENTER
    CD_BASE_LEFT
    CD_BASE_CENTER
    CD_BASE_RIGHT
)

const (                          /* style */
    CD_PLAIN  = 0
    CD_BOLD   = 1
    CD_ITALIC = 2
    CD_UNDERLINE = 4
    CD_STRIKEOUT = 8
)

const (                          /* some font sizes */
    CD_SMALL    =  8
    CD_STANDARD = 12
    CD_LARGE    = 18
)

var (
/* some predefined colors for convenience */
    CD_RED          = 0xFF0000   /* 255,  0,  0 */
    CD_DARK_RED     = 0x800000   /* 128,  0,  0 */
    CD_GREEN        = 0x00FF00   /*   0,255,  0 */
    CD_DARK_GREEN   = 0x008000   /*   0,128,  0 */
    CD_BLUE         = 0x0000FF   /*   0,  0,255 */
    CD_DARK_BLUE    = 0x000080   /*   0,  0,128 */

    CD_YELLOW       = 0xFFFF00   /* 255,255,  0 */
    CD_DARK_YELLOW  = 0x808000   /* 128,128,  0 */
    CD_MAGENTA      = 0xFF00FF   /* 255,  0,255 */
    CD_DARK_MAGENTA = 0x800080   /* 128,  0,128 */
    CD_CYAN         = 0x00FFFF   /*   0,255,255 */
    CD_DARK_CYAN    = 0x008080   /*   0,128,128 */

    CD_WHITE        = 0xFFFFFF   /* 255,255,255 */
    CD_BLACK        = 0x000000   /*   0,  0,  0 */

    CD_DARK_GRAY    = 0x808080   /* 128,128,128 */
    CD_GRAY         = 0xC0C0C0   /* 192,192,192 */
)

const (
    CD_A0 = iota
    CD_A1
    CD_A2
    CD_A3
    CD_A4
    CD_A5
    CD_LETTER
    CD_LEGAL
)

// init
func (cvs *Canvas) Activate() {
    C.cdCanvasActivate(cvs.p)
}
func (cvs *Canvas) Deactivate() {
    C.cdCanvasDeactivate(cvs.p)
}
// func UseContextPlus (use int) int {
//     return int(C.cdUseContextPlus(C.int(use)))
// }
// func InitContextPlus () {
//     C.cdInitContextPlus()
// }
// func FinishContextPlus () {
//     C.cdFinishContextPlus()
// }
// func ContextIsPlus (context *Context) bool {
//     fmt.Println(int(C.cdContextIsPlus(context.p)))
//     if int(C.cdContextIsPlus(context.p))==1 {
//         return true
//     } else {
//         return false
//     }
// }

// control
func (cvs *Canvas) Flush () {
    C.cdCanvasFlush(cvs.p)
}
func (cvs *Canvas) Clear () {
    C.cdCanvasClear(cvs.p)
}
func (cvs *Canvas) SaveState () *State {
    return NewState(C.cdCanvasSaveState(cvs.p))
}
func (cvs *Canvas) RestoreState (state *State) {
    C.cdCanvasRestoreState(cvs.p, state.p)
}
func ReleaseState (state *State) {
    C.cdReleaseState(state.p)
}
func (cvs *Canvas) SetAttribute (name, data string) {
    n := iup.NewCS(name)
    d := iup.NewCS(data)
    defer iup.FreeCS(n)
    defer iup.FreeCS(d)
    C.cdCanvasSetAttribute(cvs.p, (*C.char)(n), (*C.char)(d))
}
func (cvs *Canvas) GetAttribute (name string) string {
    n := iup.NewCS(name)
    defer iup.FreeCS(n)
    return C.GoString(C.cdCanvasGetAttribute(cvs.p, (*C.char)(n)))
}

func (cvs *Canvas) GetSize() (width, height int) {
    var w, h C.int
    C.cdCanvasGetSize(cvs.p, &w, &h, nil, nil)
    return int(w), int(h)
}
func (cvs *Canvas) GetSizemm() (width, height float64) {
    var w, h C.double
    C.cdCanvasGetSize(cvs.p, nil, nil, &w, &h)
    return float64(w), float64(h)
}
func (cvs *Canvas) UpdateYAxis(yaxis *int32) int {
    yold := C.cdCanvasUpdateYAxis(cvs.p, (*C.int)(unsafe.Pointer(yaxis)))
    return int(yold)
}

func (cvs *Canvas) TransformRotate (angle float64) {
    C.cdCanvasTransformRotate(cvs.p, C.double(angle))
}

// clipping
func (cvs *Canvas) Clip (mode int) int {
    return int(C.cdCanvasClip(cvs.p, C.int(mode)))
}
func (cvs *Canvas) ClipArea (xmin, xmax, ymin, ymax int) {
    C.cdCanvasClipArea(cvs.p, C.int(xmin), C.int(xmax), C.int(ymin), C.int(ymax))
}
func (cvs *Canvas) GetClipArea (xmin, xmax, ymin, ymax *int) int {
    return int(C.cdCanvasGetClipArea(cvs.p, (*C.int)(unsafe.Pointer(xmin)), (*C.int)(unsafe.Pointer(xmax)), (*C.int)(unsafe.Pointer(ymin)), (*C.int)(unsafe.Pointer(ymax))))
}
func (cvs *Canvas) FClipArea (xmin, xmax, ymin, ymax float64) {
    C.cdfCanvasClipArea(cvs.p, C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax))
}
func (cvs *Canvas) FGetClipArea (xmin, xmax, ymin, ymax *float64) int {
    return int(C.cdfCanvasGetClipArea(cvs.p, (*C.double)(unsafe.Pointer(xmin)), (*C.double)(unsafe.Pointer(xmax)), (*C.double)(unsafe.Pointer(ymin)), (*C.double)(unsafe.Pointer(ymax))))
}

// primitives
func (cvs *Canvas) Pixel (x, y, color int) {
    C.cdCanvasPixel(cvs.p, C.int(x), C.int(y), C.long(color))
}
func (cvs *Canvas) Mark (x, y int) {
    C.cdCanvasMark(cvs.p, C.int(x), C.int(y))
}

func (cvs *Canvas) Begin (mode int) {
    C.cdCanvasBegin(cvs.p, C.int(mode))
}
func (cvs *Canvas) PathSet (action int) {
    C.cdCanvasPathSet(cvs.p, C.int(action))
}
func (cvs *Canvas) End () {
    C.cdCanvasEnd(cvs.p)
}
func (cvs *Canvas) Polygon (mode int, coords... []float64) {
    if len(coords) < 3 { return }
    cvs.Begin(mode)
    cvs.FVertex(coords[0][0], coords[0][1])
    for i:=1; i<len(coords); i++ {
        cvs.FVertex(coords[i][0], coords[i][1])
    }
    cvs.End()
}

func (cvs *Canvas) Line (x1, y1, x2, y2 int) {
    C.cdCanvasLine(cvs.p, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (cvs *Canvas) Vertex (x, y int) {
    C.cdCanvasVertex(cvs.p, C.int(x), C.int(y))
}
func (cvs *Canvas) Rect (xmin, xmax, ymin, ymax int) {
    C.cdCanvasRect(cvs.p, C.int(xmin), C.int(xmax), C.int(ymin), C.int(ymax))
}
func (cvs *Canvas) Box (xmin, xmax, ymin, ymax int) {
    C.cdCanvasBox(cvs.p, C.int(xmin), C.int(xmax), C.int(ymin), C.int(ymax))
}
func (cvs *Canvas) Arc (xc, yc, w, h int, angle1, angle2 float64) {
    C.cdCanvasArc(cvs.p, C.int(xc), C.int(yc), C.int(w), C.int(h), C.double(angle1), C.double(angle2))
}
func (cvs *Canvas) Sector (xc, yc, w, h int, angle1, angle2 float64) {
    C.cdCanvasSector(cvs.p, C.int(xc), C.int(yc), C.int(w), C.int(h), C.double(angle1), C.double(angle2))
}
func (cvs *Canvas) Circle (xc, yc, size int) {
    C.cdCanvasArc(cvs.p, C.int(xc), C.int(yc), C.int(size), C.int(size), C.double(0.0), C.double(360.0))
}
func (cvs *Canvas) FilledCircle (xc, yc, size int) {
    C.cdCanvasSector(cvs.p, C.int(xc), C.int(yc), C.int(size), C.int(size), C.double(0.0), C.double(360.0))
}
func (cvs *Canvas) Text (x, y int, s string) {
    str := iup.NewCS(s)
    defer iup.FreeCS(str)
    C.cdCanvasText(cvs.p, C.int(x), C.int(y), (*C.char)(str))
}
func (cvs *Canvas) FLine (x1, y1, x2, y2 float64) {
    C.cdfCanvasLine(cvs.p, C.double(x1), C.double(y1), C.double(x2), C.double(y2))
}
func (cvs *Canvas) FVertex (x, y float64) {
    C.cdfCanvasVertex(cvs.p, C.double(x), C.double(y))
}
func (cvs *Canvas) FRect (xmin, xmax, ymin, ymax float64) {
    C.cdfCanvasRect(cvs.p, C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax))
}
func (cvs *Canvas) FBox (xmin, xmax, ymin, ymax float64) {
    C.cdfCanvasBox(cvs.p, C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax))
}
func (cvs *Canvas) FArc (xc, yc, w, h, angle1, angle2 float64) {
    C.cdfCanvasArc(cvs.p, C.double(xc), C.double(yc), C.double(w), C.double(h), C.double(angle1), C.double(angle2))
}
func (cvs *Canvas) FSector (xc, yc, w, h, angle1, angle2 float64) {
    C.cdfCanvasSector(cvs.p, C.double(xc), C.double(yc), C.double(w), C.double(h), C.double(angle1), C.double(angle2))
}
func (cvs *Canvas) FCircle (xc, yc, size float64) {
    C.cdfCanvasArc(cvs.p, C.double(xc), C.double(yc), C.double(size), C.double(size), C.double(0.0), C.double(360.0))
}
func (cvs *Canvas) FFilledCircle (xc, yc, size float64) {
    C.cdfCanvasSector(cvs.p, C.double(xc), C.double(yc), C.double(size), C.double(size), C.double(0.0), C.double(360.0))
}
func (cvs *Canvas) FText (x, y float64, s string) {
    str := iup.NewCS(s)
    defer iup.FreeCS(str)
    C.cdfCanvasText(cvs.p, C.double(x), C.double(y), (*C.char)(str))
}
func (cvs *Canvas) TextAlignment (alignment int) {
    C.cdCanvasTextAlignment(cvs.p, C.int(alignment))
}
func (cvs *Canvas) TextOrientation(angle float64) {
    C.cdCanvasTextOrientation(cvs.p, C.double(angle))
}
func (cvs *Canvas) VectorText (x, y int, s string) {
    str := iup.NewCS(s)
    defer iup.FreeCS(str)
    C.cdCanvasVectorText(cvs.p, C.int(x), C.int(y), (*C.char)(str))
}
func (cvs *Canvas) FVectorText (x, y float64, s string) {
    str := iup.NewCS(s)
    defer iup.FreeCS(str)
    C.cdfCanvasVectorText(cvs.p, C.double(x), C.double(y), (*C.char)(str))
}
func (cvs *Canvas) VectorFont (fn string) {
    str := iup.NewCS(fn)
    defer iup.FreeCS(str)
    rtn := C.cdCanvasVectorFont(cvs.p, (*C.char)(str))
    fmt.Println(rtn)
}
func (cvs *Canvas) VectorTextDirection(x1, y1, x2, y2 int) {
    C.cdCanvasVectorTextDirection(cvs.p, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (cvs *Canvas) FVectorTextDirection(x1, y1, x2, y2 float64) {
    C.cdfCanvasVectorTextDirection(cvs.p, C.double(x1), C.double(y1), C.double(x2), C.double(y2))
}
func (cvs *Canvas) VectorFontSize (x, y float64) {
    C.cdCanvasVectorFontSize(cvs.p, C.double(x), C.double(y))
}

// attributes
func (cvs *Canvas) Foreground (color int) int {
    return int(C.cdCanvasForeground(cvs.p, C.long(color)))
}
func (cvs *Canvas) Background (color int) int {
    return int(C.cdCanvasBackground(cvs.p, C.long(color)))
}
func (cvs *Canvas) BackOpacity (opacity int) int {
    return int(C.cdCanvasBackOpacity(cvs.p, C.int(opacity)))
}
func (cvs *Canvas) InteriorStyle (style int) int {
    return int(C.cdCanvasInteriorStyle(cvs.p, C.int(style)))
}
func (cvs *Canvas) Hatch (style int) int {
    return int(C.cdCanvasHatch(cvs.p, C.int(style)))
}
func (cvs *Canvas) WriteMode (mode int) int {
    return int(C.cdCanvasWriteMode(cvs.p, C.int(mode)))
}
func (cvs *Canvas) LineStyle (style int) int {
    return int(C.cdCanvasLineStyle(cvs.p, C.int(style)))
}
func (cvs *Canvas) LineWidth (width int) int {
    return int(C.cdCanvasLineWidth(cvs.p, C.int(width)))
}
func (cvs *Canvas) Font (typeface string, style, size int) int {
    tf := iup.NewCS(typeface)
    defer iup.FreeCS(tf)
    return int(C.cdCanvasFont(cvs.p, (*C.char)(tf), C.int(style), C.int(size)))
}

// color
func EncodeColor(r,g,b int) int {
    return int(C.cdEncodeColor(C.uchar(r), C.uchar(g), C.uchar(b)))
}
func DecodeColor(color int) (r, g, b int) {
    var cr, cg, cb *C.uchar
    C.cdDecodeColor(C.long(color), cr, cg, cb)
    return int(*cr), int(*cg), int(*cb)
}
func EncodeAlpha(color int, alpha int) int {
    return int(C.cdEncodeAlpha(C.long(color), C.uchar(alpha)))
}
func DecodeAlpha(color int) int {
    return int(C.cdDecodeAlpha(C.long(color)))
}
