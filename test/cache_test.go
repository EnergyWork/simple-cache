package test

import (
	"testing"
	"time"

	"github.com/energywork/simple-cache/internal"
	cache "github.com/energywork/simple-cache/pkg"
)

// go test -v -run TestXxx

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

func TestExpiredCache(t *testing.T) {
	c := cache.New()
	c.Set("key", 999, 3*time.Second)
	t.Log("added element")
	_, err := c.Get("key")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("after Get value still exist")
	time.Sleep(4 * time.Second)
	_, err = c.Get("key")
	if err != internal.ErrExpired || err != internal.ErrNotFound {
		t.Fatal(err)
	}
	t.Log("value exired")
}

func TestClear(t *testing.T) {
	c := cache.New()
	c.Set("a", 1, 0)
	c.Set("b", 2, 5*time.Second)
	c.Set("c", 3, 10*time.Second)
	t.Logf("CACHE: %+v", c)
	time.Sleep(6 * time.Second)
	c.Clear()
	t.Logf("CACHE: %+v", c)
	time.Sleep(6 * time.Second)
	c.Clear()
	t.Logf("CACHE: %+v", c)
	c.Free()
	t.Logf("CACHE: %+v", c)
}
