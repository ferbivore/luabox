-- see https://github.com/ferbivore/luabox/issues/5
function luabox.load()
    termbox.setinmode(termbox.inmode.mouse)
    termbox.set(1, 1, "L")
    termbox.flush()
end
function luabox.event(e)
    if e.key == termbox.key.CtrlC then
        termbox.close()
        luabox.quit()
    end
    termbox.clear()
    termbox.set(1, 1, "E")
    termbox.flush()
end