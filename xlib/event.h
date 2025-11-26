#ifndef XLIB_EVENT_H
#define XLIB_EVENT_H

#include <X11/Xlib.h>

static inline int GetEventType(XEvent *event) {
    return event->type;
}

static inline XKeyEvent *GetKeyEvent(XEvent *event) {
    return &event->xkey;
}

static inline XButtonEvent *GetButtonEvent(XEvent *event) {
    return &event->xbutton;
}

static inline XMotionEvent *GetMotionEvent(XEvent *event) {
    return &event->xmotion;
}

static inline XMapRequestEvent *GetMapRequestEvent(XEvent *event) {
    return &event->xmaprequest;
}

static inline XConfigureRequestEvent *GetConfigureRequestEvent(XEvent *event) {
    return &event->xconfigurerequest;
}

#endif
