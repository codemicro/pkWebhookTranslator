/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: BSD-2-Clause
 */

package whtranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

var ErrNoType = errors.New("whTranslate: event has no recognised type")

// TranslateEvent translates a PluralKit event into a Discord message embed.
//
// PING events and UPDATE_GROUP_MEMBERS events are ignored, and return `nil, nil`. Unregistered events return
// `nil, ErrNoType`.
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
	case EventCreateMember:
		err = t.translateCreateMember(event, embed)
	case EventUpdateMember:
		err = t.translateUpdateMember(event, embed)
	case EventDeleteMember:
		err = t.translateDeleteMember(event, embed)
	case EventCreateGroup:
		err = t.translateCreateGroup(event, embed)
	case EventUpdateGroup:
		err = t.translateUpdateGroup(event, embed)
	case EventUpdateGroupMembers:
		// TODO: inop
		return nil, nil
	case EventDeleteGroup:
		err = t.translateDeleteGroup(event, embed)
	case EventLinkAccount:
		err = t.translateLinkAccount(event, embed)
	case EventUnlinkAccount:
		err = t.translateUnlinkAccount(event, embed)
	case EventUpdateSystemGuild:
		err = t.translateUpdateSystemGuild(event, embed)
	case EventUpdateMemberGuild:
		err = t.translateUpdateMemberGuild(event, embed)
	case EventCreateMessage:
		err = t.translateCreateMessage(event, embed)
	case EventCreateSwitch:
		err = t.translateCreateSwitch(event, embed)
	case EventUpdateSwitch:
		err = t.translateUpdateSwitch(event, embed)
	case EventUpdateSwitchMembers:
		err = t.translateUpdateSwitchMembers(event, embed)
	case EventDeleteSwitch:
		err = t.translateDeleteSwitch(event, embed)
	case EventDeleteAllSwitches:
		err = t.translateDeleteAllSwitches(event, embed)
	case EventSuccessfulImport:
		err = t.translateImportSuccess(event, embed)
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
		Name        nullableString `json:"name"`
		Description nullableString `json:"description"`
		Tag         nullableString `json:"tag"`
		Timezone    nullableString `json:"timezone"`
		Colour      nullableString `json:"color" prefix:"#"`
		Banner      nullableString `json:"banner"`
		Avatar      nullableString `json:"avatar_url" readable:"Avatar"`
		Privacy     privacy        `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, false)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Avatar.HasValue {
		embed.setImage(data.Avatar.Value)
	} else if data.Banner.HasValue {
		embed.setImage(data.Banner.Value)
	}

	return nil
}

func (t *Translator) translateCreateMember(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Member created")
	embed.setStyle(actionCreate)

	var data struct {
		Name nullableString `json:"name"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	embed.setContent(
		formatStatementMessage("Name", formatString(data.Name.Value)),
	)

	return nil
}

