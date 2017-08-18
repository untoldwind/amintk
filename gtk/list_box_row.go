package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// ListBoxRow is a representation of GTK's GtkListBoxRow.
type ListBoxRow struct {
	Bin
}

// native returns a pointer to the underlying GtkListBoxRow.
func (v *ListBoxRow) native() *C.GtkListBoxRow {
	if v == nil {
		return nil
	}
	return (*C.GtkListBoxRow)(v.Native())
}

func ListBoxRowNew() *ListBoxRow {
	c := C.gtk_list_box_row_new()
	return wrapListBoxRow(unsafe.Pointer(c))
}

func wrapListBoxRow(p unsafe.Pointer) *ListBoxRow {
	if bin := wrapBin(p); bin != nil {
		return &ListBoxRow{Bin: *bin}
	}
	return nil
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *ListBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.native())
	return int(c)
}
