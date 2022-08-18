## Elementary in-memory cache

### Example
---
```go
func main() {
    c := cache.New()
    c.Set("userId", 42, 0)

    v, err := c.Get("userId")
    if err != nil {
        fmt.Println(err)
        return
    }
    t.Logf("userId: %+v", v)

    c.Delete("userId")

    v, err = c.Get("userId")
    if err == internal.ErrNotFound {
        fmt.Printf("%f -> success", err)
        return
    }
    fmt.Printf("must not find but: userId -> %+v", v)
}
```