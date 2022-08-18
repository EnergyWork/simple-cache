package test

import (
	"testing"

	cache "cache/src"
)

func TestCache(t *testing.T) {
	c := cache.New()
	c.Set("userId", 42)
	t.Logf("userId: %+v", c.Get("userId"))
	c.Delete("userId")
	t.Logf("userId: %+v", c.Get("userId"))
}

func TestCacheNil(t *testing.T) {
	c := cache.Cache{}
	c.Set("a", 1)
}
