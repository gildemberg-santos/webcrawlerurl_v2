package load

type LoadPage struct {
	Url    string
	Source string
}

func (l *LoadPage) Load(url string) error {
	l.Url = url
	return nil
}
