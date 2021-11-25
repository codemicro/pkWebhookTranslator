package whtranslate

import (
	"encoding/json"
	"errors"
	"fmt"
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
		// TODO: option to send ping webhooks?
		return nil, nil
	case EventUpdateSystem:
		err = t.translateUpdateSystem(event, embed)
	case EventCreateMember:
		err = t.translateCreateMember(event, embed)
	case EventUpdateMember:
		err = t.translateUpdateMember(event, embed)
	case EventDeleteMember:
		err = t.translateDeleteMember(event, embed)
	case EventCreateGroup:
		err = t.translateCreateGroup(event, embed)
	default:
		return nil, ErrNoType
	}

	if err != nil {
		return nil, err
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("System ID: %s", event.SystemID),
	}

	if event.ID != "" {
		embed.Footer.Text += fmt.Sprintf("\nEntity ID: %s", event.ID)
	}

	return embed.getMessageEmbed(), nil
}

func (t *Translator) translateUpdateSystem(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("System updated")
	embed.setStyle(actionUpdate)

	var data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Tag         string `json:"tag"`
		Timezone    string `json:"timezone"`
		Colour      string `json:"color" prefix:"#"`
		Banner      string `json:"banner"`
		Avatar      string `json:"avatar_url" readable:"Avatar"`
		Privacy     privacy `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Avatar != "" {
		embed.setImage(data.Avatar)
	} else if data.Banner != "" {
		embed.setImage(data.Banner)
	}

	return nil
}

func (t *Translator) translateCreateMember(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Member created")
	embed.setStyle(actionCreate)

	var data struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	embed.setContent(
		formatStatementMessage("Name", formatString(data.Name)),
	)

	return nil
}

func (t *Translator) translateUpdateMember(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Member updated")
	embed.setStyle(actionUpdate)

	var data struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Colour      string `json:"color" prefix:"#"`
		Birthday    string `json:"birthday"`
		Pronouns    string `json:"pronouns"`
		Avatar      string `json:"avatar_url" readable:"Avatar"`
		Banner      string `json:"banner"`
		Description string `json:"description"`
		KeepProxy   *bool   `json:"keep_proxy"`
		Privacy     privacy `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Avatar != "" {
		embed.setImage(data.Avatar)
	} else if data.Banner != "" {
		embed.setImage(data.Banner)
	}

	return nil
}

func (t *Translator) translateDeleteMember(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Member deleted")
	embed.setStyle(actionDelete)

	return nil
}

func (t *Translator) translateCreateGroup(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Group created")
	embed.setStyle(actionCreate)

	var data struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Description string `json:"description"`
		Icon      string `json:"icon"`
		Banner      string `json:"banner"`
		Colour      string `json:"color" prefix:"#"`
		Privacy     privacy `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Icon != "" {
		embed.setImage(data.Icon)
	} else if data.Banner != "" {
		embed.setImage(data.Banner)
	}

	return nil
}