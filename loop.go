package main

import "github.com/yuin/gopher-lua"

/* Load the file specified by luafile and start an event loop.
 * Runs luabox.load() when starting and luabox.event(e) for each event. */
func mainloop(luafile string, events chan lua.LValue) {
    L := lua.NewState()
    defer L.Close()

    /* load the termbox interface module (declared in api.go) */
    L.PreloadModule("termbox", termbox_module)

    /* create the luabox global */
    luabox := L.NewTable()
    L.SetGlobal("luabox", luabox)
    L.SetField(luabox, "events", lua.LChannel(events))
    L.SetFuncs(luabox, map[string]lua.LGFunction{
        "quit": global_quit,
    })

    /* run the given file */
    err := L.DoFile(luafile)
    if err != nil {
        termbox_close()
        panic(err)
    }

    /* call luabox.load(), panicking if it fails */
    if err := L.CallByParam(lua.P{
        Fn: L.GetField(L.GetGlobal("luabox"), "load"),
        NRet: 0,
        Protect: true,
    }); err != nil {
        termbox_close()
        panic(err)
    }

    /* start the event loop */
    for {
        msg := <- events
        err := L.CallByParam(lua.P{
            Fn: L.GetField(L.GetGlobal("luabox"), "event"),
            NRet: 0,
            Protect: true,
            }, msg)
        if err != nil {
            termbox_close()
            panic(err)
        }
    }
}

/* Sends a message to the 'quit' global channel.
 * Intended to be used from Lua. */
func global_quit(L *lua.LState) int {
    quit <- true
    return 0
}