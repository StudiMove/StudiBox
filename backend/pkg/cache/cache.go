package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache struct représente le cache utilisé dans l'application
type Cache struct {
	cache *cache.Cache
}

// NewCache crée une nouvelle instance de Cache avec un délai d'expiration et de purge défini
func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	return &Cache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Set ajoute un élément au cache avec une durée d'expiration personnalisée
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.cache.Set(key, value, duration)
}

// Get récupère un élément du cache par sa clé
func (c *Cache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

// Delete supprime un élément du cache par sa clé
func (c *Cache) Delete(key string) {
	c.cache.Delete(key)
}

// Increment incrémente une valeur numérique en cache (doit être un entier ou un float)
func (c *Cache) Increment(key string, n int64) error {
	return c.cache.Increment(key, n)
}

// Decrement décrémente une valeur numérique en cache (doit être un entier ou un float)
func (c *Cache) Decrement(key string, n int64) error {
	return c.cache.Decrement(key, n)
}

// Flush vide le cache complètement
func (c *Cache) Flush() {
	c.cache.Flush()
}
