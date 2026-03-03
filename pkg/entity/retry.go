package entity

import (
	"time"
)

// Retry represents tag for retry settings.
type Retry struct {
	*CommonTag
	start      time.Duration // start interval for retry.
	end        time.Duration // end interval for retry.
	multiplier float32       // multiplier of step.
	attempts   uint64        // attempts count of retry attempts.
}

// NewRetryTag builds new Retry instance.
func NewRetryTag(
	start time.Duration,
	end time.Duration,
	multiplier float32,
	attempts uint64,
) *Retry {
	return &Retry{
		CommonTag:  NewCommonTag(TagTypeRetry, ProxyTypeRetrier, ValueTypeEmpty),
		start:      start,
		end:        end,
		multiplier: multiplier,
		attempts:   attempts,
	}
}

func (r *Retry) IsForProxy(in ProxyType) bool {
	return r.ptype == in
}


// Start return time.Duration as int64 in nanoseconds.
func (r *Retry) Start() int64 {
	return r.start.Nanoseconds()
}

// End return time.Duration as int64 in nanoseconds.
func (r *Retry) End() int64 {
	return r.end.Nanoseconds()
}

// Multiplier return value to multiply.
func (r *Retry) Multiplier() float32 {
	return r.multiplier
}

// Attempts return value of attempts.
func (r *Retry) Attempts() uint64 {
	return r.attempts
}
