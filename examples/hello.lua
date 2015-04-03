local inspect = require("examples/inspect")

bot.on("JOIN", function(prefix, target)
  bot.privmsg(target, "Hey " .. prefix)
end)

bot.on("001", function(prefix, params)
  print("CONNECTED!")
  print(prefix)
  print(inspect(params))
  bot.join("#test-test-test")
  bot.join("#test-test-test-2")
end)

bot.on("433", function(prefix)
  print(prefix)
  bot.nick("Guest123141")
end)

orig_write = bot.write
bot.write = function(data)
  print("> " .. data)
  orig_write(data)
end
