#ifndef __GTK_LIST_STORE_GO_H_
#define __GTK_LIST_STORE_GO_H_

#include <stdlib.h>
#include <gtk/gtk.h>

static GType *
alloc_types(int n) {
	return ((GType *)g_new0(GType, n));
}

static void
set_type(GType *types, int n, GType t)
{
	types[n] = t;
}
#endif
