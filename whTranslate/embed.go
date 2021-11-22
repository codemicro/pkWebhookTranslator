package whtranslate

import "github.com/bwmarrin/discordgo"

func styleEmbed(embed *discordgo.MessageEmbed, action eventAction) {
	var colour int
	switch action {
	case actionCreate:
		colour = 0x99c140
	case actionUpdate:
		colour = 0xdb7b2b
	case actionDelete:
		colour = 0xcc3232
	}
	embed.Color = colour
}
