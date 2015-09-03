function tb_print(str, sx, y)
	local x = sx
	for char in str:gmatch(".") do
		tb_put_cell(x, y, char)
		x = x + 1
	end
end

tb_print("Hello world!", 0, 0)
tb_present()

local exit = false
while not exit do channel.select(
{"|<-", luabox_events, function(ok, value)
	if not ok then
		tb_close()
		print("Event channel closed.")
		exit = true
	else
		tb_close()
		print("Event channel closed.")
		exit = true
	end
end }) end