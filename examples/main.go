//+build example

/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: BSD-2-Clause
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	whtranslate "github.com/codemicro/pkWebhookTranslator/whTranslate"
	"github.com/gofiber/fiber/v2"
	"os"
)

// Test webhook server implementation

var (
	pkToken = os.Getenv("PK_TOKEN")
	// discord webhook info
	whID          = os.Getenv("WH_ID")
	whToken       = os.Getenv("WH_TOKEN")
	serverAddress = os.Getenv("SERVER_ADDR")
)

func run() error {

	if pkToken == "" {
		return errors.New("missing env var PK_TOKEN")
	}

	if serverAddress == "" {
		serverAddress = "127.0.0.1:8080"
	}

	if whID == "" || whToken == "" {
		fmt.Println("Not sending Discord webhooks (WH_ID and/or WH_TOKEN not set)")
	} else {
		fmt.Printf("Starting with Discord webhook ID %#v\n", whID)
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
			fmt.Println("Body unmarshal error:", err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.SendString(err.Error())
		}

		if event.SigningToken != pkToken {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}

		demb, err := trans.TranslateEvent(event)
		fmt.Printf("Translated content: %#v %v\n", demb, err)
		if err != nil {
			c.Status(500)
			return err
		}

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

	return app.Listen(serverAddress)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
