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
		embed = newDiscordEmbed()
		err   error
	)

	switch event.Type {
	case EventPing:
		return nil, nil
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

	return embed.getMessageEmbed(), nil
}

func (t *Translator) translateUpdateSystem(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("System updated")
	embed.setStyle(actionUpdate)

	var sb strings.Builder

	if name, ok := event.Data.AsString("name"); ok {
		sb.WriteString(
			formatUpdateMessage("Name", formatString(name)),
		)
	}

	if desc, ok := event.Data.AsString("description"); ok {
		sb.WriteString(
			formatUpdateMessage("Description", formatString(desc)),
		)
	}

	if colour, ok := event.Data.AsString("color"); ok {
		sb.WriteString(
			formatUpdateMessage("Colour", "#"+formatString(colour)),
		)
		embed.setImage(formatColourURL(colour))
	}

	if banner, ok := event.Data.AsString("banner"); ok {
		sb.WriteString(
			formatUpdateMessage("Banner URL", formatString(banner)),
		)
		embed.setImage(banner)
	}

	if avatar, ok := event.Data.AsString("avatar_url"); ok {
		sb.WriteString(
			formatUpdateMessage("Avatar URL", formatString(avatar)),
		)
		embed.setImage(avatar)
	}

	embed.setContent(sb.String())

	return nil
}
