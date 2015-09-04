local termbox = require("termbox")
local luabox  = require("lib/luabox")
local inspect = require("lib/inspect")

function luabox.onevent(recv)
    luabox.exit(inspect(recv))
end

luabox.loop()