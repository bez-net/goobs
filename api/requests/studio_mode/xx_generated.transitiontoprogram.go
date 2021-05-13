// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TransitionToProgramParams represents the params body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramParams struct {
	requests.ParamsBasic

	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`

		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}

// Name just returns "TransitionToProgram".
func (o *TransitionToProgramParams) Name() string {
	return "TransitionToProgram"
}

/*
TransitionToProgramResponse represents the response body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramResponse struct {
	requests.ResponseBasic
}

// TransitionToProgram sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TransitionToProgram(
	params *TransitionToProgramParams,
) (*TransitionToProgramResponse, error) {
	data := &TransitionToProgramResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}