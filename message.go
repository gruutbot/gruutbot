package gruutbot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandMessage struct {
	message    *discordgo.Message
	Parameters []string
	session    *discordgo.Session
}

func NewMessage(message *discordgo.Message, session *discordgo.Session) *CommandMessage {
	content := strings.ToLower(message.Content)
	content = strings.TrimSpace(content)
	parameters := strings.Split(content, " ")[1:]

	return &CommandMessage{
		message:    message,
		Parameters: parameters,
		session:    session,
	}
}

func (m *CommandMessage) Reply(replyMessage string, mentionAuthor bool) (err error) {
	if mentionAuthor {
		replyMessage = fmt.Sprintf("%s %s", m.message.Author.Mention(), replyMessage)
	}

	_, err = m.session.ChannelMessageSend(m.message.ChannelID, replyMessage)

	return
}

func (m *CommandMessage) Info() *discordgo.Message {
	return m.message
}

func (m *CommandMessage) GuildInfo() (*discordgo.Guild, error) {
	return m.session.Guild(m.message.GuildID)
}

func (m *CommandMessage) SessionInfo() *discordgo.Session {
	return m.session
}
