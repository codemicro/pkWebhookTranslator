package whtranslate

import "github.com/bwmarrin/discordgo"

// discordEmbed is used to wrap a *discordgo.MessageEmbed with custom methods
type discordEmbed struct {
	*discordgo.MessageEmbed
}

func newDiscordEmbed() *discordEmbed {
	return &discordEmbed{}
}

func (de *discordEmbed) getMessageEmbed() *discordgo.MessageEmbed {
	return de.MessageEmbed
}

func (de *discordEmbed) setStyle(action eventAction) {
	var colour int
	switch action {
	case actionCreate:
		colour = 0x99c140
	case actionUpdate:
		colour = 0xdb7b2b
	case actionDelete:
		colour = 0xcc3232
	}
	de.MessageEmbed.Color = colour
}
