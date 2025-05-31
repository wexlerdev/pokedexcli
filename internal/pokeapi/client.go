package pokeapi

import (
	"net/http"
	"time"
	"github.com/wexlerdev/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	
	newCache := pokecache.NewCache(5 * time.Second)

	return Client {
		httpClient: http.Client{
			Timeout:timeout,
		},
		cache: newCache,
	}
}

