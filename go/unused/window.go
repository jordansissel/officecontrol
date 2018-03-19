package main

import (
	"fmt"
	"syscall"
	//~ "time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32                  = windows.NewLazySystemDLL("user32.dll")
	procCreateWindowEx      = user32.NewProc("CreateWindowExW")
	procSetWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procGetMessage          = user32.NewProc("GetMessageW")
	procRegisterClass       = user32.NewProc("RegisterClassW")

	kernel32            = windows.NewLazySystemDLL("kernel32.dll")
	procGetModuleHandle = kernel32.NewProc("GetModuleHandleW")
	procGetLastError    = kernel32.NewProc("GetLastError")
	keyboardHook        HHOOK
)

func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(ret)
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return ret != 0
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin uint32, msgFilterMax uint32) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	return int(ret)
}

func CreateWindowEx(dwExStyle DWORD, className string, windowName string, dwStyle DWORD, x int, y int, width int, height int, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) HWND {
	wName := []byte(windowName)
	ret, _, _ := procCreateWindowEx.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(&[]byte(className)[0])),
		uintptr(unsafe.Pointer(&wName[0])),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam),
	)

	hwnd := (HWND)(ret)

	fmt.Printf("hwnd: %d\n", int(hwnd))
	return hwnd
}

func GetModuleHandle(name string) HMODULE {
	if name == "" {
		ret, _, _ := procGetModuleHandle.Call(NULL)
		return HMODULE(ret)
	} else {
		mName := []byte(name)
		ret, _, _ := procGetModuleHandle.Call(
			uintptr(unsafe.Pointer(&mName[0])),
		)
		return HMODULE(ret)
	}

}

func RegisterClass(wc WNDCLASS) ATOM {
	ret, _, _ := procRegisterClass.Call(
		uintptr(unsafe.Pointer(&wc)),
	)
	return ATOM(ret)
}

func GetLastError() DWORD {
	ret, _, _ := procGetLastError.Call()
	return DWORD(ret)
}

func Start() {
	// defer user32.Release()
	keyboardHook = SetWindowsHookEx(WH_KEYBOARD_LL,
		(HOOKPROC)(func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
			if nCode < 0 {
				return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
			}
			if nCode == 0 && wparam == WM_KEYDOWN {
				UnhookWindowsHookEx(keyboardHook)
				kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
				fmt.Printf("key pressed: %q %v\n", kbdstruct.VkCode, kbdstruct)
				if kbdstruct.VkCode == 'J' {
					return 1
				}
			}
			return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
		}), 0, 0)

	//var msg MSG
	//for GetMessage(&msg, 0, 0, 0) != 0 {
	//}

	//UnhookWindowsHookEx(keyboardHook)
	//keyboardHook = 0
}
