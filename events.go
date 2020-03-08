package gruutbot

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
)

func setupEvents(g *GruutBot) {
	filter, err := std.NewMsgFilter(context.Background(), g.client)

	if err != nil {
		g.log.Panic("Error creating message filter", err)
	}

	filter.SetPrefix(g.prefix)

	g.client.On(disgord.EvtMessageCreate, filter.NotByBot, filter.HasPrefix, filter.StripPrefix, messageCreate)
}

func messageCreate(s disgord.Session, evt *disgord.MessageCreate) {
	m := evt.Message
	content := strings.ToLower(m.Content)
	content = strings.TrimSpace(content)

	command := strings.Split(content, " ")[0]

	message := NewMessage(m, s)

	err := pluginManager.commands[command](*message, pluginManager.Log)

	if err != nil {
		s.Logger().Error(err)
	}
}
