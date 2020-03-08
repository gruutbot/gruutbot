package gruutbot

import (
	"context"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

type CommandMessage struct {
	message    *disgord.Message
	Parameters []string
	session    disgord.Session
}

func NewMessage(message *disgord.Message, session disgord.Session) *CommandMessage {
	content := strings.ToLower(message.Content)
	content = strings.TrimSpace(content)
	parameters := strings.Split(content, " ")[1:]

	return &CommandMessage{
		message:    message,
		Parameters: parameters,
		session:    session,
	}
}

func (m *CommandMessage) Reply(reply string, mentionAuthor bool) (err error) {
	if mentionAuthor {
		reply = fmt.Sprintf("%s %s", m.message.Author.Mention(), reply)
	}

	_, err = m.message.Reply(context.Background(), m.session, reply)

	return
}

func (m *CommandMessage) Info() *disgord.Message {
	return m.message
}

func (m *CommandMessage) GuildInfo() (*disgord.Guild, error) {
	return m.session.GetGuild(context.Background(), m.message.GuildID, disgord.IgnoreCache)
}

func (m *CommandMessage) SessionInfo() disgord.Session {
	return m.session
}
