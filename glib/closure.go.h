#ifndef __GLIB_CLOSURE_GO_H__
#define __GLIB_CLOSURE_GO_H__

#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>

extern void	goMarshal(GClosure *, GValue *, guint, GValue *, gpointer, GValue *);

static GClosure *
_g_closure_new()
{
	GClosure	*closure;

	closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_marshal(closure, (GClosureMarshal)(goMarshal));
	return (closure);
}

extern void	removeClosure(gpointer, GClosure *);

static void
_g_closure_add_finalize_notifier(GClosure *closure)
{
	g_closure_add_finalize_notifier(closure, NULL, removeClosure);
}

#endif
