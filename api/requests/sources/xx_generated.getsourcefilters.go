// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFiltersParams represents the params body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersParams struct {
	requests.ParamsBasic

	// Source name
	SourceName string `json:"sourceName"`
}

// Name just returns "GetSourceFilters".
func (o *GetSourceFiltersParams) Name() string {
	return "GetSourceFilters"
}

/*
GetSourceFiltersResponse represents the response body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersResponse struct {
	requests.ResponseBasic

	Filters []struct {
		// Filter name
		Name string `json:"name"`

		// Filter settings
		Settings map[string]interface{} `json:"settings"`

		// Filter type
		Type string `json:"type"`
	} `json:"filters"`
}

// GetSourceFilters sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilters(
	params *GetSourceFiltersParams,
) (*GetSourceFiltersResponse, error) {
	data := &GetSourceFiltersResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}