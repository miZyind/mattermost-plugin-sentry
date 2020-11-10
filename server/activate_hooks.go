package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

// OnActivate https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.OnActivate
func (p *Plugin) OnActivate() error {
	botID, err := p.Helpers.EnsureBot(
		&model.Bot{
			Username:    "sentry",
			DisplayName: "Sentry",
			Description: "Sends Sentry notifications to Mattermost.",
		},
		plugin.ProfileImagePath("assets/icon.png"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to ensure demo bot.")
	}

	p.botID = botID

	return nil
}
