package test

import (
	"testing"

	"github.com/energywork/simple-cache/internal"
	cache "github.com/energywork/simple-cache/pkg"
)

func TestCache(t *testing.T) {
	c := cache.New()
	c.Set("userId", 42, 0)

	v, err := c.Get("userId")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("userId: %+v", v)

	c.Delete("userId")

	v, err = c.Get("userId")
	if err == internal.ErrNotFound {
		t.Logf("%s -> success", err)
		return
	}
	t.Fatalf("must not find but: userId -> %+v", v)

}

func TestCacheNil(t *testing.T) {
	c := cache.Cache{}
	c.Set("a", 1, 0)
}
