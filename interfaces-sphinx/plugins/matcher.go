// A Matcher determines if a common.Request matches its requirements.
type Matcher interface {
	Match(common.Request) bool
}

// A MatcherFactory creates a Matcher based on a config.
type MatcherFactory interface {
	Type() string
	Create(config interface{}) (Matcher, error)
}
