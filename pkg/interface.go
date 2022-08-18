package pkg

import "time"

type ICache interface {
	Set(string, interface{}, time.Duration)
	Get(string) (interface{}, error)
	Delete(string)
	Clear()
}
