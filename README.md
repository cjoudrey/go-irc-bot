# go-irc-bot

`go-irc-bot` is a IRC bot written in go that is scriptable with Lua.

**This is still work in progress and should probably not be used in production. This was really just written as a learning exercise.**

## Usage

```
go-irc-bot [options...] <script.lua>
-host="irc.freenode.net": host to connect to
-ident="go-irc-bot": ident
-nickname="go-irc-bot": nickname
-port=6667: port to connect to
-realname="go-irc-bot": realname
```

## Scripting API

A [Lua5.1 VM](http://www.lua.org/manual/5.1/) is provided via [yuin/gopher-lua](https://github.com/yuin/gopher-lua) to script the bot.

A `bot` table is also exposed to control the bot within the VM.

#### bot.write(data)

Send a command to the IRC server.

i.e. `bot.write("JOIN #go-nuts")`

#### bot.on(event, callback(prefix, params))

Bind a handler to a given event.

An event can have many handlers.

`prefix` is the IRC message prefix.
`params` is a table of message arguments.
