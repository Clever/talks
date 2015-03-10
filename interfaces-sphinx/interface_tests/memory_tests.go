package leakybucket

import (
	"testing"
)

// START OMIT
func TestCreate(t *testing.T) {
	CreateTest(New())(t)
}

// END OMIT
