package repository

import (
	"io"
	"time"
)

type CacheManager interface {
	io.Closer
	Get(key string) (string, error)
	Set(key string, value interface{}, exp time.Duration) error
}
