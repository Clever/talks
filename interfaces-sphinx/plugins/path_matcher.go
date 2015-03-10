
type pathMatcherConfig struct {
	MatchAny []string `yaml:"match_any"`
}
type pathMatcher struct {
	Paths []*regexp.Regexp
}

type pathMatcherFactory struct{}
func (pmf pathMatcherFactory) Type() string {
	return "paths"
}
func (pmf pathMatcherFactory) Create(config interface{}) (Matcher, error) {
    ...
