-- An example shortcut for adding !commands
bot.add_command = function(command, callback)
  bot.on("PRIVMSG", function(prefix, params)
    nick = string.sub(prefix, 1, string.find(prefix, "!") - 1)
    target = params[1]
    message = params[2]

    if string.sub(message, 1, string.len(command)) == command then
      callback(nick, target, string.sub(message, string.len(command) + 2))
    end
  end)
end

require("examples/nickserv")
require("examples/8ball")

-- 001 is the event that happens when bot connects to server
-- as per IRC's RFC.
bot.on("001", function(prefix, params)
  bot.join("#go-irc-bot")
end)
