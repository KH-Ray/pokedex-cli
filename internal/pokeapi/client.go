package pokeapi

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
    cache Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
        cache: NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}