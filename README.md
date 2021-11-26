# pkWebhookTranslator

*Golang library to translate [PluralKit](https://www.pluralkit.me) webhook dispatch events to Discord message embeds*

---

## Installation

```
go get github.com/codemicro/pkWebhookTranslator/whTranslate
```

## Usage

```go
translator := whtranslate.NewTranslator()
dgSession, _ := discordgo.New()

// receive dispatch event from PluralKit, unmarshal into a *whtranslate.DispatchEvent and validate `signing_token`

discordEmbed, err := translator.Translate(dispatchEvent)
if err != nil {
	// ...
}

_, err := dgSession.WebhookExecute(whID, whToken, true, &discordgo.WebhookParams{Embeds: []*discordgo.MessageEmbed{discordEmbed}})
if err != nil {
    // ...
}
```

## Example webhook server implementation

See [main.go](main.go).

## License

`pkWebhookTranslate` is distributed under the MIT license. See [LICENSE](LICENSE) for more details.