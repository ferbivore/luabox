-- Functions for printing to Termbox.
local write = {}

write.xdefault = 1
write.ydefault = 1

function write.line(str, x, y)
    x = x or write.xdefault
    y = y or write.ydefault
    str:gsub(".", function(char)
        if char == '\n' then
            y = y + 1
            write.ydefault = y
            return
        end
        termbox.set(x, y, char)
        x = x + 1
    end)
    termbox.flush()
    return y
end

function write.lines(tab, x, y)
    for i, v in ipairs(tab) do
        y = write.line(v, x, y) + 1
    end
end

return write