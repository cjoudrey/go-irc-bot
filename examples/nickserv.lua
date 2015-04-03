local os = require("os")

bot.on("001", function(prefix, params)
  nickserv_password = os.getenv("NICKSERV_PASSWORD")

  if nickserv_password ~= "" then
    bot.privmsg("NickServ", "identify " .. os.getenv("NICKSERV_PASSWORD"))
  end
end)
