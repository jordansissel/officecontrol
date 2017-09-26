package main

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
const (
	WH_KEYBOARD_LL = 13
	WH_KEYBOARD    = 2
	WM_KEYDOWN     = 256
	WM_SYSKEYDOWN  = 260
	WM_KEYUP       = 257
	WM_SYSKEYUP    = 261
	WM_KEYFIRST    = 256
	WM_KEYLAST     = 264
	PM_NOREMOVE    = 0x000
	PM_REMOVE      = 0x001
	PM_NOYIELD     = 0x002
	WM_LBUTTONDOWN = 513
	WM_RBUTTONDOWN = 516
	NULL           = 0
)

type (
	DWORD     uint32
	WPARAM    uintptr
	LPARAM    uintptr
	LRESULT   uintptr
	HANDLE    uintptr
	HINSTANCE HANDLE
	HMODULE   HANDLE
	HHOOK     HANDLE
	HWND      HANDLE
	HMENU     HANDLE
	LPVOID    uintptr
	LPCWSTR   uintptr
	ATOM      uint16

	HICON   HANDLE
	HCURSOR HANDLE
	HBRUSH  HANDLE
)

type WNDCLASS struct {
	style         uint
	lpfnWndProc   WNDPROC
	cbClsExtra    int
	cbWndExtra    int
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCWSTR
	lpszClassName LPCWSTR
}

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

type WNDPROC func(HWND, uint, WPARAM, LPARAM) LRESULT

type KBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo uintptr
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162805.aspx
type POINT struct {
	X, Y int32
}

type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}
