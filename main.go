/*
 *	tg-autoziyao
 *	/main.go
 *	By Mozilla Public License Version 2.0.
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package main

import (
	"os"
	"strconv"

	"github.com/ziyao233/trobot"
	troLogger "github.com/ziyao233/trobot/logger"
	troCommand "github.com/ziyao233/trobot/command"
       )

func readBotToken() string {
	token, err := os.ReadFile("token")
	if err != nil {
		panic("Cannot read from ./token")
	}
	// XXX: Remove \n: so this only works on systems using
	// UNIX/Mac EOL
	return string(token[:len(token) - 1])
}

func cmdResend(cmd troCommand.Command) error {
	replied := cmd.Message.RepliedMessage
	if replied == nil || replied.ID == 0 {
		cmd.Println("You must reply to a message")
	}

	repeat := 1
	if len(cmd.Args) >= 2 {
		var err error
		repeat, err = strconv.Atoi(cmd.Args[1])
		if err != nil {
			cmd.Println("The argument must be an integer")
			return nil
		}
		if repeat > 5 {
			cmd.Println("Cannot repeat a message for more than 5 times")
			return nil
		}
	}

	for i := 0; i < repeat; i++ {
		cmd.Println(replied.Text)
	}
	return nil
}

var asciiChuang string =`
                   \ |
                   -+-
+--------+---\       O
|        ||   \
| RISC-V ||    \    O    O
|        ||     |  -+-  -+-
+--/-\---++-/-\-+ _/ \ _/ \
   \_/      \-/
`

func cmdChuang(cmd troCommand.Command) error {
	replied := cmd.Message.RepliedMessage
	var replyTo int64
	if replied == nil || replied.ID == 0 {
		replyTo = cmd.Message.ID
	} else {
		replyTo = replied.ID
	}

	cmd.Markup = "HTML"

	cmd.ReplyTof(replyTo, "<pre>%s</pre>", asciiChuang)

	return nil
}

func main() {
	trobot.SetAPIToken(readBotToken())
	trobot.SetLogLevel(troLogger.LDebug)

	troCommand.Register("hello",
		func(cmd troCommand.Command) error {
			cmd.Reply("Hi")
			return nil
		})

	troCommand.Register("resend", cmdResend)
	troCommand.Register("chuang", cmdChuang)

	trobot.Run()
}
