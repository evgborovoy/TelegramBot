package main

import (
	"flag"
	"log"
)

func main() {
	token := mustToken()

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token to access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
