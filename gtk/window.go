package gtk

// #cgo pkg-config: gtk+-3.0
// #include "window.go.h"
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// WindowType is a representation of GTK's GtkWindowType.
type WindowType int

const (
	WindowToplevel WindowType = C.GTK_WINDOW_TOPLEVEL
	WindowPopup    WindowType = C.GTK_WINDOW_POPUP
)

// Window is a representation of GTK's GtkWindow.
type Window struct {
	Bin
}

// native returns a pointer to the underlying GtkWindow.
func (v *Window) native() *C.GtkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkWindow)(p)
}

// WindowNew is a wrapper around gtk_window_new().
func WindowNew(t WindowType) *Window {
	c := C.gtk_window_new(C.GtkWindowType(t))
	if c == nil {
		return nil
	}
	return wrapWindow(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapWindow(obj *glib.Object) *Window {
	return &Window{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// SetTitle is a wrapper around gtk_window_set_title().
func (v *Window) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_title(v.native(), (*C.gchar)(cstr))
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(v.native(), gbool(resizable))
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (v *Window) GetResizable() bool {
	c := C.gtk_window_get_resizable(v.native())
	return gobool(c)
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (v *Window) ActivateFocus() bool {
	c := C.gtk_window_activate_focus(v.native())
	return gobool(c)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (v *Window) ActivateDefault() bool {
	c := C.gtk_window_activate_default(v.native())
	return gobool(c)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(v.native(), gbool(modal))
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (v *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(v.native(), C.gint(width), C.gint(height))
}

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (v *Window) GetDefaultSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_default_size(v.native(), &w, &h)
	return int(w), int(h)
}