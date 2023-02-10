package pkg

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

type NormalizeUrl struct {
	Url string
}

func (l *NormalizeUrl) GetUrl() (url string, err error) {
	l.normalizeUrl()

	if !l.isUrl() {
		err = errors.New("Url is invalid")
	}

	url = l.Url

	return
}

func (l *NormalizeUrl) normalizeUrl() {
	if !strings.HasPrefix(l.Url, "https://") && !strings.HasPrefix(l.Url, "http://") {
		l.Url = "https://" + l.Url
	} else if strings.HasPrefix(l.Url, "http://") {
		l.Url = strings.Replace(l.Url, "http://", "https://", 1)
	}
}

func (l *NormalizeUrl) isUrl() bool {
	url, err := url.ParseRequestURI(l.Url)
	if err != nil {
		return false
	}

	address := net.ParseIP(url.Host)
	if address == nil {
		return strings.Contains(url.Host, ".")
	}

	return true
}
