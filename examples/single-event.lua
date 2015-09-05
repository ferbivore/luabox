-- Displays a bunch of Unicode text on the screen, then waits for a single
-- event, closes termbox and prints the event out.

inspect = require("lib/inspect")
write   = require("lib/write")

-- Print a string to termbox, on a single line, starting at x, y.

function luabox.load()
    write.line("𝒜wesome!")
end

function luabox.event(e)
    termbox.close()
    print(inspect(e))
    luabox.quit()
end