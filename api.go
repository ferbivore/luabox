package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "unicode/utf8"

/* This file implements a Lua-accessible API for termbox-go.
   Use L.PreloadModule("termbox", termbox_module) to import it. */

func termbox_module(L *lua.LState) int {
    mod := L.SetFuncs(L.NewTable(), termbox_exports)

    /* Painstakingly exported Termbox definitions. */

    color := L.NewTable()
    L.SetField(color, "default", lua.LNumber(termbox.ColorDefault))
    L.SetField(color, "black",   lua.LNumber(termbox.ColorBlack))
    L.SetField(color, "red",     lua.LNumber(termbox.ColorRed))
    L.SetField(color, "green",   lua.LNumber(termbox.ColorGreen))
    L.SetField(color, "yellow",  lua.LNumber(termbox.ColorYellow))
    L.SetField(color, "blue",    lua.LNumber(termbox.ColorBlue))
    L.SetField(color, "magenta", lua.LNumber(termbox.ColorMagenta))
    L.SetField(color, "cyan",    lua.LNumber(termbox.ColorCyan))
    L.SetField(color, "white",   lua.LNumber(termbox.ColorWhite))
    L.SetField(mod, "color", color)

    event := L.NewTable()
    L.SetField(event, "key",       lua.LNumber(termbox.EventKey))
    L.SetField(event, "resize",    lua.LNumber(termbox.EventResize))
    L.SetField(event, "mouse",     lua.LNumber(termbox.EventMouse))
    L.SetField(event, "error",     lua.LNumber(termbox.EventError))
    L.SetField(event, "interrupt", lua.LNumber(termbox.EventInterrupt))
    L.SetField(mod, "event", event)

    modifier := L.NewTable()
    L.SetField(modifier, "alt", lua.LNumber(termbox.ModAlt))
    L.SetField(mod, "modifier", modifier)

    inmode := L.NewTable()
    L.SetField(inmode, "esc",     lua.LNumber(termbox.InputEsc))
    L.SetField(inmode, "alt",     lua.LNumber(termbox.InputAlt))
    L.SetField(inmode, "mouse",   lua.LNumber(termbox.InputMouse))
    L.SetField(inmode, "current", lua.LNumber(termbox.InputCurrent))
    L.SetField(mod, "inmode", inmode)

    outmode := L.NewTable()
    L.SetField(outmode, "current", lua.LNumber(termbox.OutputCurrent))
    L.SetField(outmode, "normal",  lua.LNumber(termbox.OutputNormal))
    L.SetField(outmode, "c256",    lua.LNumber(termbox.Output256))
    L.SetField(outmode, "c216",    lua.LNumber(termbox.Output216))
    L.SetField(outmode, "gray",    lua.LNumber(termbox.OutputGrayscale))
    L.SetField(mod, "outmode", outmode)

    key := L.NewTable()
    L.SetField(key, "F1",            lua.LNumber(termbox.KeyF1))
    L.SetField(key, "F2",            lua.LNumber(termbox.KeyF2))
    L.SetField(key, "F3",            lua.LNumber(termbox.KeyF3))
    L.SetField(key, "F4",            lua.LNumber(termbox.KeyF4))
    L.SetField(key, "F5",            lua.LNumber(termbox.KeyF5))
    L.SetField(key, "F6",            lua.LNumber(termbox.KeyF6))
    L.SetField(key, "F7",            lua.LNumber(termbox.KeyF7))
    L.SetField(key, "F8",            lua.LNumber(termbox.KeyF8))
    L.SetField(key, "F9",            lua.LNumber(termbox.KeyF9))
    L.SetField(key, "F10",           lua.LNumber(termbox.KeyF10))
    L.SetField(key, "F11",           lua.LNumber(termbox.KeyF11))
    L.SetField(key, "F12",           lua.LNumber(termbox.KeyF12))
    L.SetField(key, "Insert",        lua.LNumber(termbox.KeyInsert))
    L.SetField(key, "Delete",        lua.LNumber(termbox.KeyDelete))
    L.SetField(key, "Home",          lua.LNumber(termbox.KeyHome))
    L.SetField(key, "End",           lua.LNumber(termbox.KeyEnd))
    L.SetField(key, "Pgup",          lua.LNumber(termbox.KeyPgup))
    L.SetField(key, "Pgdn",          lua.LNumber(termbox.KeyPgdn))
    L.SetField(key, "ArrowUp",       lua.LNumber(termbox.KeyArrowUp))
    L.SetField(key, "ArrowDown",     lua.LNumber(termbox.KeyArrowDown))
    L.SetField(key, "ArrowLeft",     lua.LNumber(termbox.KeyArrowLeft))
    L.SetField(key, "ArrowRight",    lua.LNumber(termbox.KeyArrowRight))
    L.SetField(key, "MouseLeft",     lua.LNumber(termbox.MouseLeft))
    L.SetField(key, "MouseMiddle",   lua.LNumber(termbox.MouseMiddle))
    L.SetField(key, "MouseRight",    lua.LNumber(termbox.MouseRight))
    L.SetField(key, "CtrlTilde",     lua.LNumber(termbox.KeyCtrlTilde))
    L.SetField(key, "Ctrl2",         lua.LNumber(termbox.KeyCtrl2))
    L.SetField(key, "CtrlSpace",     lua.LNumber(termbox.KeyCtrlSpace))
    L.SetField(key, "CtrlA",         lua.LNumber(termbox.KeyCtrlA))
    L.SetField(key, "CtrlB",         lua.LNumber(termbox.KeyCtrlB))
    L.SetField(key, "CtrlC",         lua.LNumber(termbox.KeyCtrlC))
    L.SetField(key, "CtrlD",         lua.LNumber(termbox.KeyCtrlD))
    L.SetField(key, "CtrlE",         lua.LNumber(termbox.KeyCtrlE))
    L.SetField(key, "CtrlF",         lua.LNumber(termbox.KeyCtrlF))
    L.SetField(key, "CtrlG",         lua.LNumber(termbox.KeyCtrlG))
    L.SetField(key, "Backspace",     lua.LNumber(termbox.KeyBackspace))
    L.SetField(key, "CtrlH",         lua.LNumber(termbox.KeyCtrlH))
    L.SetField(key, "Tab",           lua.LNumber(termbox.KeyTab))
    L.SetField(key, "CtrlI",         lua.LNumber(termbox.KeyCtrlI))
    L.SetField(key, "CtrlJ",         lua.LNumber(termbox.KeyCtrlJ))
    L.SetField(key, "CtrlK",         lua.LNumber(termbox.KeyCtrlK))
    L.SetField(key, "CtrlL",         lua.LNumber(termbox.KeyCtrlL))
    L.SetField(key, "Enter",         lua.LNumber(termbox.KeyEnter))
    L.SetField(key, "CtrlM",         lua.LNumber(termbox.KeyCtrlM))
    L.SetField(key, "CtrlN",         lua.LNumber(termbox.KeyCtrlN))
    L.SetField(key, "CtrlO",         lua.LNumber(termbox.KeyCtrlO))
    L.SetField(key, "CtrlP",         lua.LNumber(termbox.KeyCtrlP))
    L.SetField(key, "CtrlQ",         lua.LNumber(termbox.KeyCtrlQ))
    L.SetField(key, "CtrlR",         lua.LNumber(termbox.KeyCtrlR))
    L.SetField(key, "CtrlS",         lua.LNumber(termbox.KeyCtrlS))
    L.SetField(key, "CtrlT",         lua.LNumber(termbox.KeyCtrlT))
    L.SetField(key, "CtrlU",         lua.LNumber(termbox.KeyCtrlU))
    L.SetField(key, "CtrlV",         lua.LNumber(termbox.KeyCtrlV))
    L.SetField(key, "CtrlW",         lua.LNumber(termbox.KeyCtrlW))
    L.SetField(key, "CtrlX",         lua.LNumber(termbox.KeyCtrlX))
    L.SetField(key, "CtrlY",         lua.LNumber(termbox.KeyCtrlY))
    L.SetField(key, "CtrlZ",         lua.LNumber(termbox.KeyCtrlZ))
    L.SetField(key, "Esc",           lua.LNumber(termbox.KeyEsc))
    L.SetField(key, "Ctrl3",         lua.LNumber(termbox.KeyCtrl3))
    L.SetField(key, "Ctrl4",         lua.LNumber(termbox.KeyCtrl4))
    L.SetField(key, "CtrlBackslash", lua.LNumber(termbox.KeyCtrlBackslash))
    L.SetField(key, "Ctrl5",         lua.LNumber(termbox.KeyCtrl5))
    L.SetField(key, "Ctrl6",         lua.LNumber(termbox.KeyCtrl6))
    L.SetField(key, "Ctrl7",         lua.LNumber(termbox.KeyCtrl7))
    L.SetField(key, "CtrlSlash",     lua.LNumber(termbox.KeyCtrlSlash))
    L.SetField(key, "Space",         lua.LNumber(termbox.KeySpace))
    L.SetField(key, "Backspace2",    lua.LNumber(termbox.KeyBackspace2))
    L.SetField(key, "Ctrl8",         lua.LNumber(termbox.KeyCtrl8))
    L.SetField(mod, "key", key)

    L.Push(mod)
    return 1
}

