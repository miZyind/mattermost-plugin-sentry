package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Plugin is running!")
	case "/webhook":
		p.handleWebhook(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *Plugin) handleWebhook(w http.ResponseWriter, r *http.Request) {
	var webhook *SentryWebhook

	if err := json.NewDecoder(r.Body).Decode(&webhook); err != nil {
		p.API.LogError("Failed to decode body", "error", err.Error())
		return
	}

	var (
		user     = p.botID
		channel  = r.URL.Query().Get("channel")
		project  = strings.ToUpper(webhook.Project)
		message  = webhook.Event.Title
		url      = webhook.URL
		location = webhook.Event.Location
		color    = "#EC5E44"
	)

	post := &model.Post{
		UserId:    user,
		ChannelId: channel,
		Props: model.StringInterface{
			"attachments": []*model.SlackAttachment{
				{
					Fallback: fmt.Sprintf("[%s] %s", project, message),
					Title:    fmt.Sprintf("[%s] [%s](%s)", project, message, url),
					Text:     fmt.Sprintf("Occurred from: `%s`", location),
					Color:    color,
				},
			},
		},
	}

	if _, err := p.API.CreatePost(post); err != nil {
		p.API.LogError("Failed to create post", "error", err.Error())
	}

	w.WriteHeader(http.StatusOK)
}
