package normalize

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

type NormalizeUrl struct {
	Url     string
	BaseUrl string
}

// NewNormalizeUrl creates a new instance of NormalizeUrl with the provided URL and sets the BaseUrl to the same value.
//
// Parameters:
// - url: The URL string to initialize the NormalizeUrl instance with.
//
// Returns:
// - *NormalizeUrl: A pointer to the newly created NormalizeUrl instance.
func NewNormalizeUrl(url string) *NormalizeUrl {
	return &NormalizeUrl{Url: url, BaseUrl: url}
}

// GetUrl retrieves the normalized URL after applying domain normalization and HTTP normalization.
//
// No parameters.
// Returns a string representing the normalized URL and an error if the URL is invalid.
func (l *NormalizeUrl) GetUrl() (url string, err error) {
	l.normalizeDomain()
	l.normalizeHttp()

	if !l.isUrl() {
		err = errors.New("url is invalid")
	}

	url = l.Url

	return
}

func (l *NormalizeUrl) MD5() string {
	url, _ := l.GetUrl()
	url = strings.Replace(url, "https://", "", 1)
	url = strings.Replace(url, "www.", "", 1)
	hash := md5.Sum([]byte(url))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

// normalizeHttp normalizes the URL and BaseUrl by ensuring they have the correct HTTP/HTTPS format.
//
// No parameters.
// Does not return any value.
func (l *NormalizeUrl) normalizeHttp() {
	// URL
	if !regexp.MustCompile(`^https?://`).MatchString(l.Url) {
		l.Url = "https://" + l.Url
	}
	if regexp.MustCompile(`^https:///`).MatchString(l.Url) {
		l.Url = strings.Replace(l.Url, "https:///", "https://", 1)
		l.Url = strings.Replace(l.Url, "http:///", "https://", 1)
	}

	// BASE URL
	if !regexp.MustCompile(`^https?://`).MatchString(l.BaseUrl) {
		l.BaseUrl = "https://" + l.BaseUrl
	}
	if regexp.MustCompile(`^https:/`).MatchString(l.BaseUrl) && !regexp.MustCompile(`^https://`).MatchString(l.BaseUrl) {
		l.BaseUrl = strings.Replace(l.BaseUrl, "https:/", "https://", 1)
		l.BaseUrl = strings.Replace(l.BaseUrl, "http:/", "https://", 1)
	}
}

// normalizeDomain normalizes the URL domain based on the BaseUrl.
//
// It parses the current URL and the BaseUrl, then adjusts the path and host accordingly.
// Returns nothing.
func (l *NormalizeUrl) normalizeDomain() {
	if l.BaseUrl == "" {
		return
	}

	linkCurrent, err := url.Parse(l.Url)
	if err != nil {
		return
	}
	linkBase, err := url.Parse(l.BaseUrl)
	if err != nil {
		return
	}

	linkCurrent.Path = strings.TrimSuffix(linkCurrent.Path, "/")
	if l.isExtension(linkCurrent.Path) {
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

	if linkCurrent.Host != "" {
		linkCurrent.Host = strings.TrimSuffix(linkBase.Host, "/")
		l.Url = fmt.Sprintf("https://%s", linkCurrent.Host)
		return
	}
}

// isUrl checks if the given URL is valid and matches the base URL, if provided.
//
// It parses the current URL and the BaseUrl, then checks if the host matches.
// If the BaseUrl is provided and the host does not match, it returns false.
// It then parses the URL as a request URI and checks if the host is an IP address.
// If the host is not an IP address, it checks if the host contains a dot (".") and returns the result.
// If the host is an IP address, it returns true.
//
// Returns a boolean indicating whether the URL is valid and matches the base URL.
func (l *NormalizeUrl) isUrl() bool {
	linkCurrent, err := url.Parse(l.Url)
	if err != nil {
		return false
	}
	linkBase, err := url.Parse(l.BaseUrl)
	if err != nil {
		return false
	}

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

// isExtension checks if the provided linkPath has a valid extension.
//
// Parameters:
// - linkPath: The path of the link.
// Returns a boolean indicating if the linkPath has a valid extension.
func (l *NormalizeUrl) isExtension(linkPath string) bool {
	if linkPath != "" {
		linkPath = strings.TrimSuffix(linkPath, "/")

		validationExtension := func(path string, valid string) bool {
			return strings.HasSuffix(path, valid)
		}

		for _, extension := range []string{".png", ".jpg", ".jpeg", ".gif", ".svg", ".css", ".js", ".ico", ".woff", ".woff2", ".ttf", ".eot", ".otf", ".mp4", ".mp3", ".webm", ".ogg", ".wav", ".flac", ".aac", ".zip", ".tar", ".gz", ".rar", ".7z", ".exe", ".dmg", ".apk", ".csv", ".xls", ".xlsx", ".doc", ".docx", ".pdf", ".epub", ".iso", ".dmg", ".bin", ".ppt", ".pptx", ".odt", ".avi", ".mkv", ".json", ".yml", ".yaml", ".rss", ".atom", ".swf", ".txt", ".dart", ".webp", ".bmp", ".tif", ".psd", ".ai", ".indd", ".eps", ".ps", ".zipx", ".srt", ".wasm", ".m4v", ".m4a", ".webp", ".weba", ".m4b", ".opus", ".ogv", ".ogm", ".oga", ".spx", ".ogx", ".flv", ".3gp", ".3g2", ".jxr", ".wdp", ".jng", ".hief", ".avif", ".apng", ".avifs", ".heif", ".heic", ".cur", ".ico", ".ani", ".jp2", ".jpm", ".jpx", ".mj2", ".wmv", ".wma", ".aac", ".tif", ".tiff", ".mpg", ".mpeg", ".mov", ".avi", ".wmv", ".flv", ".swf", ".mkv", ".m4v", ".m4p", ".m4b", ".m4r", ".m4a", ".mp3", ".wav", ".wma", ".ogg", ".oga", ".webm", ".3gp", ".3g2", ".flac", ".spx", ".amr", ".mid", ".midi", ".mka", ".dts", ".ac3", ".eac3", ".weba", ".m3u", ".m3u8", ".ts", ".wpl", ".pls", ".vob", ".ifo", ".bup", ".svcd", ".drc", ".dsm", ".dsv", ".dsa", ".dss", ".vivo", ".ivf", ".dvd", ".fli", ".flc", ".flic", ".flic", ".mng", ".asf", ".m2v", ".asx", ".ram", ".ra", ".rm", ".rpm", ".roq", ".smi", ".smil", ".wmf", ".wmz", ".wmd", ".wvx", ".wmx", ".movie", ".wri", ".ins", ".isp", ".acsm", ".djvu", ".fb2", ".xps", ".oxps", ".ps", ".eps", ".ai", ".prn", ".svg", ".dwg", ".dxf", ".ttf", ".fnt", ".fon", ".otf", ".cab"} {
			if validationExtension(linkPath, extension) {
				return true
			}
		}
	}

	return false
}
