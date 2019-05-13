package command

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//Roll - Rolls n d-sided dices, sums them all with b bonus
func Roll(n, d, b int, s *discordgo.Session, m *discordgo.MessageCreate) {
	result := "Rolled: ["
	total := b
	for i := 0; i < n-1; i++ {
		val := rand.Intn(d) + 1
		result += strconv.Itoa(val) + ", "
		total += val
	}
	val := rand.Intn(d) + 1
	result += strconv.Itoa(val)
	total += val
	if b > 0 {
		result += "] +" + strconv.Itoa(b)
	} else if b == 0 {
		result += "]"
	} else {
		result += "] " + strconv.Itoa(b)
	}

	result += " = " + strconv.Itoa(total)
	s.ChannelMessageSend(m.ChannelID, result)
}
