-- Displays the terminal's size in real time.

write = require("lib/write")

function draw()
    termbox.clear()
    size = termbox.size()
    write.line("W: " .. tostring(size.width), 1, 1)
    write.line("H: " .. tostring(size.height), 1, 2)
end

function luabox.load()
    draw()
end

function luabox.event(e)
    if e.tbtype == termbox.event.resize then
        draw()
    else
        luabox.quit()
    end
end