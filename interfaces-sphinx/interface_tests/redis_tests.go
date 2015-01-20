package leakybucket

import (
	"testing"
)

// START OMIT
func TestCreate(t *testing.T) {
	s, err := New("tcp", os.Getenv("REDIS_URL"))
	assert.Nil(t, err)
	CreateTest(s)(t)
}

// END OMIT
