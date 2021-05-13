// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetMuteParams represents the params body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteParams struct {
	requests.ParamsBasic

	// Desired mute status.
	Mute bool `json:"mute"`

	// Source name.
	Source string `json:"source"`
}

// Name just returns "SetMute".
func (o *SetMuteParams) Name() string {
	return "SetMute"
}

/*
SetMuteResponse represents the response body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteResponse struct {
	requests.ResponseBasic
}

// SetMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetMute(params *SetMuteParams) (*SetMuteResponse, error) {
	data := &SetMuteResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}