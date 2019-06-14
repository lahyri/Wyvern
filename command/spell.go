package command

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func GetSpell(s *discordgo.Session, m *discordgo.MessageCreate, name string) error {
	spell := Spell{}
	fullURL := baseURL + "spells/" + name
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	jsonResponse := buf.String()
	json.Unmarshal([]byte(jsonResponse), &spell)
	SpellMessage(s, m, spell)

	return nil
}

func SpellMessage(s *discordgo.Session, m *discordgo.MessageCreate, spell Spell) {
	s.ChannelMessageSend(m.ChannelID, ("Spell: " + spell.Name))
	for _, desc := range spell.Description {
		s.ChannelMessageSend(m.ChannelID, desc)
	}
}
