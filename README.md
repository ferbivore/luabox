Luabox is a tiny library for writing console apps in Lua. It uses [gopher-lua](https://github.com/yuin/gopher-lua) as the interpreter and [termbox-go](https://github.com/nsf/termbox-go) to interface with the terminal.

```shell
$ git clone https://github.com/ferbivore/luabox
$ go build
$ ./luabox examples/single-event.lua
```

You can also `go get https://github.com/ferbivore/luabox` if you like executables in your path. Here's what an extremely simple program, which prints `!` to your terminal and quits when you press a key, looks like:

```lua
termbox = require("termbox")
function luabox.load()
    termbox.set(1, 1, "!")
end
function luabox.event(e)
    luabox.quit()
end
```

Luabox calls `luabox.load()` when starting up and `luabox.event(e)` when receiving a Termbox event. `e` is just a table - you can run the `examples/single-event.lua` file to see what it looks like. The `termbox` module offers a bunch of functions and constants you can use - see `api.go` for details.