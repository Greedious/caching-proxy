package cache

import "time"

type CacheObject struct {
	Content   []byte
	CreatedAt time.Time
}
