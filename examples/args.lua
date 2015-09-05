-- Prints command line arguments.

function printLine(x, y, str)
    str:gsub(".", function(char)
        termbox.set(x, y, char)
        x = x + 1
    end)
    termbox.flush()
end

function printLines(x, y, tab)
    for i, v in ipairs(tab) do
        printLine(x, y, v)
        y = y + 1
    end
end

function luabox.load()
    printLines(1, 1, luabox.args)
end

function luabox.event(e)
    luabox.quit()
end