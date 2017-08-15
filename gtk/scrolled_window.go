package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

// native returns a pointer to the underlying GtkScrolledWindow.
func (v *ScrolledWindow) native() *C.GtkScrolledWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkScrolledWindow)(p)
}

// ScrolledWindowNew() is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	c := C.gtk_scrolled_window_new(hadjustment.native(),
		vadjustment.native())
	return wrapScrolledWindow(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapScrolledWindow(obj *glib.Object) *ScrolledWindow {
	return &ScrolledWindow{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}
