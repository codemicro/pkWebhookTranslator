package whtranslate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var ErrNoType = errors.New("whTranslate: event has no recognised type")

func (t *Translator) TranslateEvent(event *DispatchEvent) (*discordgo.MessageEmbed, error) {

	if event.Type == "" {
		return nil, ErrNoType
	}

	var (
		embed = new(discordgo.MessageEmbed)
		err   error
	)

	switch event.Type {
	case EventUpdateSystem:
		err = t.translateUpdateSystem(event, embed)
	default:
		return nil, ErrNoType
	}

	if err != nil {
		return nil, err
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("System ID: %s", event.SystemID),
	}

	return embed, nil
}

func (t *Translator) translateUpdateSystem(event *DispatchEvent, embed *discordgo.MessageEmbed) error {

	var sb strings.Builder

	if name, ok := event.Data.AsString("name"); ok {
		sb.WriteString(
			fmt.Sprintf("System name updated: new value `%s`\n", formatString(name)),
		)
	}

	embed.Title = "System updated"
	styleEmbed(embed, actionUpdate)

	return nil
}
