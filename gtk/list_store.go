package gtk

// #cgo pkg-config: gtk+-3.0
// #include "list_store.go.h"
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// ListStore is a representation of GTK's GtkListStore.
type ListStore struct {
	*glib.Object

	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkListStore.
func (v *ListStore) native() *C.GtkListStore {
	if v == nil {
		return nil
	}
	return (*C.GtkListStore)(v.Native())
}

// ListStoreNew is a wrapper around gtk_list_store_newv().
func ListStoreNew(types ...glib.Type) *ListStore {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_list_store_newv(C.gint(len(types)), gtypes)

	return wrapListStore(unsafe.Pointer(c))
}

func (v *ListStore) toTreeModel() *C.GtkTreeModel {
	return (*C.GtkTreeModel)(v.Native())
}

func wrapListStore(p unsafe.Pointer) *ListStore {
	if obj := glib.WrapObject(p); obj != nil {
		tm := wrapTreeModel(obj)
		return &ListStore{Object: obj, TreeModel: *tm}
	}
	return nil
}
