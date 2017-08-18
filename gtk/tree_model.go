package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"github.com/untoldwind/amintk/glib"
)

// TreeModel is a representation of GTK's GtkTreeModel GInterface.
type TreeModel struct {
	*glib.Object
}

type ITreeModel interface {
	toTreeModel() *C.GtkTreeModel
}

// native returns a pointer to the underlying GObject as a GtkTreeModel.
func (v *TreeModel) native() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeModel)(v.Native())
}

func (v *TreeModel) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return v.native()
}

func wrapTreeModel(obj *glib.Object) *TreeModel {
	return &TreeModel{Object: obj}
}
