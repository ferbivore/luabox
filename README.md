Luabox is a tiny library for writing console apps in Lua. It uses [gopher-lua](https://github.com/yuin/gopher-lua) as the interpreter and [termbox-go](https://github.com/nsf/termbox-go) to interface with the terminal.

    $ git clone https://github.com/ferbivore/luabox
    $ go build
    $ ./luabox examples/single-event.lua

You can also `go get -u https://github.com/ferbivore/luabox` if you like executables in your path. Here's what an extremely simple program, which prints `!` to your terminal and quits when you press Ctrl-C, looks like:

```lua
function luabox.load()
    termbox.set(1, 1, "!")
end
function luabox.event(e)
    if e.key == termbox.key.CtrlC then
        luabox.quit()
    end
end
```

Luabox calls `luabox.load()` when starting up and `luabox.event(e)` when receiving a Termbox event. `e` is just a table - you can run the `examples/single-event.lua` file to see what it looks like. The `termbox` module offers a bunch of functions and constants you can use - see `api.go` for details.
