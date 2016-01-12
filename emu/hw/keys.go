package hw

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCANCODE_A = sdl.SCANCODE_A
	SCANCODE_B = sdl.SCANCODE_B
	SCANCODE_C = sdl.SCANCODE_C
	SCANCODE_D = sdl.SCANCODE_D
	SCANCODE_E = sdl.SCANCODE_E
	SCANCODE_F = sdl.SCANCODE_F
	SCANCODE_G = sdl.SCANCODE_G
	SCANCODE_H = sdl.SCANCODE_H
	SCANCODE_I = sdl.SCANCODE_I
	SCANCODE_J = sdl.SCANCODE_J
	SCANCODE_K = sdl.SCANCODE_K
	SCANCODE_L = sdl.SCANCODE_L
	SCANCODE_M = sdl.SCANCODE_M
	SCANCODE_N = sdl.SCANCODE_N
	SCANCODE_O = sdl.SCANCODE_O
	SCANCODE_P = sdl.SCANCODE_P
	SCANCODE_Q = sdl.SCANCODE_Q
	SCANCODE_R = sdl.SCANCODE_R
	SCANCODE_S = sdl.SCANCODE_S
	SCANCODE_T = sdl.SCANCODE_T
	SCANCODE_U = sdl.SCANCODE_U
	SCANCODE_V = sdl.SCANCODE_V
	SCANCODE_W = sdl.SCANCODE_W
	SCANCODE_X = sdl.SCANCODE_X
	SCANCODE_Y = sdl.SCANCODE_Y
	SCANCODE_Z = sdl.SCANCODE_Z

	SCANCODE_1 = sdl.SCANCODE_1
	SCANCODE_2 = sdl.SCANCODE_2
	SCANCODE_3 = sdl.SCANCODE_3
	SCANCODE_4 = sdl.SCANCODE_4
	SCANCODE_5 = sdl.SCANCODE_5
	SCANCODE_6 = sdl.SCANCODE_6
	SCANCODE_7 = sdl.SCANCODE_7
	SCANCODE_8 = sdl.SCANCODE_8
	SCANCODE_9 = sdl.SCANCODE_9
	SCANCODE_0 = sdl.SCANCODE_0

	SCANCODE_RETURN    = sdl.SCANCODE_RETURN
	SCANCODE_ESCAPE    = sdl.SCANCODE_ESCAPE
	SCANCODE_BACKSPACE = sdl.SCANCODE_BACKSPACE
	SCANCODE_TAB       = sdl.SCANCODE_TAB
	SCANCODE_SPACE     = sdl.SCANCODE_SPACE

	SCANCODE_MINUS        = sdl.SCANCODE_MINUS
	SCANCODE_EQUALS       = sdl.SCANCODE_EQUALS
	SCANCODE_LEFTBRACKET  = sdl.SCANCODE_LEFTBRACKET
	SCANCODE_RIGHTBRACKET = sdl.SCANCODE_RIGHTBRACKET
	SCANCODE_BACKSLASH    = sdl.SCANCODE_BACKSLASH
	SCANCODE_NONUSHASH    = sdl.SCANCODE_NONUSHASH
	SCANCODE_SEMICOLON    = sdl.SCANCODE_SEMICOLON
	SCANCODE_APOSTROPHE   = sdl.SCANCODE_APOSTROPHE
	SCANCODE_GRAVE        = sdl.SCANCODE_GRAVE
	SCANCODE_COMMA        = sdl.SCANCODE_COMMA
	SCANCODE_PERIOD       = sdl.SCANCODE_PERIOD
	SCANCODE_SLASH        = sdl.SCANCODE_SLASH
	SCANCODE_CAPSLOCK     = sdl.SCANCODE_CAPSLOCK
	SCANCODE_F1           = sdl.SCANCODE_F1
	SCANCODE_F2           = sdl.SCANCODE_F2
	SCANCODE_F3           = sdl.SCANCODE_F3
	SCANCODE_F4           = sdl.SCANCODE_F4
	SCANCODE_F5           = sdl.SCANCODE_F5
	SCANCODE_F6           = sdl.SCANCODE_F6
	SCANCODE_F7           = sdl.SCANCODE_F7
	SCANCODE_F8           = sdl.SCANCODE_F8
	SCANCODE_F9           = sdl.SCANCODE_F9
	SCANCODE_F10          = sdl.SCANCODE_F10
	SCANCODE_F11          = sdl.SCANCODE_F11
	SCANCODE_F12          = sdl.SCANCODE_F12
	SCANCODE_PRINTSCREEN  = sdl.SCANCODE_PRINTSCREEN
	SCANCODE_SCROLLLOCK   = sdl.SCANCODE_SCROLLLOCK
	SCANCODE_PAUSE        = sdl.SCANCODE_PAUSE
	SCANCODE_INSERT       = sdl.SCANCODE_INSERT
	SCANCODE_HOME         = sdl.SCANCODE_HOME
	SCANCODE_PAGEUP       = sdl.SCANCODE_PAGEUP
	SCANCODE_DELETE       = sdl.SCANCODE_DELETE
	SCANCODE_END          = sdl.SCANCODE_END
	SCANCODE_PAGEDOWN     = sdl.SCANCODE_PAGEDOWN
	SCANCODE_RIGHT        = sdl.SCANCODE_RIGHT
	SCANCODE_LEFT         = sdl.SCANCODE_LEFT
	SCANCODE_DOWN         = sdl.SCANCODE_DOWN
	SCANCODE_UP           = sdl.SCANCODE_UP

	SCANCODE_NUMLOCKCLEAR = sdl.SCANCODE_NUMLOCKCLEAR
	SCANCODE_KP_DIVIDE    = sdl.SCANCODE_KP_DIVIDE
	SCANCODE_KP_MULTIPLY  = sdl.SCANCODE_KP_MULTIPLY
	SCANCODE_KP_MINUS     = sdl.SCANCODE_KP_MINUS
	SCANCODE_KP_PLUS      = sdl.SCANCODE_KP_PLUS
	SCANCODE_KP_ENTER     = sdl.SCANCODE_KP_ENTER
	SCANCODE_KP_1         = sdl.SCANCODE_KP_1
	SCANCODE_KP_2         = sdl.SCANCODE_KP_2
	SCANCODE_KP_3         = sdl.SCANCODE_KP_3
	SCANCODE_KP_4         = sdl.SCANCODE_KP_4
	SCANCODE_KP_5         = sdl.SCANCODE_KP_5
	SCANCODE_KP_6         = sdl.SCANCODE_KP_6
	SCANCODE_KP_7         = sdl.SCANCODE_KP_7
	SCANCODE_KP_8         = sdl.SCANCODE_KP_8
	SCANCODE_KP_9         = sdl.SCANCODE_KP_9
	SCANCODE_KP_0         = sdl.SCANCODE_KP_0
	SCANCODE_KP_PERIOD    = sdl.SCANCODE_KP_PERIOD

	SCANCODE_NONUSBACKSLASH = sdl.SCANCODE_NONUSBACKSLASH
	SCANCODE_APPLICATION    = sdl.SCANCODE_APPLICATION
	SCANCODE_POWER          = sdl.SCANCODE_POWER
	SCANCODE_KP_EQUALS      = sdl.SCANCODE_KP_EQUALS
	SCANCODE_F13            = sdl.SCANCODE_F13
	SCANCODE_F14            = sdl.SCANCODE_F14
	SCANCODE_F15            = sdl.SCANCODE_F15
	SCANCODE_F16            = sdl.SCANCODE_F16
	SCANCODE_F17            = sdl.SCANCODE_F17
	SCANCODE_F18            = sdl.SCANCODE_F18
	SCANCODE_F19            = sdl.SCANCODE_F19
	SCANCODE_F20            = sdl.SCANCODE_F20
	SCANCODE_F21            = sdl.SCANCODE_F21
	SCANCODE_F22            = sdl.SCANCODE_F22
	SCANCODE_F23            = sdl.SCANCODE_F23
	SCANCODE_F24            = sdl.SCANCODE_F24
	SCANCODE_EXECUTE        = sdl.SCANCODE_EXECUTE
	SCANCODE_HELP           = sdl.SCANCODE_HELP
	SCANCODE_MENU           = sdl.SCANCODE_MENU
	SCANCODE_SELECT         = sdl.SCANCODE_SELECT
	SCANCODE_STOP           = sdl.SCANCODE_STOP
	SCANCODE_AGAIN          = sdl.SCANCODE_AGAIN
	SCANCODE_UNDO           = sdl.SCANCODE_UNDO
	SCANCODE_CUT            = sdl.SCANCODE_CUT
	SCANCODE_COPY           = sdl.SCANCODE_COPY
	SCANCODE_PASTE          = sdl.SCANCODE_PASTE
	SCANCODE_FIND           = sdl.SCANCODE_FIND
	SCANCODE_MUTE           = sdl.SCANCODE_MUTE
	SCANCODE_VOLUMEUP       = sdl.SCANCODE_VOLUMEUP
	SCANCODE_VOLUMEDOWN     = sdl.SCANCODE_VOLUMEDOWN
	SCANCODE_KP_COMMA       = sdl.SCANCODE_KP_COMMA
	SCANCODE_KP_EQUALSAS400 = sdl.SCANCODE_KP_EQUALSAS400

	SCANCODE_INTERNATIONAL1 = sdl.SCANCODE_INTERNATIONAL1
	SCANCODE_INTERNATIONAL2 = sdl.SCANCODE_INTERNATIONAL2
	SCANCODE_INTERNATIONAL3 = sdl.SCANCODE_INTERNATIONAL3
	SCANCODE_INTERNATIONAL4 = sdl.SCANCODE_INTERNATIONAL4
	SCANCODE_INTERNATIONAL5 = sdl.SCANCODE_INTERNATIONAL5
	SCANCODE_INTERNATIONAL6 = sdl.SCANCODE_INTERNATIONAL6
	SCANCODE_INTERNATIONAL7 = sdl.SCANCODE_INTERNATIONAL7
	SCANCODE_INTERNATIONAL8 = sdl.SCANCODE_INTERNATIONAL8
	SCANCODE_INTERNATIONAL9 = sdl.SCANCODE_INTERNATIONAL9
	SCANCODE_LANG1          = sdl.SCANCODE_LANG1
	SCANCODE_LANG2          = sdl.SCANCODE_LANG2
	SCANCODE_LANG3          = sdl.SCANCODE_LANG3
	SCANCODE_LANG4          = sdl.SCANCODE_LANG4
	SCANCODE_LANG5          = sdl.SCANCODE_LANG5
	SCANCODE_LANG6          = sdl.SCANCODE_LANG6
	SCANCODE_LANG7          = sdl.SCANCODE_LANG7
	SCANCODE_LANG8          = sdl.SCANCODE_LANG8
	SCANCODE_LANG9          = sdl.SCANCODE_LANG9

	SCANCODE_ALTERASE   = sdl.SCANCODE_ALTERASE
	SCANCODE_SYSREQ     = sdl.SCANCODE_SYSREQ
	SCANCODE_CANCEL     = sdl.SCANCODE_CANCEL
	SCANCODE_CLEAR      = sdl.SCANCODE_CLEAR
	SCANCODE_PRIOR      = sdl.SCANCODE_PRIOR
	SCANCODE_RETURN2    = sdl.SCANCODE_RETURN2
	SCANCODE_SEPARATOR  = sdl.SCANCODE_SEPARATOR
	SCANCODE_OUT        = sdl.SCANCODE_OUT
	SCANCODE_OPER       = sdl.SCANCODE_OPER
	SCANCODE_CLEARAGAIN = sdl.SCANCODE_CLEARAGAIN
	SCANCODE_CRSEL      = sdl.SCANCODE_CRSEL
	SCANCODE_EXSEL      = sdl.SCANCODE_EXSEL

	SCANCODE_KP_00              = sdl.SCANCODE_KP_00
	SCANCODE_KP_000             = sdl.SCANCODE_KP_000
	SCANCODE_THOUSANDSSEPARATOR = sdl.SCANCODE_THOUSANDSSEPARATOR
	SCANCODE_DECIMALSEPARATOR   = sdl.SCANCODE_DECIMALSEPARATOR
	SCANCODE_CURRENCYUNIT       = sdl.SCANCODE_CURRENCYUNIT
	SCANCODE_CURRENCYSUBUNIT    = sdl.SCANCODE_CURRENCYSUBUNIT
	SCANCODE_KP_LEFTPAREN       = sdl.SCANCODE_KP_LEFTPAREN
	SCANCODE_KP_RIGHTPAREN      = sdl.SCANCODE_KP_RIGHTPAREN
	SCANCODE_KP_LEFTBRACE       = sdl.SCANCODE_KP_LEFTBRACE
	SCANCODE_KP_RIGHTBRACE      = sdl.SCANCODE_KP_RIGHTBRACE
	SCANCODE_KP_TAB             = sdl.SCANCODE_KP_TAB
	SCANCODE_KP_BACKSPACE       = sdl.SCANCODE_KP_BACKSPACE
	SCANCODE_KP_A               = sdl.SCANCODE_KP_A
	SCANCODE_KP_B               = sdl.SCANCODE_KP_B
	SCANCODE_KP_C               = sdl.SCANCODE_KP_C
	SCANCODE_KP_D               = sdl.SCANCODE_KP_D
	SCANCODE_KP_E               = sdl.SCANCODE_KP_E
	SCANCODE_KP_F               = sdl.SCANCODE_KP_F
	SCANCODE_KP_XOR             = sdl.SCANCODE_KP_XOR
	SCANCODE_KP_POWER           = sdl.SCANCODE_KP_POWER
	SCANCODE_KP_PERCENT         = sdl.SCANCODE_KP_PERCENT
	SCANCODE_KP_LESS            = sdl.SCANCODE_KP_LESS
	SCANCODE_KP_GREATER         = sdl.SCANCODE_KP_GREATER
	SCANCODE_KP_AMPERSAND       = sdl.SCANCODE_KP_AMPERSAND
	SCANCODE_KP_DBLAMPERSAND    = sdl.SCANCODE_KP_DBLAMPERSAND
	SCANCODE_KP_VERTICALBAR     = sdl.SCANCODE_KP_VERTICALBAR
	SCANCODE_KP_DBLVERTICALBAR  = sdl.SCANCODE_KP_DBLVERTICALBAR
	SCANCODE_KP_COLON           = sdl.SCANCODE_KP_COLON
	SCANCODE_KP_HASH            = sdl.SCANCODE_KP_HASH
	SCANCODE_KP_SPACE           = sdl.SCANCODE_KP_SPACE
	SCANCODE_KP_AT              = sdl.SCANCODE_KP_AT
	SCANCODE_KP_EXCLAM          = sdl.SCANCODE_KP_EXCLAM
	SCANCODE_KP_MEMSTORE        = sdl.SCANCODE_KP_MEMSTORE
	SCANCODE_KP_MEMRECALL       = sdl.SCANCODE_KP_MEMRECALL
	SCANCODE_KP_MEMCLEAR        = sdl.SCANCODE_KP_MEMCLEAR
	SCANCODE_KP_MEMADD          = sdl.SCANCODE_KP_MEMADD
	SCANCODE_KP_MEMSUBTRACT     = sdl.SCANCODE_KP_MEMSUBTRACT
	SCANCODE_KP_MEMMULTIPLY     = sdl.SCANCODE_KP_MEMMULTIPLY
	SCANCODE_KP_MEMDIVIDE       = sdl.SCANCODE_KP_MEMDIVIDE
	SCANCODE_KP_PLUSMINUS       = sdl.SCANCODE_KP_PLUSMINUS
	SCANCODE_KP_CLEAR           = sdl.SCANCODE_KP_CLEAR
	SCANCODE_KP_CLEARENTRY      = sdl.SCANCODE_KP_CLEARENTRY
	SCANCODE_KP_BINARY          = sdl.SCANCODE_KP_BINARY
	SCANCODE_KP_OCTAL           = sdl.SCANCODE_KP_OCTAL
	SCANCODE_KP_DECIMAL         = sdl.SCANCODE_KP_DECIMAL
	SCANCODE_KP_HEXADECIMAL     = sdl.SCANCODE_KP_HEXADECIMAL

	SCANCODE_LCTRL          = sdl.SCANCODE_LCTRL
	SCANCODE_LSHIFT         = sdl.SCANCODE_LSHIFT
	SCANCODE_LALT           = sdl.SCANCODE_LALT
	SCANCODE_LGUI           = sdl.SCANCODE_LGUI
	SCANCODE_RCTRL          = sdl.SCANCODE_RCTRL
	SCANCODE_RSHIFT         = sdl.SCANCODE_RSHIFT
	SCANCODE_RALT           = sdl.SCANCODE_RALT
	SCANCODE_RGUI           = sdl.SCANCODE_RGUI
	SCANCODE_MODE           = sdl.SCANCODE_MODE
	SCANCODE_AUDIONEXT      = sdl.SCANCODE_AUDIONEXT
	SCANCODE_AUDIOPREV      = sdl.SCANCODE_AUDIOPREV
	SCANCODE_AUDIOSTOP      = sdl.SCANCODE_AUDIOSTOP
	SCANCODE_AUDIOPLAY      = sdl.SCANCODE_AUDIOPLAY
	SCANCODE_AUDIOMUTE      = sdl.SCANCODE_AUDIOMUTE
	SCANCODE_MEDIASELECT    = sdl.SCANCODE_MEDIASELECT
	SCANCODE_WWW            = sdl.SCANCODE_WWW
	SCANCODE_MAIL           = sdl.SCANCODE_MAIL
	SCANCODE_CALCULATOR     = sdl.SCANCODE_CALCULATOR
	SCANCODE_COMPUTER       = sdl.SCANCODE_COMPUTER
	SCANCODE_AC_SEARCH      = sdl.SCANCODE_AC_SEARCH
	SCANCODE_AC_HOME        = sdl.SCANCODE_AC_HOME
	SCANCODE_AC_BACK        = sdl.SCANCODE_AC_BACK
	SCANCODE_AC_FORWARD     = sdl.SCANCODE_AC_FORWARD
	SCANCODE_AC_STOP        = sdl.SCANCODE_AC_STOP
	SCANCODE_AC_REFRESH     = sdl.SCANCODE_AC_REFRESH
	SCANCODE_AC_BOOKMARKS   = sdl.SCANCODE_AC_BOOKMARKS
	SCANCODE_BRIGHTNESSDOWN = sdl.SCANCODE_BRIGHTNESSDOWN
	SCANCODE_BRIGHTNESSUP   = sdl.SCANCODE_BRIGHTNESSUP
	SCANCODE_DISPLAYSWITCH  = sdl.SCANCODE_DISPLAYSWITCH
	SCANCODE_KBDILLUMTOGGLE = sdl.SCANCODE_KBDILLUMTOGGLE
	SCANCODE_KBDILLUMDOWN   = sdl.SCANCODE_KBDILLUMDOWN
	SCANCODE_KBDILLUMUP     = sdl.SCANCODE_KBDILLUMUP
	SCANCODE_EJECT          = sdl.SCANCODE_EJECT
	SCANCODE_SLEEP          = sdl.SCANCODE_SLEEP
	SCANCODE_APP1           = sdl.SCANCODE_APP1
	SCANCODE_APP2           = sdl.SCANCODE_APP2
)