var termbox_exports = map[string]lua.LGFunction{
    "clear":      termbox_clear,
    "close":      termbox_close,
    "flush":      termbox_flush,
    "set":        termbox_set,
    "cursor":     termbox_cursor,
    "size":       termbox_size,
    "sync":       termbox_sync,
    "setinmode":  termbox_setinmode,
    "setoutmode": termbox_setoutmode,
}

/* clear(fg?, bg?) - clear the backbuffer */
func termbox_clear(L *lua.LState) int {
    fg := termbox.ColorDefault
    bg := termbox.ColorDefault
    if L.GetTop() == 2 {
        fg = termbox.Attribute(L.CheckInt(1))
        bg = termbox.Attribute(L.CheckInt(2))
    }
    termbox.Clear(fg, bg)
    return 0
}

/* close() - close termbox and restore the terminal state */
func termbox_close(L *lua.LState) int {
    termbox.Close()
    return 0
}

/* flush() - sync the backbuffer and terminal */
func termbox_flush(L *lua.LState) int {
    termbox.Flush()
    return 0
}

/* set(x, y, char, fg?, bg?) - set a character in the backbuffer */
func termbox_set(L *lua.LState) int {
    if L.GetTop() < 3 || L.GetTop() > 5 {
        return 0
    }
    x := L.CheckInt(1)
    y := L.CheckInt(2)
    char, _ := utf8.DecodeRuneInString(L.CheckString(3))
    fg := termbox.ColorDefault
    bg := termbox.ColorDefault
    if L.GetTop() == 4 {
        fg = termbox.Attribute(L.CheckInt(4))
    }
    if L.GetTop() == 5 {
        fg = termbox.Attribute(L.CheckInt(4))
        bg = termbox.Attribute(L.CheckInt(5))
    }
    termbox.SetCell(x, y, char, fg, bg)
    return 0
}

