// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionParams struct {
	requests.ParamsBasic

	// Name of the desired scene collection.
	ScName string `json:"sc-name"`
}

// Name just returns "SetCurrentSceneCollection".
func (o *SetCurrentSceneCollectionParams) Name() string {
	return "SetCurrentSceneCollection"
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionResponse struct {
	requests.ResponseBasic
}

// SetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneCollection(
	params *SetCurrentSceneCollectionParams,
) (*SetCurrentSceneCollectionResponse, error) {
	data := &SetCurrentSceneCollectionResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}