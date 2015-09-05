-- Prints command line arguments.

write = require("lib/write")

function luabox.load()
    write.lines(luabox.args)
end

function luabox.event(e)
    luabox.quit()
end