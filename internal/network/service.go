package network

import (
	"fmt"
	"net/url"
	"strings"
)

// Service controls network validations
type Service interface {
	IsValidUrl(string) bool
	IsUrl(string) bool
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
	if parsedUrl, err := s.parseUrl(maybeUrl); err == nil {
		return s.isAllowedHost(parsedUrl.Host)
	}

	return false
}

func (s service) IsUrl(maybeUrl string) bool {
	_, err := s.parseUrl(maybeUrl)
	return err == nil
}

func (s service) parseUrl(maybeUrl string) (*url.URL, error) {
	u, err := url.ParseRequestURI(maybeUrl)
	if err != nil {
		return nil, err
	} else if u == nil || u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("invalid url %s", maybeUrl)
	}

	return u, nil
}

func (s service) isAllowedHost(host string) bool {
	for _, allowedHost := range s.allowedHosts {
		if strings.EqualFold(host, allowedHost) {
			return true
		}
	}

	return false
}
