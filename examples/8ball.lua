local string = require("string")
local math = require("math")
local os = require("os")

local magic8ball = {"It is certain", "It is decidedly so", "Without a doubt", "Yes - definitely", "You may rely on it", "As I see it, yes", "Most likely", "Outlook good", "Yes", "Signs point to yes", "Reply hazy, try again", "Ask again later", "Better not tell you now", "Cannot predict now", "Concentrate and ask again", "Don't count on it", "My reply is no", "My sources say no", "Outlook not so good", "Very doubtful"}

math.randomseed(os.time())

bot.on("PRIVMSG", function(prefix, params)
  target = params[1]
  message = params[2]

  if string.sub(target, 1, 1) == "#" and string.sub(message, 1, 7) == "!8ball " then
    bot.privmsg(target, magic8ball[math.random(#magic8ball)])
  end
end)
