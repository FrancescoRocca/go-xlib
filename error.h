#ifndef XLIB_ERROR_H
#define XLIB_ERROR_H

#include <X11/Xlib.h>

static int errorCode = 0;

static int errorHandler(Display *d, XErrorEvent *event) {
    errorCode = event->error_code;
    return 0;
}

static void SetErrorHandler() {
    XSetErrorHandler(errorHandler);
}

static void ResetErrorCode() {
    errorCode = 0;
}

static int GetErrorCode() {
    return errorCode;
}

#endif
