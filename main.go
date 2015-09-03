package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "encoding/json"
import "unicode/utf8"
import "os"

func tb_close(L *lua.LState) int {
	termbox.Close()
	return 0
}

/* Returns the width of the backbuffer. */
func tb_width(L *lua.LState) int {
	w, _ := termbox.Size()
	L.Push(lua.LNumber(w))
	return 1
}

/* Returns the height of the backbuffer. */
func tb_height(L *lua.LState) int {
	_, h := termbox.Size()
	L.Push(lua.LNumber(h))
	return 1
}

/* Clears the backbuffer. */
func tb_clear(L *lua.LState) int {
	// TODO: implement support for custom colors
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
	return 0
}

/* Synchronizes the backbuffer with the terminal. */
func tb_present(L *lua.LState) int {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
	return 0
}

/* Puts a character to the specified cell.
 * tb_put_cell(int x, int y, rune c) */
func tb_put_cell(L *lua.LState) int {
	// TODO: implement support for custom colors (tb_change_cell)
	// TODO: figure out what happens if we fuck up in Lua-verse
	// TODO: should we use 0- or 1-based indexing?
	x := L.CheckInt(1)
	y := L.CheckInt(2)
	cs := L.CheckString(3) // TODO: check length
	cc, _ := utf8.DecodeRuneInString(cs)
	termbox.SetCell(x, y, cc, termbox.ColorDefault, termbox.ColorDefault)
	return 0
}

/* Set the cursor position. Use (-1, -1) to hide. */
func tb_set_cursor(L *lua.LState) int {
	x := L.CheckInt(1)
	y := L.CheckInt(2)
	termbox.SetCursor(x, y)
	return 0
}


/* We use two goroutines: one to run the actual Lua code and one to listen for
 * termbox events and send them to the other. */

func runner(events chan lua.LValue, quit chan bool) {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("luabox_events", lua.LChannel(events))

	/* see https://github.com/nsf/termbox-go/issues/80 */
	os.Setenv("TERM","xterm")
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	/* Function declarations. */
	L.SetGlobal("tb_close",      L.NewFunction(tb_close))
	L.SetGlobal("tb_width",      L.NewFunction(tb_width))
	L.SetGlobal("tb_height",     L.NewFunction(tb_height))
	L.SetGlobal("tb_clear",      L.NewFunction(tb_clear))
	L.SetGlobal("tb_present",    L.NewFunction(tb_present))
	L.SetGlobal("tb_put_cell",   L.NewFunction(tb_put_cell))
	L.SetGlobal("tb_set_cursor", L.NewFunction(tb_set_cursor))

	/* Run main.lua and quit when it finishes. */
	err := L.DoFile("main.lua")
	if err != nil {
		panic(err)
	}

	quit <- true
}

func listener(events chan lua.LValue) {
	// TODO: do we actually need to create a state here?
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("luabox_events", lua.LChannel(events))

	for {
		// TODO: figure out a way to quit this thing
		e := termbox.PollEvent()
		j, _ := json.Marshal(e)
		events <- lua.LString(j)
	}
}


func main() {
	L := lua.NewState()
	defer L.Close()

	// L.SetGlobal("tb_init", L.NewFunction(tb_init))
	// ...

	events := make(chan lua.LValue)
	quit   := make(chan bool)
	go listener(events)
	go runner(events, quit)
	<- quit
}