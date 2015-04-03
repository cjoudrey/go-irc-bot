package main

import "github.com/cjoudrey/irc"
import "github.com/yuin/gopher-lua"
import "os"
import "fmt"
import "flag"

func main() {
	var host = flag.String("host", "irc.freenode.net", "host to connect to")
	var port = flag.Int("port", 6667, "port to connect to")
	var nickname = flag.String("nickname", "go-irc-bot", "nickname")
	var ident = flag.String("ident", "go-irc-bot", "ident")
	var realname = flag.String("realname", "go-irc-bot", "realname")

	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...] <script.lua>\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	handler := *irc.NewEventHandler()

	client := irc.Client{
		Host: *host,
		Port: *port,
		Nickname: *nickname,
		Ident:    *ident,
		Realname: *realname,
		Handler:  handler,
	}

	l := lua.NewState()
	defer l.Close()

	registerBotTable(l, &client, &handler)
	registerBotFunctions(l)

	if err := l.DoFile(flag.Arg(0)); err != nil {
		panic(err)
	}

	go client.Connect()

	select {}
}

func registerBotTable(l *lua.LState, client *irc.Client, handler *irc.EventHandler) {
	bot := l.NewTable()
	bot.RawSet(lua.LString("write"), l.NewFunction(func(l *lua.LState) int {
		data := l.ToString(1)
		client.Write(data)
		return 1
	}))
	bot.RawSet(lua.LString("on"), l.NewFunction(func(l *lua.LState) int {
		command := l.ToString(1)
		function := l.ToFunction(2)

		handler.On(command, func(client *irc.Client, message *irc.Message) {
			lparams := l.NewTable()
			for _, param := range message.Params {
				lparams.Append(lua.LString(param))
			}

			l.CallByParam(lua.P{
				Fn:      function,
				NRet:    0,
				Protect: true,
			}, lua.LString(message.Prefix), lparams)
		})

		return 1
	}))
	l.SetGlobal("bot", bot)
}

func registerBotFunctions(l *lua.LState) {
	if err := l.DoString(`
bot.join = function(channel)
	bot.write("JOIN " .. channel)
end

bot.nick = function(new_nick)
	bot.write("NICK " .. new_nick)
end

bot.privmsg = function(target, message)
	bot.write("PRIVMSG " .. target .. " :" .. message)
end

bot.on("PING", function(prefix, params)
	print("PING? PONG!")
	bot.write("PONG " .. params[1])
end)
	`); err != nil {
		panic(err)
	}
}