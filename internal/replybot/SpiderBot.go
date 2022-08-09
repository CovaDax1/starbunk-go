package replybot

import (
	"golang-discord-bot/internal/utils"
	"golang-discord-bot/internal/webhook"

	"github.com/bwmarrin/discordgo"
)

type SpiderBot struct {
}

func (b SpiderBot) ObserverName() string {
	return "Spider-Bot"
}

func (b SpiderBot) AvatarURL() string {
	return "https://i.dlpng.com/static/png/6569125_preview.png"
}

func (b SpiderBot) Pattern() string {
	return "\\bspider\\s?man\\b"
}

func (b SpiderBot) Response() string {
	return `Hey, it's "Spider-Man"! Don't forget the hyphen!`
}

func (b SpiderBot) HandleMessage(session *discordgo.Session, message discordgo.Message) {
	if utils.Match(b.Pattern(), message.Content) {
		webhook.WriteMessage(session, message.ChannelID, b.Response(), b.ObserverName(), b.AvatarURL())
	}
}
