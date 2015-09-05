-- Prints every received event.

inspect = require("lib/inspect")
write   = require("lib/write")

function luabox.load()
    write.line("Enter any key to see the event.\nQuit with Ctrl-C.")
end

function luabox.event(e)
    if(e.key == termbox.key.CtrlC) then
        termbox.close()
        luabox.quit()
    end
    termbox.clear()
    write.line(inspect(e), 1, 1)
end