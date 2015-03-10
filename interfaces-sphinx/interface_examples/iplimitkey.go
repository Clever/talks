// START OMIT
type ipLimitKey struct{}

func (ilk ipLimitKey) Type() string {
	return "ip"
}

func (ilk ipLimitKey) Key(request common.Request) (string, error) {

	if _, ok := request["remoteaddr"]; !ok {
		return "", EmptyKeyError{ilk, "No remoteaddr key in request"}
	}
	return "ip:" + request["remoteaddr"].(string), nil
}

// END OMIT
