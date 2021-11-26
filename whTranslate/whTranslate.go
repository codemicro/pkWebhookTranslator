package whtranslate

import (
	"encoding/json"
	"time"
)

// Translator performs translation between PluralKit dispatch events and Discord message embeds
type Translator struct{}

func NewTranslator() *Translator {
	return &Translator{}
}

// DispatchEvent represents a dispatch event sent by PluralKit
type DispatchEvent struct {
	Type         EventType       `json:"type"`
	SigningToken string          `json:"signing_token"` // not used
	SystemID     string          `json:"system_id"`
	ID           string          `json:"id,omitempty"`
	Data         json.RawMessage `json:"data,omitempty"`
}

// privacy represents all known privacy fields used in models used in webhooks
type privacy struct {
	Name         string `json:"name_privacy"`
	Description  string `json:"description_privacy"`
	Avatar       string `json:"avatar_privacy"`
	Icon         string `json:"icon_privacy"`
	MemberList   string `json:"member_list_privacy"`
	GroupList    string `json:"group_list_privacy"`
	List         string `json:"list_privacy"`
	Front        string `json:"front_privacy"`
	FrontHistory string `json:"front_history_privacy"`
	Visibility   string `json:"visibility"`
	Birthday     string `json:"birthday_privacy"`
	Pronoun      string `json:"pronoun_privacy"`
	Metadata     string `json:"metadata_privacy"`
}

// switchModel represents a switch model from the PluralKit API.
type switchModel struct {
	ID        nullableString `json:"id" readable:"ID"`
	Timestamp *time.Time     `json:"timestamp" readable:"time"`
	Members   []string       `json:"members" readable:"Member IDs"`
}
