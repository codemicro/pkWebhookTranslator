package whtranslate

// Translator performs translation between PluralKit dispatch events and Discord message embeds
type Translator struct {
	options Option
}

func NewTranslator(options Option) *Translator {
	return &Translator{
		options: options,
	}
}

func (t *Translator) ignorePrivacyChanges() bool {
	return t.options & OptionIgnorePrivacyChanges != 0
}

// DispatchEvent represents a dispatch event sent by PluralKit
type DispatchEvent struct {
	Type         EventType `json:"type"`
	SigningToken string    `json:"signing_token"` // not used
	SystemID     string    `json:"system_id"`
	ID           string    `json:"id,omitempty"`
	GuildID      string    `json:"guild_id,omitempty"`
	Data         EventData `json:"data,omitempty"`
}

// EventData holds the data sent by PK in a DispatchEvent
type EventData map[string]interface{}

func (e EventData) AsString(key string) (string, bool) {
	if x, ok := e[key]; ok {
		if x == nil {
			return "", true
		}
		if y, ok := x.(string); ok {
			return y, true
		}
	}
	return "", false
}

type Option uint8

const (
	OptionDefault Option = 0
	OptionIgnorePrivacyChanges Option = 1 << iota
)