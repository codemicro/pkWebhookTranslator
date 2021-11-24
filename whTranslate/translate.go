package whtranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
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

	var data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Tag         string `json:"tag"`
		Timezone    string `json:"timezone"`
		Colour      string `json:"color"`
		Banner      string `json:"banner"`
		Avatar      string `json:"avatar_url"`
		Privacy     struct {
			Description  string `json:"description_privacy"`
			MemberList   string `json:"member_list_privacy"`
			GroupList    string `json:"group_list_privacy"`
			Front        string `json:"front_privacy"`
			FrontHistory string `json:"front_history_privacy"`
		} `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	var sb strings.Builder

	if data.Name != "" {
		sb.WriteString(
			formatUpdateMessage("Name", formatString(data.Name)),
		)
	}

	if data.Description != "" {
		sb.WriteString(
			formatUpdateMessage("Description", formatString(data.Description)),
		)
	}

	if data.Tag != "" {
		sb.WriteString(
			formatUpdateMessage("Tag", formatString(data.Tag)),
		)
	}

	if data.Timezone != "" {
		sb.WriteString(
			formatUpdateMessage("Timezone", formatString(data.Timezone)),
		)
	}

	if data.Colour != "" {
		sb.WriteString(
			formatUpdateMessage("Colour", "#"+formatString(data.Colour)),
		)
	}

	if data.Banner != "" {
		sb.WriteString(
			formatUpdateMessage("Banner URL", formatString(data.Banner)),
		)
		embed.setImage(data.Banner)
	}

	if data.Avatar != "" {
		sb.WriteString(
			formatUpdateMessage("Avatar URL", formatString(data.Avatar)),
		)
		embed.setImage(data.Avatar)
	}

	if !t.ignorePrivacyChanges() {

		var (
			privacy = &data.Privacy
			x       [][2]string
		)

		if privacy.Description != "" {
			x = append(x, [2]string{"Description", privacy.Description})
		}

		if privacy.MemberList != "" {
			x = append(x, [2]string{"Member list", privacy.MemberList})
		}

		if privacy.GroupList != "" {
			x = append(x, [2]string{"Group list", privacy.GroupList})
		}

		if privacy.Front != "" {
			x = append(x, [2]string{"Front", privacy.Front})
		}

		if privacy.FrontHistory != "" {
			x = append(x, [2]string{"Front history", privacy.FrontHistory})
		}

		if len(x) != 0 {
			sb.WriteString("Privacy settings updated:\n")
			for _, y := range x {
				sb.WriteString(fmt.Sprintf(" • %s is now `%s`\n", y[0], y[1]))
			}
		}

	}

	embed.setContent(sb.String())

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
		Colour      string `json:"color"`
		Birthday    string `json:"birthday"`
		Pronouns    string `json:"pronouns"`
		Avatar      string `json:"avatar_url"`
		Banner      string `json:"banner"`
		Description string `json:"description"`
		KeepProxy   bool   `json:"keep_proxy"`
		Privacy     struct {
			Visibility  string `json:"visibility"`
			Name        string `json:"name_privacy"`
			Description string `json:"description_privacy"`
			Birthday    string `json:"birthday_privacy"`
			Pronoun     string `json:"pronoun_privacy"`
			Avatar      string `json:"avatar_privacy"`
			Metadata    string `json:"metadata_privacy"`
		} `json:"privacy"`
	}

	// TODO: This.

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	return nil
}
