package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// TreeView is a representation of GTK's GtkTreeView.
type TreeView struct {
	Container
}

// native returns a pointer to the underlying GtkTreeView.
func (v *TreeView) native() *C.GtkTreeView {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeView)(v.Native())
}

func wrapTreeView(p unsafe.Pointer) *TreeView {
	if container := wrapContainer(p); container != nil {
		return &TreeView{Container: *container}
	}
	return nil
}
