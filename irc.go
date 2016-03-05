package main

import (
	"github.com/thoj/go-ircevent"
	l "log"
)

const (
	ircServer     = "irc.freenode.net"
	ircPort       = "6667"
	ircSslEnabled = false
	ircDebug      = false
	ircNick       = "gnatsbot"
	ircName       = "gnatsbugbot"
	ircChannel    = "#gnatsbottest"
)

// Create a message to send to IRC
func createMsg() {
}

func main() {
	// Create a channel taking a string
	//ch := make(chan string)
	connectIrc()
	//ch <- "hello"
}

func sendMsg(conn *irc.Connection, msg string) {
	conn.Privmsg(ircChannel, msg)
}

func connectIrc() {
	ircobj := irc.IRC(ircNick, ircName)
	ircobj.Debug = ircDebug
	ircobj.UseTLS = ircSslEnabled
	err := ircobj.Connect(ircServer + ":" + ircPort)
	if err != nil {
		l.Println(err.Error())
		l.Fatal("Can't connect to freenode.")
	}
	ircobj.AddCallback("001", func(e *irc.Event) {
		ircobj.Join(ircChannel)
	})

	//ircobj.AddCallback("PRIVMSG", func(e *irc.Event) {
	//	ircobj.Privmsg(ircChannel, e.Message())
	//})
	sendMsg(ircobj, "hello")
	//ircobj.Loop()
	ircobj.Disconnect()
}
