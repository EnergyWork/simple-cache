## Elementary in-memory cache

### Example
---
```go
func main() {

    c := cache.New() // initiate cache

    c.Set("userId", 42)

    fmt.Printf("userId: %+v", c.Get("userId")) // will print "userId: 42"

    c.Delete("userId")

    fmt.Printf("userId: %+v", c.Get("userId")) // will print "userId: <nil>"
    
}
```