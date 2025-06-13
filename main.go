package main

import (
	"log"
	"strings"
	"time"

	"github.com/skinnyfads/english-vocabs-bot/config"
	tele "gopkg.in/telebot.v4"
)

func main() {
	config.LoadEnv()

	pref := tele.Settings{
		Token:  config.Env.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		word := strings.Fields(c.Text())[0]
		meaning, err := GetMeaning(word)
		if err != nil {
			log.Printf("Error fetching meaning for %s: %v", word, err)
			return c.Reply("Sorry, I couldn't find the meaning of that word.")
		}
		if meaning == "" {
			return c.Reply("No definitions found for that word.")
		}
		return c.Reply(meaning)
	})

	b.Start()
}
