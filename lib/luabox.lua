local termbox = require("termbox")
local luabox = {}

luabox.running = true

function luabox.exit(msg)
    termbox.close()
    if msg then print(msg) end
    luabox.running = false
end

function luabox.loop()
    while luabox.running do
        ok, event = luabox_events:receive()
        if ok then
            luabox.onevent(event)
        else
            luabox.exit("Event channel closed unexpectedly.")
        end
    end
end

function luabox.onevent(event)
    -- override this function, for now
end

return luabox