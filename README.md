Luabox is a tiny library for writing console apps in Lua. It uses [gopher-lua](https://github.com/yuin/gopher-lua) as the interpreter and [termbox-go](https://github.com/nsf/termbox-go) to interface with the terminal.

    $ git clone https://github.com/ferbivore/luabox
    $ go build
    $ ./luabox examples/single-event.lua