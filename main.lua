function tb_print(str, sx, y)
    local x = sx
    for char in str:gmatch(".") do
        tb_put_cell(x, y, char)
        x = x + 1
    end
end

tb_print("Hello world!", 0, 0)
tb_present()

function repr (obj, indent)
    indent = indent or 0
    if type(obj) == "table" then
        local result = "{\n"
        for key, val in pairs(obj) do
            result = result .. string.rep(" ", indent+2) ..
                tostring(key) .. " = " .. repr(val, indent+2) .. "\n"
        end
        return result .. string.rep(" ", indent) .. "}"
    else
        return tostring(obj)
    end
end

local luabox = {}
luabox.EXIT = false
luabox.exit = function(msg)
    tb_close()
    if msg then
        print(msg)
    end
    luabox.EXIT = true
end

luabox.onevent = function(recv)
    luabox.exit(repr(recv))
end

while not luabox.EXIT do channel.select(
{"|<-", luabox_events, function(ok, value)
    if ok then
        luabox.onevent(value)
    else
        luabox.exit("Event channel closed unexpectedly.")
    end
end }) end