func (t *Translator) translateUpdateMember(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Member updated")
	embed.setStyle(actionUpdate)

	var data struct {
		Name        nullableString `json:"name"`
		DisplayName nullableString `json:"display_name"`
		Colour      nullableString `json:"color" prefix:"#"`
		Birthday    nullableString `json:"birthday"`
		Pronouns    nullableString `json:"pronouns"`
		Avatar      nullableString `json:"avatar_url" readable:"Avatar"`
		Banner      nullableString `json:"banner"`
		Description nullableString `json:"description"`
		KeepProxy   *bool          `json:"keep_proxy"`
		Privacy     privacy        `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, false)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Avatar.HasValue {
		embed.setImage(data.Avatar.Value)
	} else if data.Banner.HasValue {
		embed.setImage(data.Banner.Value)
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
		Name nullableString `json:"name"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	embed.setContent(
		formatStatementMessage("Name", formatString(data.Name.Value)),
	)

	return nil
}

func (t *Translator) translateUpdateGroup(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Group created")
	embed.setStyle(actionCreate)

	var data struct {
		Name        nullableString `json:"name"`
		DisplayName nullableString `json:"display_name"`
		Description nullableString `json:"description"`
		Icon        nullableString `json:"icon"`
		Banner      nullableString `json:"banner"`
		Colour      nullableString `json:"color" prefix:"#"`
		Privacy     privacy        `json:"privacy"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, false)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Icon.HasValue {
		embed.setImage(data.Icon.Value)
	} else if data.Banner.HasValue {
		embed.setImage(data.Banner.Value)
	}

	return nil
}

func (t *Translator) translateDeleteGroup(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Group deleted")
	embed.setStyle(actionDelete)

	return nil
}

func (t *Translator) translateLinkAccount(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Discord account linked")
	embed.setStyle(actionCreate)

	return nil
}

func (t *Translator) translateUnlinkAccount(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Discord account unlinked")
	embed.setStyle(actionDelete)

	return nil
}

func (t *Translator) translateUpdateSystemGuild(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Guild settings updated")
	embed.setStyle(actionUpdate)

	var data struct {
		ProxyingEnabled *bool          `json:"proxying_enabled"`
		AutoproxyMode   nullableString `json:"autoproxy_mode"`
		AutoproxyMember nullableString `json:"autoproxy_member"`
		Tag             nullableString `json:"tag"`
		TagEnabled      *bool          `json:"tag_enabled"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, false)
	if err != nil {
		return err
	}

	embed.setContent(c)

	return nil
}

func (t *Translator) translateUpdateMemberGuild(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Member guild settings updated")
	embed.setStyle(actionUpdate)

	var data struct {
		GuildID     nullableString `json:"guild_id" readable:"Guild ID"`
		DisplayName nullableString `json:"display_name"`
		Avatar      nullableString `json:"avatar_url" readable:"Avatar"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, false)
	if err != nil {
		return err
	}

	embed.setContent(c)

	if data.Avatar.HasValue {
		embed.setImage(data.Avatar.Value)
	}

	return nil
}

func (t *Translator) translateCreateMessage(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Message created")
	embed.setStyle(actionCreate)

	var data struct {
		Timestamp *time.Time     `json:"timestamp" readable:"Time"`
		ID        nullableString `json:"id" readable:"New message ID"`
		Original  nullableString `json:"original" readable:"Original message ID"`
		Sender    nullableString `json:"sender" readable:"Discord account ID"`
		Channel   nullableString `json:"channel" readable:"Channel ID"`
		Member    nullableString `json:"member" readable:"Member ID"`
	}

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, true)
	if err != nil {
		return err
	}

	embed.setContent(c)

	return nil
}

func (t *Translator) translateCreateSwitch(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Switch created")
	embed.setStyle(actionCreate)

	var data switchModel

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, true)
	if err != nil {
		return err
	}

	embed.setContent(c)

	return nil
}

func (t *Translator) translateUpdateSwitch(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Switch updated")
	embed.setStyle(actionUpdate)

	var data switchModel

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	c, err := structToString(data, false)
	if err != nil {
		return err
	}

	embed.setContent(c)

	return nil
}

func (t *Translator) translateUpdateSwitchMembers(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Switch members updated")
	embed.setStyle(actionUpdate)

	var data []string

	if err := json.Unmarshal(event.Data, &data); err != nil {
		return err
	}

	embed.setContent(
		formatStatementMessage("Member IDs", formatStringSlice(data)),
	)

	return nil
}

func (t *Translator) translateDeleteSwitch(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Switch deleted")
	embed.setStyle(actionDelete)

	return nil
}

func (t *Translator) translateDeleteAllSwitches(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("All switches deleted")
	embed.setStyle(actionDelete)

	return nil
}

func (t *Translator) translateImportSuccess(event *DispatchEvent, embed *discordEmbed) error {

	embed.setTitle("Import success")
	embed.setStyle(actionCreate)

	return nil
}