/* cursor(x, y) - move the cursor to a specific position */
func termbox_cursor(L *lua.LState) int {
    if L.GetTop() != 2 {
        return 0
    }
    x := L.CheckInt(1)
    y := L.CheckInt(2)
    termbox.SetCursor(x, y)
    return 0
}

/* size() - returns a {width: w, height: h} table */
func termbox_size(L *lua.LState) int {
    w, h := termbox.Size()
    t := L.NewTable()
    L.SetField(t, "width",  lua.LNumber(w))
    L.SetField(t, "height", lua.LNumber(h))
    L.Push(t)
    return 1
}

/* sync() - forces a sync between the backbuffer and terminal */
func termbox_sync(L *lua.LState) int {
    termbox.Sync()
    return 0
}

/* setinmode(mode) - set the termbox input mode */
func termbox_setinmode(L *lua.LState) int {
    if L.GetTop() != 1 {
        return 0
    }
    mode := termbox.InputMode(L.CheckInt(1))
    L.Push(lua.LNumber(termbox.SetInputMode(mode)))
    return 1
}

/* setoutmode(mode) - set the termbox output mode */
func termbox_setoutmode(L *lua.LState) int {
    if L.GetTop() != 1 {
        return 0
    }
    mode := termbox.OutputMode(L.CheckInt(1))
    L.Push(lua.LNumber(termbox.SetOutputMode(mode)))
    return 1
}

/* This is a listener that polls termbox for events and sends them, as tables,
 * to the given channel. Run it as a goroutine. */
func termbox_listener(events chan lua.LValue) {
    L := lua.NewState()
    defer L.Close()

    t := L.NewTable()
    for {
        e := termbox.PollEvent()
        L.SetField(t, "type",     lua.LNumber(e.Type))
        L.SetField(t, "modifier", lua.LNumber(e.Mod))
        L.SetField(t, "key",      lua.LNumber(e.Key))
        L.SetField(t, "char",     lua.LString(string(e.Ch)))
        L.SetField(t, "width",    lua.LNumber(e.Width))
        L.SetField(t, "height",   lua.LNumber(e.Height))
        L.SetField(t, "mousex",   lua.LNumber(e.MouseX))
        L.SetField(t, "mousey",   lua.LNumber(e.MouseY))
        events <- t
    }
}