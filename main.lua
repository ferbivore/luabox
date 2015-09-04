local termbox = require("termbox")
local inspect = require("lib/inspect")

local luabox = {}
luabox.EXIT = false
luabox.exit = function(msg)
    termbox.close()
    if msg then
        print(msg)
    end
    luabox.EXIT = true
end

luabox.onevent = function(recv)
    luabox.exit(inspect(recv))
end

while not luabox.EXIT do
    ok, value = luabox_events:receive()
    if ok then
        luabox.onevent(value)
    else
        luabox.exit("Event channel closed unexpectedly.")
    end
end