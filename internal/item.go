package internal

import "time"

// represents cache Item
type Item struct {
	ExpirationTime *time.Time
	Value          interface{}
}

func (i *Item) Expired() bool {
	if i.ExpirationTime == nil {
		return false
	}
	return time.Now().After(*i.ExpirationTime)
}
