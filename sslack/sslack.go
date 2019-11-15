package sslack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SlackRequestBody struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func SendSlackNotification(webhookUrl string, schan string, msg string) error {
	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg, Channel: schan})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
