package kbd

/*
#cgo CFLAGS: -Qunused-arguments
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>
#include <Carbon/Carbon.h>

void keyevt(int keycode, bool isdown) {
    CGEventRef evt;
    evt = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)keycode, isdown);
    CGEventPost(kCGSessionEventTap, evt);
}
*/
import "C"

var KEYS = map[string] int32{
    "KEY_UP"    : C.kVK_UpArrow,
    "KEY_DOWN"  : C.kVK_DownArrow,
    "KEY_LEFT"  : C.kVK_LeftArrow,
    "KEY_RIGHT" : C.kVK_RightArrow,

    "KEY_L"     : C.kVK_ANSI_Q,
    "KEY_R"     : C.kVK_ANSI_E,

    "KEY_A"     : C.kVK_ANSI_A,
    "KEY_B"     : C.kVK_ANSI_S,

    "KEY_X"     : C.kVK_ANSI_Z,
    "KEY_Y"     : C.kVK_ANSI_X,

    "KEY_START" : C.kVK_Return,
    "KEY_SELECT": C.kVK_Space,
}

func KeyDown(key string) {
    C.keyevt(C.int(KEYS[key]), C.bool(true))
}

func KeyUp(key string) {
    C.keyevt(C.int(KEYS[key]), C.bool(false))
}
