package gtk

// #cgo pkg-config: gtk+-3.0
// #include "label.go.h"
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Label is a representation of GTK's GtkLabel.
type Label struct {
	Widget
}

// native returns a pointer to the underlying GtkLabel.
func (v *Label) native() *C.GtkLabel {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkLabel)(p)
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) *Label {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new((*C.gchar)(cstr))
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapLabel(obj)
}

func wrapLabel(obj *glib.Object) *Label {
	return &Label{Widget{glib.InitiallyUnowned{Object: obj}}}
}
