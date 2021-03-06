package handler

import (
	"strconv"
	"strings"

	"github.com/lahyri/Wyvern/command"

	"github.com/bwmarrin/discordgo"
)

// CommandHandler - This function will be called (due to AddHandler in main.go) every time
// a new message is created on any channel that the autenticated bot has access to.
func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Detects the "!" to see if the message is a command
	if m.Content[:1] == "!" {
		fullCommand := strings.Split(m.Content, " ")
		method := fullCommand[0][1:]

		switch method {
		case "r":
			n, err1 := strconv.Atoi(fullCommand[1])
			d, err2 := strconv.Atoi(fullCommand[2])
			b := 0
			if len(fullCommand) >= 4 {
				b, _ = strconv.Atoi(fullCommand[3])
			}

			if err1 == nil && err2 == nil {
				command.Roll(n, d, b, s, m)
			}
		case "help":
			command.Help(s, m)
		case "class":
			command.GetClass(s, m, fullCommand[1])
		case "spell":
			command.GetSpell(s, m, fullCommand[1])
		default:
			command.Help(s, m)
		}

	}

}
