package main

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Plugin represents MattermostPlugin struct
type Plugin struct {
	plugin.MattermostPlugin
	botID string
}

func main() {
	plugin.ClientMain(&Plugin{})
}
