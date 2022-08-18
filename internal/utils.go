package internal

import "time"

//	getExpirationTime : returns the time after which the cache value must be cleared
//		if ttl == 0 then indefinite value
func GetExpirationTime(ttl time.Duration) *time.Time {
	if ttl == 0 {
		return nil
	}
	expirationTime := time.Now().Add(ttl)
	return &expirationTime
}
