package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

// TreeIter is a representation of GTK's GtkTreeIter.
type TreeIter struct {
	GtkTreeIter C.GtkTreeIter
}
