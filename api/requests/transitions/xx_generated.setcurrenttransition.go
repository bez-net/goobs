// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentTransitionParams represents the params body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionParams struct {
	requests.ParamsBasic

	// The name of the transition.
	TransitionName string `json:"transition-name"`
}

// Name just returns "SetCurrentTransition".
func (o *SetCurrentTransitionParams) Name() string {
	return "SetCurrentTransition"
}

/*
SetCurrentTransitionResponse represents the response body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionResponse struct {
	requests.ResponseBasic
}

// SetCurrentTransition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentTransition(
	params *SetCurrentTransitionParams,
) (*SetCurrentTransitionResponse, error) {
	data := &SetCurrentTransitionResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}