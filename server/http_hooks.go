package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type webhookRequest struct {
	ChannelID string `json:"channelID"`
	Message   string `json:"message"`
}

func decodeWebhookRequestFromJSON(data io.Reader) *webhookRequest {
	var o *webhookRequest
	err := json.NewDecoder(data).Decode(&o)
	if err != nil {
		return nil
	}
	return o
}

func (r *webhookRequest) ToJSON() []byte {
	b, _ := json.Marshal(r)
	return b
}

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
	body := decodeWebhookRequestFromJSON(r.Body)

	if _, err := p.API.CreatePost(&model.Post{
		UserId:    p.botID,
		ChannelId: body.ChannelID,
		Message:   body.Message,
	}); err != nil {
		p.API.LogError("Failed to create post", "err", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
