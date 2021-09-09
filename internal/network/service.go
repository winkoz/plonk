package network

import (
	"fmt"
	"net/url"
	"strings"
)

// Service controls network validations
type Service interface {
	IsValidUrl(string) bool
}

type service struct {
	allowedHosts []string
}

func NewService() Service {
	return &service{
		allowedHosts: []string{"github.com"},
	}
}

// IsValidUrl checks the the passed in url is a valid URI and also
// matches the allow-list
func (s service) IsValidUrl(maybeUrl string) bool {
	_, err := url.ParseRequestURI(maybeUrl)
	if err != nil {
		return false
	}

	u, err := url.Parse(maybeUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	fmt.Printf("Host: %s", u.Host)

	return s.isAllowedHost(u.Host)
}

func (s service) isAllowedHost(host string) bool {
	for _, allowedHost := range s.allowedHosts {
		fmt.Printf("\t%s", allowedHost)
		if strings.EqualFold(host, allowedHost) {
			return true
		}
	}

	return false
}
