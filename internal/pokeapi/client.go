package pokeapi

import (
	"net/http"
	"time"

	"github.com/jeremyu25/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		http.Client{
			Timeout: timeout,
		},
		pokecache.NewCache(interval),
	}
}
