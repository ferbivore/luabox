-- Prints every received event.

inspect = require("lib/inspect")
write   = require("lib/write")

local last_time_event = nil
local last_termbox_event = nil

function luabox.load()
    termbox.setinmode(termbox.inmode.altMouse)
    display()
end

function luabox.event(e)
    if(e.key == termbox.key.CtrlC) then
        termbox.close()
        luabox.quit()
    end

    if(e.type == "time") then
        last_time_event = e
    end

    if(e.type == "termbox") then
        last_termbox_event = e
    end

    display()
end

function find_key(tab, val)
    for k, v in pairs(tab) do
        if v == val then
            return k
        end
    end
    return "none"
end

function display()
    termbox.clear()
    write.xdefault = 1
    write.ydefault = 1
    if last_termbox_event then
        e = last_termbox_event
        write.line("Event type: " .. find_key(termbox.event, e.tbtype))
        if e.key == 0 then
            write.line("Character:  " .. e.char)
        else
            write.line("Key:        " .. find_key(termbox.key, e.key))
        end
        write.line("Modifier:   " .. find_key(termbox.modifier, e.modifier))
        write.line("Width:      " .. tostring(e.width))
        write.line("Height:     " .. tostring(e.height))
        write.line("Mouse X:    " .. tostring(e.mousex))
        write.line("Mouse Y:    " .. tostring(e.mousey))
    else
        write.line("Enter any key to see the event.\nQuit with Ctrl-C.")
    end
    if last_time_event then
        e = last_time_event
        write.line()
        write.line(tostring(e.tick) .. " seconds since startup.")
        write.ydefault = write.ydefault - 2
    end
end