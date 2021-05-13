// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionListParams represents the params body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListParams struct {
	requests.ParamsBasic
}

// Name just returns "GetTransitionList".
func (o *GetTransitionListParams) Name() string {
	return "GetTransitionList"
}

/*
GetTransitionListResponse represents the response body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListResponse struct {
	requests.ResponseBasic

	// Name of the currently active transition.
	CurrentTransition string `json:"current-transition"`

	Transitions []struct {
		// Name of the transition.
		Name string `json:"name"`
	} `json:"transitions"`
}

// GetTransitionList sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionList(
	paramss ...*GetTransitionListParams,
) (*GetTransitionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionListParams{{}}
	}
	params := paramss[0]
	data := &GetTransitionListResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}