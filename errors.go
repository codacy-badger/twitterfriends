package twitterfriends

import (
	"errors"
)

var (
	errNoStatus = errors.New("A status is required, none given.")
)

