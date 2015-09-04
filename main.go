package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "os"

/* This is the function in charge of running the actual Lua code. The events
 * channel is passed directly to Lua, and the quit channel is used to send a
 * message when it's time to quit the main loop. */
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

    /* termbox_module is declared in api.go */
    L.PreloadModule("termbox", termbox_module)

    /* Run main.lua and quit when it finishes. */
    err := L.DoFile("main.lua")
    if err != nil {
        panic(err)
    }
    quit <- true
}

func main() {
    events := make(chan lua.LValue)
    quit   := make(chan bool)
    go termbox_listener(events)
    go runner(events, quit)
    <- quit
}