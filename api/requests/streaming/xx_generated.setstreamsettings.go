// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetStreamSettingsParams represents the params body for the "SetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetStreamSettings.
*/
type SetStreamSettingsParams struct {
	requests.ParamsBasic

	// Persist the settings to disk.
	Save bool `json:"save"`

	Settings struct {
		// The publish key.
		Key string `json:"key"`

		// The password for the streaming service.
		Password string `json:"password"`

		// The publish URL.
		Server string `json:"server"`

		// Indicates whether authentication should be used when connecting to the streaming server.
		UseAuth bool `json:"use-auth"`

		// The username for the streaming service.
		Username string `json:"username"`
	} `json:"settings"`

	// The type of streaming service configuration, usually `rtmp_custom` or `rtmp_common`.
	Type string `json:"type"`
}

// Name just returns "SetStreamSettings".
func (o *SetStreamSettingsParams) Name() string {
	return "SetStreamSettings"
}

/*
SetStreamSettingsResponse represents the response body for the "SetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetStreamSettings.
*/
type SetStreamSettingsResponse struct {
	requests.ResponseBasic
}

// SetStreamSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStreamSettings(
	params *SetStreamSettingsParams,
) (*SetStreamSettingsResponse, error) {
	data := &SetStreamSettingsResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}