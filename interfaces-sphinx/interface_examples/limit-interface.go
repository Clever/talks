package limitkeys

import (
	"fmt"
	"github.com/Clever/sphinx/common"
)

// START OMIT
type LimitKey interface {
	Key(common.Request) (string, error)
	Type() string
}

// END OMIT
