-- Displays a bunch of Unicode text on the screen, then waits for a single
-- event, closes termbox and prints the event out.

local inspect = require("lib/inspect")

-- Print a string to termbox, on a single line, starting at x, y.
function tbprint(x, y, str)
    str:gsub(".", function(char)
        termbox.set(x, y, char)
        x = x + 1
    end)
    termbox.flush()
end

function luabox.load()
    tbprint(1, 1, "𝒜wesome!")
end

function luabox.event(e)
    termbox.close()
    print(inspect(e))
    luabox.quit()
end