package handlers

import (
	"fmt"
	"godiscordbot/utility"

	"github.com/bwmarrin/discordgo"
)

var BotId string

func SetBotId(id string) {
	BotId = id
}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	var a utility.Config
	err := utility.GetConfig(&a)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if m.Content == a.BotPrefix {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Noroc fra")
	}
}
