package main

import (
	f "fmt"
	"godiscordbot/handlers"
	"godiscordbot/utility"

	"github.com/bwmarrin/discordgo"
)

var (
	cfg   utility.Config
	BotId string
)

func init() {
	err := utility.GetConfig(&cfg)

	if err != nil {
		f.Println(err.Error())
		return
	}

}

func Start() {
	goBot, err := discordgo.New("Bot " + cfg.Token)

	if err != nil {
		f.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		f.Println(err.Error())
		return
	}

	BotId = u.ID
	handlers.SetBotId(BotId)

	goBot.AddHandler(MessageHandler)

	err = goBot.Open()

	if err != nil {
		f.Println(err.Error())
		return
	}

	f.Println("Everything is up..")
}

func main() {
	Start()
	<-make(chan struct{})
	return
}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	var a utility.Config
	err := utility.GetConfig(&a)

	if err != nil {
		f.Println(err.Error())
		return
	}

	if m.Content == a.BotPrefix {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Noroc, frate!")
	}
}
