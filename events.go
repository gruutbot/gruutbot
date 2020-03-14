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
	if mc.Author.Bot {
		return
	}

	if !strings.HasPrefix(mc.Content, g.prefix) {
		return
	}

	m := mc.Message
	content := strings.ToLower(m.Content)
	content = strings.TrimPrefix(content, g.prefix)
	content = strings.TrimSpace(content)

	command := strings.Split(content, " ")[0]

	message := NewMessage(m, s)

	commandFunc := pluginManager.commands[command]

	if commandFunc == nil {
		return
	}

	g.log.Debugf("Received command \"%s\"", command)

	err := commandFunc(*message, pluginManager.Log)

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
