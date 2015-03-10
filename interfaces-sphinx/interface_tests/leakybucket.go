package leakybucket

import (
	"time"
)

// START OMIT
type Storage interface {
	Create(name string, capacity uint, rate time.Duration) (Bucket, error)
}

// END OMIT

type Bucket interface {
	Capacity() uint
	Remaining() uint
	// Reset returns when the bucket will be drained.
	Reset() time.Time
	// Add to the bucket. Returns bucket state after adding.
	Add(uint) (BucketState, error)
}

type BucketState struct {
	Capacity  uint
	Remaining uint
	Reset     time.Time
}
