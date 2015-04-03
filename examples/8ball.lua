local string = require("string")
local math = require("math")
local os = require("os")

local magic8ball = {"It is certain", "It is decidedly so", "Without a doubt", "Yes - definitely", "You may rely on it", "As I see it, yes", "Most likely", "Outlook good", "Yes", "Signs point to yes", "Reply hazy, try again", "Ask again later", "Better not tell you now", "Cannot predict now", "Concentrate and ask again", "Don't count on it", "My reply is no", "My sources say no", "Outlook not so good", "Very doubtful"}

math.randomseed(os.time())

-- Whenever someone types "!8ball message" the bot will
-- reply to the person directly or in the channel that
-- it was typed in.
bot.add_command("!8ball", function(nick, target, message)
  if message ~= "" then
    if string.sub(target, 1, 1) == "#" then
      bot.privmsg(target, nick .. ": " .. magic8ball[math.random(#magic8ball)])
    else
      bot.privmsg(nick, magic8ball[math.random(#magic8ball)])
    end
  end
end)
