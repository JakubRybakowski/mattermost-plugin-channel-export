// Copyright (c) 2020-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-plugin-channel-export/server/pluginapi"
)

func isLicensed(_ *model.License, api *pluginapi.Wrapper) bool {
	return true
}
