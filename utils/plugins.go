package utils

import (
	"encoding/binary"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func GetUserVoiceChannel(guild *discordgo.Guild, userID string) string {
	var authorChannel string

	voiceStates := guild.VoiceStates
	foundChannel := false

	for _, vs := range voiceStates {
		if vs.UserID == userID {
			authorChannel = vs.ChannelID
			foundChannel = true

			break
		}
	}

	if foundChannel {
		return authorChannel
	}

	return ""
}

// Function based on DiscordGo's airhorn example (https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/main.go)
func LoadAudioFile(fs http.FileSystem, path string) (buffer [][]byte, err error) {
	buffer = make([][]byte, 0)

	var opuslen int16

	file, err := fs.Open(path)
	if err != nil {
		return
	}

	defer func() {
		_ = file.Close()
	}()

	for {
		err = binary.Read(file, binary.LittleEndian, &opuslen)
		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err = file.Close()
			if err != nil {
				return
			}

			return
		}

		if err != nil {
			return
		}

		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			return
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}
