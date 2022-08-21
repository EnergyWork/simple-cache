## Elementary in-memory cache with TTL for value

## Installation
go get -u github.com/EnergyWork/simple-cache

### Example
---
```go

import (
    cache "github.com/energywork/simple-cache/pkg"
)

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
    }
}
```
---