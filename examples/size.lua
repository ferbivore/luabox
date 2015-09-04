-- Displays the terminal's size in real time.

termbox = require("termbox")
function tbprint(x, y, str)
    str:gsub(".", function(char)
        termbox.set(x, y, char)
        x = x + 1
    end)
    termbox.flush()
end

function draw()
    termbox.clear()
    size = termbox.size()
    tbprint(1, 1, "W: " .. tostring(size.width))
    tbprint(1, 2, "H: " .. tostring(size.height))
    termbox.flush()
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