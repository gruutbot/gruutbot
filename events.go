package gruutbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func setupEvents(g *GruutBot) {
	g.client.AddHandler(g.messageCreate)
}

func (g *GruutBot) messageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	m := mc.Message
	content := strings.ToLower(m.Content)
	content = strings.TrimPrefix(content, g.prefix)
	content = strings.TrimSpace(content)

	command := strings.Split(content, " ")[0]

	message := NewMessage(m, s)

	err := pluginManager.commands[command](*message, pluginManager.Log)

	if err != nil {
		g.log.Error(err)
	}
}
