-- Prints command line arguments.

write = require("lib/write")

function luabox.load()
    write.lines(luabox.args)
end

function luabox.event(e)
    if e.key == termbox.key.CtrlC then
        luabox.quit()
    end
end