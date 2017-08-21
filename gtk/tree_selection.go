package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// TreeSelection is a representation of GTK's GtkTreeSelection.
type TreeSelection struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTreeView.
func (v *TreeSelection) native() *C.GtkTreeSelection {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeSelection)(v.Native())
}

func wrapTreeSelection(p unsafe.Pointer) *TreeSelection {
	if obj := glib.WrapObject(p); obj != nil {
		return &TreeSelection{Object: obj}
	}
	return nil
}

// GetSelected is a wrapper around gtk_tree_selection_get_selected().
func (v *TreeSelection) GetSelected() (model ITreeModel, iter *TreeIter, ok bool) {
	var cmodel *C.GtkTreeModel
	var citer C.GtkTreeIter
	c := C.gtk_tree_selection_get_selected(v.native(),
		&cmodel, &citer)
	model = wrapTreeModel(glib.WrapObject(unsafe.Pointer(cmodel)))
	iter = &TreeIter{citer}
	ok = gobool(c)
	return
}
