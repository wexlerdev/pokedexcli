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

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	
	newCache := pokecache.NewCache(cacheInterval)

	return Client { 
		httpClient: http.Client{
			Timeout:timeout,
		},
		cache: newCache,
	}
}

