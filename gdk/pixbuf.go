package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib.Object
}

func WrapPixbuf(p unsafe.Pointer) *Pixbuf {
	if obj := glib.WrapObject(p); obj != nil {
		return &Pixbuf{Object: obj}
	}
	return nil
}
