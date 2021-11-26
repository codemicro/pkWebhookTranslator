/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: BSD-2-Clause
 */

package whtranslate

type EventType string

const (
	EventPing                EventType = "PING"
	EventUpdateSystem        EventType = "UPDATE_SYSTEM"
	EventCreateMember        EventType = "CREATE_MEMBER"
	EventUpdateMember        EventType = "UPDATE_MEMBER"
	EventDeleteMember        EventType = "DELETE_MEMBER"
	EventCreateGroup         EventType = "CREATE_GROUP"
	EventUpdateGroup         EventType = "UPDATE_GROUP"
	EventUpdateGroupMembers  EventType = "UPDATE_GROUP_MEMBERS"
	EventDeleteGroup         EventType = "DELETE_GROUP"
	EventLinkAccount         EventType = "LINK_ACCOUNT"
	EventUnlinkAccount       EventType = "UNLINK_ACCOUNT"
	EventUpdateSystemGuild   EventType = "UPDATE_SYSTEM_GUILD"
	EventUpdateMemberGuild   EventType = "UPDATE_MEMBER_GUILD"
	EventCreateMessage       EventType = "CREATE_MESSAGE"
	EventCreateSwitch        EventType = "CREATE_SWITCH"
	EventUpdateSwitch        EventType = "UPDATE_SWITCH"
	EventUpdateSwitchMembers EventType = "UPDATE_SWITCH_MEMBERS"
	EventDeleteSwitch        EventType = "DELETE_SWITCH"
	EventDeleteAllSwitches   EventType = "DELETE_ALL_SWITCHES"
	EventSuccessfulImport    EventType = "SUCCESSFUL_IMPORT"
)

// eventAction is used for styling embeds
type eventAction uint8

const (
	actionUndefined eventAction = iota
	actionUpdate
	actionCreate
	actionDelete
)
