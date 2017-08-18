package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// FileChooserButton is a representation of GTK's GtkFileChooserButton.
type FileChooserButton struct {
	Box

	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserButton.
func (v *FileChooserButton) native() *C.GtkFileChooserButton {
	if v == nil {
		return nil
	}
	return (*C.GtkFileChooserButton)(v.Native())
}

// FileChooserButtonNew is a wrapper around gtk_file_chooser_button_new().
func FileChooserButtonNew(title string, action FileChooserAction) *FileChooserButton {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_button_new((*C.gchar)(cstr),
		(C.GtkFileChooserAction)(action))
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapFileChooserButton(obj)
}

func wrapFileChooserButton(obj *glib.Object) *FileChooserButton {
	fc := wrapFileChooser(obj)
	return &FileChooserButton{Box: Box{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}, FileChooser: *fc}
}
