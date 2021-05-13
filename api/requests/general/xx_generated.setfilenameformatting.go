// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetFilenameFormattingParams represents the params body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingParams struct {
	requests.ParamsBasic

	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}

// Name just returns "SetFilenameFormatting".
func (o *SetFilenameFormattingParams) Name() string {
	return "SetFilenameFormatting"
}

/*
SetFilenameFormattingResponse represents the response body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingResponse struct {
	requests.ResponseBasic
}

// SetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetFilenameFormatting(
	params *SetFilenameFormattingParams,
) (*SetFilenameFormattingResponse, error) {
	data := &SetFilenameFormattingResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}