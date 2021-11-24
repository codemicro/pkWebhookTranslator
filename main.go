package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	whtranslate "github.com/codemicro/pkWebhookTranslator/whTranslate"
	"github.com/gofiber/fiber/v2"
)

// This file is meant as a test webhook server

var (
	pkToken = os.Getenv("PK_TOKEN")
	// discord webhook info
	whID = os.Getenv("WH_ID")
	whToken = os.Getenv("WH_TOKEN")
)

func run() error {

	if pkToken == "" {
		return errors.New("missing env var PK_TOKEN")
	}

	if whID == "" || whToken == "" {
		fmt.Println("Not sending Discord webhooks (WH_ID and/or WH_TOKEN not set)")
	}

	fmt.Printf("Starting with token %#v\n", pkToken)

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Request: %s %s %#v\n", c.Method(), c.OriginalURL(), c.Request().String())
		err := c.Next()
		fmt.Printf("Response: %#v\n", c.Response().String())
		return err
	})

	trans := whtranslate.NewTranslator() // haha hehe trans <3
	dgSession, _ := discordgo.New()

	app.Post("/wh", func(c *fiber.Ctx) error {

		event := new(whtranslate.DispatchEvent)

		if err := json.Unmarshal(c.Body(), event); err != nil {
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if event.SigningToken != pkToken {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}

		demb, err := trans.TranslateEvent(event)
		fmt.Printf("Translated content: %#v %v\n", demb, err)

		if !(whID == "" || whToken == "") {
			fmt.Println("Sending WH")
			_, err := dgSession.WebhookExecute(whID, whToken, true, &discordgo.WebhookParams{Embeds: []*discordgo.MessageEmbed{demb}})
			if err != nil {
				fmt.Printf("Webhook send error: %v\n", err)
			}
		}

		c.Status(fiber.StatusNoContent)
		return nil
	})

	return app.Listen("127.0.0.1:8080")
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
