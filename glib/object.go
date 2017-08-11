package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "object.go.h"
import "C"
import (
	"runtime"
	"unsafe"
)

// Object is a representation of GLib's GObject.
type Object struct {
	GObject *C.GObject
}

// newObject creates a new Object from a GObject pointer.
func newObject(p *C.GObject) *Object {
	obj := &Object{GObject: p}

	if C.g_object_is_floating(C.gpointer(obj.GObject)) != 0 {
		C.g_object_ref_sink(C.gpointer(obj.GObject))
	}

	runtime.SetFinalizer(obj, finalizeObject)

	return obj
}

func finalizeObject(obj *Object) {
	C.g_object_unref(C.gpointer(obj.GObject))
}

func NewObject(gType Type) *Object {
	return newObject(C._g_object_new(C.GType(gType)))
}

// Ref is a wrapper around g_object_ref().
func (v *Object) Ref() {
	if v == nil {
		return
	}
	C.g_object_ref(C.gpointer(v.GObject))
}

// Unref is a wrapper around g_object_unref().
func (v *Object) Unref() {
	if v == nil {
		return
	}
	C.g_object_unref(C.gpointer(v.GObject))
}

// IsFloating is a wrapper around g_object_is_floating().
func (v *Object) IsFloating() bool {
	if v == nil {
		return false
	}
	c := C.g_object_is_floating(C.gpointer(v.GObject))
	return c != 0
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	if v == nil {
		return
	}
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

// GetPropertyType returns the Type of a property of the underlying GObject.
// If the property is missing it will return TYPE_INVALID and an error.
func (v *Object) GetPropertyType(name string) Type {
	if v == nil {
		return TYPE_INVALID
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	paramSpec := C.g_object_class_find_property(C._g_object_get_class(v.GObject), (*C.gchar)(cstr))
	if paramSpec == nil {
		return TYPE_INVALID
	}
	return Type(paramSpec.value_type)
}

// GetProperty is a wrapper around g_object_get_property().
func (v *Object) GetProperty(name string) interface{} {
	if v == nil {
		return nil
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	t := v.GetPropertyType(name)

	p := ValueInit(t)
	if p == nil {
		return nil
	}
	C.g_object_get_property(v.GObject, (*C.gchar)(cstr), p.GValue)
	return p.GoValue()
}

// SetProperty is a wrapper around g_object_set_property().
func (v *Object) SetProperty(name string, value interface{}) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	if _, ok := value.(Object); ok {
		value = value.(Object).GObject
	}

	if p := GValue(value); p != nil {
		C.g_object_set_property(v.GObject, (*C.gchar)(cstr), p.GValue)
	}
}
