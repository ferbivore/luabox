package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "os"

// This has to be a global - we write to it from a function called from Lua.
// See loop.go, function qlobal_quit()
var quit chan bool

func main() {
    events := make(chan lua.LValue, 256)
    quit = make(chan bool)

    // start up termbox
    // see https://github.com/nsf/termbox-go/issues/80
    os.Setenv("TERM", "xterm")
    if err := termbox.Init(); err != nil {
        panic(err)
    }
    defer termbox_close()

    go termbox_listener(events)
    go mainloop("main.lua", events)
    <- quit
}