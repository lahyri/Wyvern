package handler

import (
	"fmt"
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
	fmt.Println(m.Content)
	// Detects the "!" to see if the message is a command
	if m.Content[:3] == "dnd" {
		fullCommand := strings.Split(m.Content, " ")
		method := fullCommand[1]

		switch method {
		case "r":
			n, err1 := strconv.Atoi(fullCommand[2])
			d, err2 := strconv.Atoi(fullCommand[3])
			b := 0
			if len(fullCommand) >= 5 {
				b, _ = strconv.Atoi(fullCommand[4])
			}

			if err1 == nil && err2 == nil {
				command.Roll(n, d, b, s, m)
			}
		case "help":
			command.Help(s, m)
		default:
			command.Help(s, m)
		}

	}

}
