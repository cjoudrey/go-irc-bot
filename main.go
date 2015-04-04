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
	var secure = flag.Bool("secure", false, "connect over tls")
	var password = flag.String("password", "", "password")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...] <script.lua>\n", os.Args[0])
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Options:\n")
		fmt.Fprint(os.Stderr, "\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.Arg(0) == "" {
		flag.Usage()
		return
	}

	handler := *irc.NewEventHandler()

	client := irc.Client{
		Host:     *host,
		Port:     *port,
		Nickname: *nickname,
		Ident:    *ident,
		Realname: *realname,
		Secure:   *secure,
		Password: *password,
		Handler:  handler,
	}

	l := lua.NewState()
	defer l.Close()

	registerBotTable(l, &client, &handler)
	registerBotFunctions(l)

	if err := l.DoFile(flag.Arg(0)); err != nil {
		panic(err)
	}

	fmt.Printf("Connecting to %s:%d as %s\n", *host, *port, *nickname)

	client.Connect()

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

bot.notice = function(target, message)
	bot.write("NOTICE " .. target .. " :" .. message)
end
	`); err != nil {
		panic(err)
	}
}
