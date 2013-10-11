package kbd

/*
#cgo CFLAGS: -Qunused-arguments
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>
#include <stdio.h>

void keyevt(int keycode, bool isdown) {
    CGEventRef evt;
    evt = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)keycode, isdown);
    CGEventPost(kCGSessionEventTap, evt);
}
*/
import "C"

var KEYS = map[string] int32{
    "KEY_UP": 126,    // UP
    "KEY_DOWN": 125,  // DOWN
    "KEY_LEFT": 123,  // LEFT
    "KEY_RIGHT": 124, // RIGHT

    "KEY_L": 7,       // x
    "KEY_R": 2,       // d

    "KEY_A": 27,      // -
    "KEY_B": 24,      // =

    "KEY_X": 6,       // z
    "KEY_Y": 12,      // q

    "KEY_START": 44,  // /
    "KEY_SELECT": 11, // b
}

func KeyDown(key string) {
    C.keyevt(C.int(KEYS[key]), C.bool(true))
}

func KeyUp(key string) {
    C.keyevt(C.int(KEYS[key]), C.bool(false))
}
