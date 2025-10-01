#include "event.h"

int GetEventType(XEvent *event) {
  return event->type;
}
