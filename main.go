package main

//go:generate rsrc -manifest main.manifest -o rsrc.syso

import (
	"fmt"
	"time"
	"unsafe"

	// Until https://github.com/lxn/walk/pull/228 is merged
	"github.com/MakeNowJust/hotkey"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

// NoRepeat is a flag for RegisterHotKey.
// It is not defined by the hotkey library.
const NoRepeat = 0x4000

func movemouse(x, y int) {
	input := []win.MOUSE_INPUT{
		{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{
				Dx:      int32(x),
				Dy:      int32(y),
				DwFlags: win.MOUSEEVENTF_MOVE,
			},
		},
	}
	//fmt.Printf("SendInput(%d, %#v, %d)\n", len(input), input, unsafe.Sizeof(input[0]))
	count := win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))

	if int(count) != len(input) {
		fmt.Printf("SendInput failed. %d of %d items sent.\n", count, len(input))
	}
	//fmt.Printf(" => %d (lastError: %d)\n", i, win.GetLastError())
}

func typeText(text string) {
	input := make([]win.KEYBD_INPUT, len(text)*2)

	for i, value := range text {
		keydown := win.KEYBD_INPUT{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				WScan:   uint16(value),
				DwFlags: win.KEYEVENTF_UNICODE,
			},
		}

		keyup := keydown // copy of keydown
		keyup.Ki.DwFlags |= win.KEYEVENTF_KEYUP

		input[i*2] = keydown
		input[i*2+1] = keyup
	}

	//fmt.Printf("sendinput(%d, %#v, %d)\n", uint32(len(input)), input, int32(unsafe.Sizeof(input[0])))
	count := win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))

	if int(count) != len(input) {
		fmt.Printf("SendInput failed. %d of %d items sent.\n", count, len(input))
	}
}

func typeKey(shortcut walk.Shortcut) {
	var keydown []win.KEYBD_INPUT
	if shortcut.Modifiers&walk.ModShift > 0 {
		keydown = append(keydown, win.KEYBD_INPUT{
			Type: win.INPUT_KEYBOARD,
			Ki:   win.KEYBDINPUT{WVk: uint16(walk.KeyShift)},
		})
	}
	if shortcut.Modifiers&walk.ModControl > 0 {
		keydown = append(keydown, win.KEYBD_INPUT{
			Type: win.INPUT_KEYBOARD,
			Ki:   win.KEYBDINPUT{WVk: uint16(walk.KeyControl)},
		})
	}
	if shortcut.Modifiers&walk.ModAlt > 0 {
		keydown = append(keydown, win.KEYBD_INPUT{
			Type: win.INPUT_KEYBOARD,
			Ki:   win.KEYBDINPUT{WVk: uint16(walk.KeyAlt)},
		})
	}

	keydown = append(keydown, win.KEYBD_INPUT{
		Type: win.INPUT_KEYBOARD,
		Ki:   win.KEYBDINPUT{WVk: uint16(shortcut.Key)},
	})

	keyup := make([]win.KEYBD_INPUT, len(keydown))

	copy(keyup, keydown)
	for _, k := range keyup {
		k.Ki.DwFlags = win.KEYEVENTF_KEYUP
	}

	input := append(keyup, keydown...)

	count := win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))

	if int(count) != len(input) {
		fmt.Printf("SendInput failed. %d of %d items sent.\n", count, len(input))
	}
}

func hotkeys() {
	hk := hotkey.New()

	// Quoting https://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
	// > VK_OEM_1 0xBA
	// > For the US standard keyboard, the ';:' key
	key := hotkey.OEM_1

	hk.Register(hotkey.Ctrl|NoRepeat, key, func() {
		// Workaround. On Windows, for the `win.SetActiveWindow` call to succeed, this app must be the
		// foreground window.
		// Quoting https://msdn.microsoft.com/en-us/library/windows/desktop/ms646311(v=vs.85).aspx
		// > The SetActiveWindow function activates a window, but not if the application is in the
		// > background. The window will be brought into the foreground (top of Z-Order) if its
		// > application is in the foreground when the system activates the window. "
		//win.SetForegroundWindow(app.MainWindow.Handle())
		//app.Activate()
		Start()
	})
}

func main() {
	hotkeys()
	<-time.After(1 * time.Minute)
}
