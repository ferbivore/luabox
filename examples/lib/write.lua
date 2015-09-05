-- Functions for printing to Termbox.
local write = {}

write.xdefault = 1
write.ydefault = 1

function write.line(str, ox, oy)
    str = str or ""
    x = ox or write.xdefault
    y = oy or write.ydefault
    str:gsub(".", function(char)
        if char == "\n" then
            y = y + 1
            x = ox or write.xdefault
        else
            termbox.set(x, y, char)
            x = x + 1
        end
    end)
    termbox.flush()
    write.xdefault = ox or write.xdefault
    write.ydefault = y + 1
    return y + 1
end

function write.lines(tab, x, y)
    for i, v in ipairs(tab) do
        y = write.line(v, x, y)
    end
end

return write