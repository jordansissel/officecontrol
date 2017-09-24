package main

//go:generate rsrc -manifest main.manifest -o rsrc.syso

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/MakeNowJust/hotkey"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

// NoRepeat is a flag for RegisterHotKey.
// It is not defined by the hotkey library.
const NoRepeat = 0x4000

// App is a thing
type App struct {
	*walk.MainWindow
}

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

func keyboard() {
	input := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				WVk:     'O',
				DwFlags: 0,
			},
		},
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				WVk:     'O',
				DwFlags: win.KEYEVENTF_KEYUP,
			},
		},
	}

	count := win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))

	if int(count) != len(input) {
		fmt.Printf("SendInput failed. %d of %d items sent.\n", count, len(input))
	}
}

func hotkeys(app *App) {
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
		win.SetForegroundWindow(app.MainWindow.Handle())
		app.Activate()
	})
}

func main() {
	app := &App{}
	hotkeys(app)

	MainWindow{
		Title:    "SCREAMO",
		AssignTo: &app.MainWindow,
		MinSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "&Operations",
				Items: []MenuItem{
					Action{
						Text:     "Hello",
						Shortcut: Shortcut{Key: walk.KeyO},
						OnTriggered: func() {
							movemouse(10, 10)
							<-time.After(3 * time.Second)
							keyboard()
						},
					},
				},
			},
		},
	}.Run()
}
