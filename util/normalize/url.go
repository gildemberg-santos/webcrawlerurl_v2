package normalize

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
)

type NormalizeUrl struct {
	Url     string
	BaseUrl string
}

func NewNormalizeUrl(url string) *NormalizeUrl {
	return &NormalizeUrl{Url: url}
}

func (l *NormalizeUrl) GetUrl() (url string, err error) {
	l.normalizeDomain()
	l.normalizeHttp()

	if !l.isUrl() {
		err = errors.New("url is invalid")
	}

	url = l.Url

	return
}

func (l *NormalizeUrl) normalizeHttp() {
	if !strings.HasPrefix(l.Url, "https://") && !strings.HasPrefix(l.Url, "http://") {
		l.Url = "https://" + l.Url
	} else if strings.HasPrefix(l.Url, "http://") {
		l.Url = strings.Replace(l.Url, "http://", "https://", 1)
	}
}

func (l *NormalizeUrl) normalizeDomain() {
	if l.BaseUrl == "" {
		return
	}

	linkCurrent, _ := url.Parse(l.Url)
	linkBase, _ := url.Parse(l.BaseUrl)
	linkCurrent.Path = strings.TrimSuffix(linkCurrent.Path, "/")
	if l.isExtension(linkCurrent.String(), linkCurrent.Path) {
		linkCurrent.Path = ""
	}

	if linkCurrent.Host == "" && linkCurrent.Path != "" {
		linkCurrent.Host = strings.TrimSuffix(linkBase.Host, "/")
		linkCurrent.Path = strings.TrimPrefix(linkCurrent.Path, "/")
		l.Url = fmt.Sprintf("https://%s/%s", linkCurrent.Host, linkCurrent.Path)
		return
	}

	if linkCurrent.Host != "" && linkCurrent.Path != "" {
		linkCurrent.Host = strings.TrimSuffix(linkCurrent.Host, "/")
		linkCurrent.Path = strings.TrimPrefix(linkCurrent.Path, "/")
		l.Url = fmt.Sprintf("https://%s/%s", linkCurrent.Host, linkCurrent.Path)
		return
	}

	if linkCurrent.Host == "" {
		linkCurrent.Host = strings.TrimSuffix(linkBase.Host, "/")
		l.Url = fmt.Sprintf("https://%s", linkCurrent.Host)
		return
	}
}

func (l *NormalizeUrl) isUrl() bool {

	linkCurrent, _ := url.Parse(l.Url)
	linkBase, _ := url.Parse(l.BaseUrl)
	if l.BaseUrl != "" && linkCurrent.Host != linkBase.Host {
		return false
	}

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

func (l *NormalizeUrl) isExtension(linkHost, linkPath string) bool {
	if linkPath != "" {
		linkPath = strings.TrimSuffix(linkPath, "/")

		validationExtension := func(path string, valid string) bool {
			return strings.HasSuffix(path, valid)
		}

		for _, extension := range []string{".pdf", ".jpg", ".gif", ".png"} {
			if validationExtension(linkPath, extension) {
				return true
			}
		}
	}

	return false
}
