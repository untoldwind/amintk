package gtk

// #cgo pkg-config: gtk+-3.0
// #include "container.go.h"
import "C"
import "unsafe"

// Container is a representation of GTK's GtkContainer.
type Container struct {
	Widget
}

// native returns a pointer to the underlying GtkContainer.
func (v *Container) native() *C.GtkContainer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkContainer)(p)
}

// Add is a wrapper around gtk_container_add().
func (v *Container) Add(w IWidget) {
	C.gtk_container_add(v.native(), w.toWidget())
}

// Remove is a wrapper around gtk_container_remove().
func (v *Container) Remove(w IWidget) {
	C.gtk_container_remove(v.native(), w.toWidget())
}
