package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationResp, error) {
	url := baseUrl + "location-area/" + location + "/"

	cached, exists := c.cache.Get(url)
	if exists {
		locationResp := LocationResp{}
		err := json.Unmarshal(cached, &locationResp)
		if err != nil {
			return LocationResp{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResp{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResp{}, err
	}

	c.cache.Add(url, data)

	locationResp := LocationResp{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationResp{}, err
	}

	return locationResp, nil
}
