package utils

import (
	"errors"
	"regexp"
)

type GitURI struct {
	Host  string
	Owner string
	Repo  string
}

func ParseURI(uri string) (*GitURI, error) {
	gitURI := GitURI{
		Host: "github.com",
	}
	r := regexp.MustCompile("^.*github.com[/:](?P<org>[^/]+)/(?P<name>[^/\\.]+)")
	if !r.MatchString(uri) {
		return &gitURI, errors.New("Could not parse uri")
	}

	matches := r.FindStringSubmatch(uri)
	gitURI.Owner = matches[1]
	gitURI.Repo = matches[2]

	return &gitURI, nil
}