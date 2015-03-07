package logutil

import "time"

var Rate = time.Millisecond * 300

func NewRateLimit(limit time.Duration) *ratelimit {
	return &ratelimit{limit: limit}
}

type ratelimit struct {
	limit time.Duration
	last  time.Time
}

func (r *ratelimit) Limit() bool {
	now := time.Now()
	if now.Sub(r.last) > r.limit {
		r.last = now
		return false
	}
	return true
}
