package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationsListResp, error) {
	url := baseUrl + "location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

	cached, exists  := c.cache.Get(url)
	if exists {
		locationResp := LocationsListResp{}
		err := json.Unmarshal(cached, &locationResp)
		if err != nil {
			return LocationsListResp{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsListResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsListResp{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsListResp{}, err
	}

	c.cache.Add(url, data)

	locationResp := LocationsListResp{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationsListResp{}, err
	}

	return locationResp, nil
}
