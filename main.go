/*
 *	tg-autoziyao
 *	/main.go
 *	By Mozilla Public License Version 2.0.
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package main

import (
	"os"

	"github.com/ziyao233/trobot"
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

func main() {
	trobot.SetAPIToken(readBotToken())
	trobot.Run()
}
