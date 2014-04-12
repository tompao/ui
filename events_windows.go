// 25 march 2014

package ui

// Virtual key codes.
const (
	// from winuser.h
	_VK_LBUTTON             = 0x01
	_VK_RBUTTON             = 0x02
	_VK_CANCEL              = 0x03
	_VK_MBUTTON             = 0x04
	_VK_XBUTTON1            = 0x05
	_VK_XBUTTON2            = 0x06
	_VK_BACK                = 0x08
	_VK_TAB                 = 0x09
	_VK_CLEAR               = 0x0C
	_VK_RETURN              = 0x0D
	_VK_SHIFT               = 0x10
	_VK_CONTROL             = 0x11
	_VK_MENU                = 0x12
	_VK_PAUSE               = 0x13
	_VK_CAPITAL             = 0x14
	_VK_KANA                = 0x15
	_VK_HANGEUL             = 0x15
	_VK_HANGUL              = 0x15
	_VK_JUNJA               = 0x17
	_VK_FINAL               = 0x18
	_VK_HANJA               = 0x19
	_VK_KANJI               = 0x19
	_VK_ESCAPE              = 0x1B
	_VK_CONVERT             = 0x1C
	_VK_NONCONVERT          = 0x1D
	_VK_ACCEPT              = 0x1E
	_VK_MODECHANGE          = 0x1F
	_VK_SPACE               = 0x20
	_VK_PRIOR               = 0x21
	_VK_NEXT                = 0x22
	_VK_END                 = 0x23
	_VK_HOME                = 0x24
	_VK_LEFT                = 0x25
	_VK_UP                  = 0x26
	_VK_RIGHT               = 0x27
	_VK_DOWN                = 0x28
	_VK_SELECT              = 0x29
	_VK_PRINT               = 0x2A
	_VK_EXECUTE             = 0x2B
	_VK_SNAPSHOT            = 0x2C
	_VK_INSERT              = 0x2D
	_VK_DELETE              = 0x2E
	_VK_HELP                = 0x2F
	_VK_LWIN                = 0x5B
	_VK_RWIN                = 0x5C
	_VK_APPS                = 0x5D
	_VK_SLEEP               = 0x5F
	_VK_NUMPAD0             = 0x60
	_VK_NUMPAD1             = 0x61
	_VK_NUMPAD2             = 0x62
	_VK_NUMPAD3             = 0x63
	_VK_NUMPAD4             = 0x64
	_VK_NUMPAD5             = 0x65
	_VK_NUMPAD6             = 0x66
	_VK_NUMPAD7             = 0x67
	_VK_NUMPAD8             = 0x68
	_VK_NUMPAD9             = 0x69
	_VK_MULTIPLY            = 0x6A
	_VK_ADD                 = 0x6B
	_VK_SEPARATOR           = 0x6C
	_VK_SUBTRACT            = 0x6D
	_VK_DECIMAL             = 0x6E
	_VK_DIVIDE              = 0x6F
	_VK_F1                  = 0x70
	_VK_F2                  = 0x71
	_VK_F3                  = 0x72
	_VK_F4                  = 0x73
	_VK_F5                  = 0x74
	_VK_F6                  = 0x75
	_VK_F7                  = 0x76
	_VK_F8                  = 0x77
	_VK_F9                  = 0x78
	_VK_F10                 = 0x79
	_VK_F11                 = 0x7A
	_VK_F12                 = 0x7B
	_VK_F13                 = 0x7C
	_VK_F14                 = 0x7D
	_VK_F15                 = 0x7E
	_VK_F16                 = 0x7F
	_VK_F17                 = 0x80
	_VK_F18                 = 0x81
	_VK_F19                 = 0x82
	_VK_F20                 = 0x83
	_VK_F21                 = 0x84
	_VK_F22                 = 0x85
	_VK_F23                 = 0x86
	_VK_F24                 = 0x87
	_VK_NUMLOCK             = 0x90
	_VK_SCROLL              = 0x91
	_VK_OEM_NEC_EQUAL       = 0x92
	_VK_OEM_FJ_JISHO        = 0x92
	_VK_OEM_FJ_MASSHOU      = 0x93
	_VK_OEM_FJ_TOUROKU      = 0x94
	_VK_OEM_FJ_LOYA         = 0x95
	_VK_OEM_FJ_ROYA         = 0x96
	_VK_LSHIFT              = 0xA0
	_VK_RSHIFT              = 0xA1
	_VK_LCONTROL            = 0xA2
	_VK_RCONTROL            = 0xA3
	_VK_LMENU               = 0xA4
	_VK_RMENU               = 0xA5
	_VK_BROWSER_BACK        = 0xA6
	_VK_BROWSER_FORWARD     = 0xA7
	_VK_BROWSER_REFRESH     = 0xA8
	_VK_BROWSER_STOP        = 0xA9
	_VK_BROWSER_SEARCH      = 0xAA
	_VK_BROWSER_FAVORITES   = 0xAB
	_VK_BROWSER_HOME        = 0xAC
	_VK_VOLUME_MUTE         = 0xAD
	_VK_VOLUME_DOWN         = 0xAE
	_VK_VOLUME_UP           = 0xAF
	_VK_MEDIA_NEXT_TRACK    = 0xB0
	_VK_MEDIA_PREV_TRACK    = 0xB1
	_VK_MEDIA_STOP          = 0xB2
	_VK_MEDIA_PLAY_PAUSE    = 0xB3
	_VK_LAUNCH_MAIL         = 0xB4
	_VK_LAUNCH_MEDIA_SELECT = 0xB5
	_VK_LAUNCH_APP1         = 0xB6
	_VK_LAUNCH_APP2         = 0xB7
	_VK_OEM_1               = 0xBA
	_VK_OEM_PLUS            = 0xBB
	_VK_OEM_COMMA           = 0xBC
	_VK_OEM_MINUS           = 0xBD
	_VK_OEM_PERIOD          = 0xBE
	_VK_OEM_2               = 0xBF
	_VK_OEM_3               = 0xC0
	_VK_OEM_4               = 0xDB
	_VK_OEM_5               = 0xDC
	_VK_OEM_6               = 0xDD
	_VK_OEM_7               = 0xDE
	_VK_OEM_8               = 0xDF
	_VK_OEM_AX              = 0xE1
	_VK_OEM_102             = 0xE2
	_VK_ICO_HELP            = 0xE3
	_VK_ICO_00              = 0xE4
	_VK_PROCESSKEY          = 0xE5
	_VK_ICO_CLEAR           = 0xE6
	_VK_PACKET              = 0xE7
	_VK_OEM_RESET           = 0xE9
	_VK_OEM_JUMP            = 0xEA
	_VK_OEM_PA1             = 0xEB
	_VK_OEM_PA2             = 0xEC
	_VK_OEM_PA3             = 0xED
	_VK_OEM_WSCTRL          = 0xEE
	_VK_OEM_CUSEL           = 0xEF
	_VK_OEM_ATTN            = 0xF0
	_VK_OEM_FINISH          = 0xF1
	_VK_OEM_COPY            = 0xF2
	_VK_OEM_AUTO            = 0xF3
	_VK_OEM_ENLW            = 0xF4
	_VK_OEM_BACKTAB         = 0xF5
	_VK_ATTN                = 0xF6
	_VK_CRSEL               = 0xF7
	_VK_EXSEL               = 0xF8
	_VK_EREOF               = 0xF9
	_VK_PLAY                = 0xFA
	_VK_ZOOM                = 0xFB
	_VK_NONAME              = 0xFC
	_VK_PA1                 = 0xFD
	_VK_OEM_CLEAR           = 0xFE
)

