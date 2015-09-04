package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "unicode/utf8"

/* This file implements a Lua-accessible API for termbox-go.
   Use L.PreloadModule("termbox", termbox_module) to import it. */

func termbox_module(L *lua.LState) int {
    mod := L.SetFuncs(L.NewTable(), termbox_exports)

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

    L.Push(mod)
    return 1
}

var termbox_exports = map[string]lua.LGFunction{
    "clear":  termbox_clear,
    "close":  termbox_close,
    "flush":  termbox_flush,
    "set":    termbox_set,
    "cursor": termbox_cursor,
    "size":   termbox_size,
    "sync":   termbox_sync,
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