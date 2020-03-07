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
	message := evt.Message
	content := strings.ToLower(message.Content)
	content = strings.TrimSpace(content)

	var err error
	if content == "ping" {
		_, err = evt.Message.Reply(context.Background(), s, "Pong!")
	} else if content == "pong" {
		_, err = evt.Message.Reply(context.Background(), s, "Ping!")
	}

	if err != nil {
		s.Logger().Error(err)
	}
}
