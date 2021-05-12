package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/radovskyb/watcher"
)

var (
	ChatId = os.Getenv("CHAT_ID")
	ApiKey = os.Getenv("T_API")
)

func main() {
	fmt.Println("Server-Watcher Started!")
	CHAT_ID, err := strconv.ParseInt(ChatId, 10, 64)
	if err != nil {
		log.Fatalf("%v", err)
	}
	bot := initTelegram(CHAT_ID)
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write)
	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
				var out bytes.Buffer
				fls := exec.Command("bash", "-c", "tail -5 /var/log/auth.log | grep 'Accepted'")
				fls.Stdout = &out
				err := fls.Run()
				if err != nil {
					log.Printf("Error -:%v", err)
					break
				}
				msg := tgbotapi.NewMessage(CHAT_ID, out.String())
				bot.Send(msg)
			}
		}
	}()
	if err := w.Add("/var/log/auth.log"); err != nil {
		log.Fatalln(err)
	}
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}
	if err := w.Start(time.Second * 2); err != nil {
		log.Fatalln(err)
	}
}

func initTelegram(CHAT_ID int64) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(ApiKey)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	msg := tgbotapi.NewMessage(CHAT_ID, "Bot Started!")
	bot.Send(msg)
	return bot
}
