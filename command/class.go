package command

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func GetClass(s *discordgo.Session, m *discordgo.MessageCreate, name string) error {
	class := Class{}
	fullURL := baseURL + "classes/" + name
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	jsonResponse := buf.String()
	json.Unmarshal([]byte(jsonResponse), &class)
	ClassMessage(s, m, class)

	return nil
}

func ClassMessage(s *discordgo.Session, m *discordgo.MessageCreate, class Class) {
	s.ChannelMessageSend(m.ChannelID, ("Class: " + class.Name))
	s.ChannelMessageSend(m.ChannelID, ("Hit Die: " + strconv.Itoa(class.HitDie)))

	for _, choices := range class.ProficiencyChoices {
		s.ChannelMessageSend(m.ChannelID, ("Choose " + strconv.Itoa(choices.Ammount) + " from the proficiencies below:"))
		var skills string
		for _, proficiency := range choices.From {
			skills += proficiency.Name + ", "
		}
		s.ChannelMessageSend(m.ChannelID, skills)
	}

	s.ChannelMessageSend(m.ChannelID, "You are proficient in the following:")
	var proficiencies string
	for _, proficiency := range class.Proficiences {
		proficiencies += proficiency.Name + ", "
	}
	s.ChannelMessageSend(m.ChannelID, proficiencies)

	s.ChannelMessageSend(m.ChannelID, "Saving Throws:")
	var savingThrows string
	for _, throw := range class.SavingThrows {
		savingThrows += throw.Name + ", "
	}
	s.ChannelMessageSend(m.ChannelID, savingThrows)

	s.ChannelMessageSend(m.ChannelID, "Subclasses:")
	for _, subclass := range class.Subclasses {
		s.ChannelMessageSend(m.ChannelID, subclass.Name)
	}
}