// Mouse event modifier masks.
const (
	// from winuser.h
	_MK_LBUTTON  = 0x0001
	_MK_RBUTTON  = 0x0002
	_MK_SHIFT    = 0x0004
	_MK_CONTROL  = 0x0008
	_MK_MBUTTON  = 0x0010
	_MK_XBUTTON1 = 0x0020
	_MK_XBUTTON2 = 0x0040
)

// Window mouse event messages.
const (
	_WM_MOUSEACTIVATE = 0x0021

	// from winuser.h
	_WM_MOUSEFIRST    = 0x0200
	_WM_MOUSEMOVE     = 0x0200
	_WM_LBUTTONDOWN   = 0x0201
	_WM_LBUTTONUP     = 0x0202
	_WM_LBUTTONDBLCLK = 0x0203
	_WM_RBUTTONDOWN   = 0x0204
	_WM_RBUTTONUP     = 0x0205
	_WM_RBUTTONDBLCLK = 0x0206
	_WM_MBUTTONDOWN   = 0x0207
	_WM_MBUTTONUP     = 0x0208
	_WM_MBUTTONDBLCLK = 0x0209
	_WM_MOUSEWHEEL    = 0x020A
	_WM_XBUTTONDOWN   = 0x020B
	_WM_XBUTTONUP     = 0x020C
	_WM_XBUTTONDBLCLK = 0x020D
)

// Window keyboard event messages and related constants.
const (
	// from winuser.h
	_WM_KEYDOWN     = 0x0100
	_WM_KEYUP       = 0x0101
	_WM_CHAR        = 0x0102
	_WM_DEADCHAR    = 0x0103
	_WM_SYSKEYDOWN  = 0x0104
	_WM_SYSKEYUP    = 0x0105
	_WM_SYSCHAR     = 0x0106
	_WM_SYSDEADCHAR = 0x0107
	_WM_UNICHAR     = 0x0109
	_UNICODE_NOCHAR = 0xFFFF // used by _WM_UNICHAR
)
