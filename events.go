package gruutbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func setupEvents(g *GruutBot) {
	g.client.AddHandler(g.messageCreate)
	g.client.AddHandler(g.ready)
}

func (g *GruutBot) messageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	m := mc.Message
	content := strings.ToLower(m.Content)
	content = strings.TrimPrefix(content, g.prefix)
	content = strings.TrimSpace(content)

	command := strings.Split(content, " ")[0]

	message := NewMessage(m, s)

	g.log.Debugf("Received command \"%s\"", command)

	err := pluginManager.commands[command](*message, pluginManager.Log)

	if err != nil {
		g.log.Error(err)
	}
}

func (g *GruutBot) ready(s *discordgo.Session, mc *discordgo.Ready) {
	err := s.UpdateStatus(1, g.prefix)
	if err != nil {
		g.log.Error("Failed to set status.", err)
	}
}
