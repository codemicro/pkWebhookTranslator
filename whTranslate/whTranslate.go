package whtranslate

import "encoding/json"

// Translator performs translation between PluralKit dispatch events and Discord message embeds
type Translator struct {
	options Option
}

func NewTranslator(options ...Option) *Translator {
	var o Option
	// combine multiple options into one value
	for _, option := range options {
		o = o | option
	}

	return &Translator{
		options: o,
	}
}

func (t *Translator) ignorePrivacyChanges() bool {
	return t.options&OptionIgnorePrivacyChanges != 0
}

// DispatchEvent represents a dispatch event sent by PluralKit
type DispatchEvent struct {
	Type         EventType       `json:"type"`
	SigningToken string          `json:"signing_token"` // not used
	SystemID     string          `json:"system_id"`
	ID           string          `json:"id,omitempty"`
	GuildID      string          `json:"guild_id,omitempty"`
	Data         json.RawMessage `json:"data,omitempty"`
}

type Option uint8

const (
	OptionIgnorePrivacyChanges Option = 1 << iota
)
