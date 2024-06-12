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
	return &NormalizeUrl{Url: url, BaseUrl: url}
}

func (l *NormalizeUrl) GetUrl() (url string, err error) {
	l.normalizeHttp()
	l.normalizeDomain()

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
	if !strings.HasPrefix(l.BaseUrl, "https://") && !strings.HasPrefix(l.BaseUrl, "http://") {
		l.BaseUrl = "https://" + l.BaseUrl
	} else if strings.HasPrefix(l.BaseUrl, "http://") {
		l.BaseUrl = strings.Replace(l.BaseUrl, "http://", "https://", 1)
	}
}

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
	if l.isExtension(linkCurrent.String(), linkCurrent.Path) {
		linkCurrent.Path = ""
	}

	if linkCurrent.Host == "" && linkCurrent.Path != "" {
		linkCurrent.Host = strings.TrimSuffix(linkBase.Host, "/")
		linkCurrent.Path = strings.TrimPrefix(linkCurrent.Path, "/")
		l.Url = fmt.Sprintf("https:/%s/%s", linkCurrent.Host, linkCurrent.Path)
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

func (l *NormalizeUrl) isExtension(linkHost, linkPath string) bool {
	if linkPath != "" {
		linkPath = strings.TrimSuffix(linkPath, "/")

		validationExtension := func(path string, valid string) bool {
			return strings.HasSuffix(path, valid)
		}

		for _, extension := range []string{".png", ".jpg", ".jpeg", ".gif", ".svg", ".css", ".js", ".ico", ".woff", ".woff2", ".ttf", ".eot", ".otf", ".mp4", ".mp3", ".webm", ".ogg", ".wav", ".flac", ".aac", ".zip", ".tar", ".gz", ".rar", ".7z", ".exe", ".dmg", ".apk", ".csv", ".xls", ".xlsx", ".doc", ".docx", ".pdf", ".epub", ".iso", ".dmg", ".bin", ".ppt", ".pptx", ".odt", ".avi", ".mkv", ".xml", ".json", ".yml", ".yaml", ".rss", ".atom", ".swf", ".txt", ".dart", ".webp", ".bmp", ".tif", ".psd", ".ai", ".indd", ".eps", ".ps", ".zipx", ".srt", ".wasm", ".m4v", ".m4a", ".webp", ".weba", ".m4b", ".opus", ".ogv", ".ogm", ".oga", ".spx", ".ogx", ".flv", ".3gp", ".3g2", ".jxr", ".wdp", ".jng", ".hief", ".avif", ".apng", ".avifs", ".heif", ".heic", ".cur", ".ico", ".ani", ".jp2", ".jpm", ".jpx", ".mj2", ".wmv", ".wma", ".aac", ".tif", ".tiff", ".mpg", ".mpeg", ".mov", ".avi", ".wmv", ".flv", ".swf", ".mkv", ".m4v", ".m4p", ".m4b", ".m4r", ".m4a", ".mp3", ".wav", ".wma", ".ogg", ".oga", ".webm", ".3gp", ".3g2", ".flac", ".spx", ".amr", ".mid", ".midi", ".mka", ".dts", ".ac3", ".eac3", ".weba", ".m3u", ".m3u8", ".ts", ".wpl", ".pls", ".vob", ".ifo", ".bup", ".svcd", ".drc", ".dsm", ".dsv", ".dsa", ".dss", ".vivo", ".ivf", ".dvd", ".fli", ".flc", ".flic", ".flic", ".mng", ".asf", ".m2v", ".asx", ".ram", ".ra", ".rm", ".rpm", ".roq", ".smi", ".smil", ".wmf", ".wmz", ".wmd", ".wvx", ".wmx", ".movie", ".wri", ".ins", ".isp", ".acsm", ".djvu", ".fb2", ".xps", ".oxps", ".ps", ".eps", ".ai", ".prn", ".svg", ".dwg", ".dxf", ".ttf", ".fnt", ".fon", ".otf", ".cab"} {
			if validationExtension(linkPath, extension) {
				return true
			}
		}
	}

	return false
}
