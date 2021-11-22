package whtranslate

// Translator performs translation between PluralKit dispatch events and Discord message embeds
type Translator struct{}

func NewTranslator() *Translator {
	return &Translator{}
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
