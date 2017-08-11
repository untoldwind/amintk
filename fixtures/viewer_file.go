package fixtures

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "viewer_file.go.h"
import "C"
import "github.com/untoldwind/amintk/glib"

func ViewerFileGetType() glib.Type {
	return glib.Type(C.viewer_file_get_type())
}
