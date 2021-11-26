/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: BSD-2-Clause
 */

package whtranslate

import "github.com/bwmarrin/discordgo"

// discordEmbed is used to wrap a *discordgo.MessageEmbed with custom methods
type discordEmbed struct {
	*discordgo.MessageEmbed
}

func newDiscordEmbed() *discordEmbed {
	return &discordEmbed{
		MessageEmbed: new(discordgo.MessageEmbed),
	}
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

func (de *discordEmbed) setTitle(title string) {
	de.MessageEmbed.Title = title
}

func (de *discordEmbed) setContent(content string) {
	de.MessageEmbed.Description = content
}

func (de *discordEmbed) setImage(url string) {
	if url == "" {
		return
	}

	de.MessageEmbed.Image = &discordgo.MessageEmbedImage{
		URL: url,
	}
}
