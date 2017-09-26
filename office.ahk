class OSD {
  __New() {
    global
    local CustomColor = eeaa99  ; can be any rgb color (it will be made transparent below).
    Gui, OSD: New, +lastfound +alwaysontop -caption +toolwindow  ; +toolwindow avoids a taskbar button and an alt-tab menu item.
    Gui, OSD: Color, % CustomColor
    Gui, OSD: Font, s32  ; Set a large font size (32-point).
    Gui, OSD: Add, Text, vMyText cLime, ABCDEFGHIJ ; XX & YY serve to auto-size the window.
    ; Make all pixels of this color transparent and make the text itself translucent (150):
    WinSet, TransColor, %CustomColor% 150
    y := ("x0 y" . %A_ScreenHeight% - 100)
    Gui, OSD: Show, % "x" A_SCreenWidth - 400 "y" A_ScreenHeight - 150 NoActivate  ; NoActivate avoids deactivating the currently active window.
    Gui, OSD: Hide

    this.timer := ObjBindMethod(this, "HideOSD")
  }

  Display(text, duration) {
    GuiControl, OSD: , MyText, %text%
    Gui, OSD: Show, NoActivate
    timer := this.timer
    SetTimer % timer , % duration
  }

  HideOSD() {
    Gui, OSD: Hide
    timer := this.timer
    SetTimer % timer , Off
  }
}

osd := new OSD()
return

#IfWinActive, ahk_exe idea64.exe
^;::
Input, key, L1
if (key = "b") {
  SendInput, ^{F9} ; Intellij: build project 
  osd.Display("Build", 1000)
} Else {
  ; Fallthrough. Just type the key the user sent.
  SendInput, %key%
}
return