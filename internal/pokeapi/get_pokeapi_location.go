package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetShallowLocation(pageURL *string) (RespShallowLocations, error) {
    url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if ok {
		var locationResp RespShallowLocations
		if err := json.Unmarshal(data, &locationResp); err != nil {
			return RespShallowLocations{}, err
		}
	
		return locationResp, nil
	}

    res, err := http.Get(url)
    if err != nil {
        return RespShallowLocations{}, fmt.Errorf("error creating request: %w", err)
    }
    defer res.Body.Close()

    data, err = io.ReadAll(res.Body)
    if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error reading response: %w", err)
	}

	var locationResp RespShallowLocations
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}

func (c *Client) GetDeepLocation(pageURL string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + pageURL

	data, ok := c.cache.Get(url)
	if ok {
		var locationResp RespDeepLocations
		if err := json.Unmarshal(data, &locationResp); err != nil {
			return RespDeepLocations{}, err
		}
	
		return locationResp, nil
	}

    res, err := http.Get(url)
    if err != nil {
        return RespDeepLocations{}, fmt.Errorf("error creating request: %w", err)
    }
    defer res.Body.Close()

    data, err = io.ReadAll(res.Body)
    if err != nil {
		return RespDeepLocations{}, fmt.Errorf("error reading response: %w", err)
	}

	var locationResp RespDeepLocations
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return RespDeepLocations{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
