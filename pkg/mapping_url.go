package pkg

import "errors"

type MappingUrl struct {
	Url   string
	Limit int
}

type responseSuccessUrls struct {
	MappingUrl struct {
		Urls []string `json:"urls"`
	} `json:"mapping_url"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

type responseErroUrls struct {
	Erro       string  `json:"erro"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

func (m *MappingUrl) Call() (interface{}, error) {
	ts := Timestamp{}
	ts.Start()

	if m.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroUrls{
			Erro:       err.Error(),
			Url:        m.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: 500,
		}
		return responseErro, err
	}

	pagina := LoadPage{
		Url: m.Url,
	}

	err := pagina.Load()
	if err != nil {
		ts.End()
		responseErro := responseErroUrls{
			Erro:       err.Error(),
			Url:        m.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: pagina.StatusCode,
		}
		return responseErro, err
	}

	extract_url := ExtractUrl{
		Url:   m.Url,
		Limit: m.Limit,
	}
	extract_url.Init(pagina.Source)
	extract_url.Call()
	ts.End()

	responseSuccess := responseSuccessUrls{
		Url:        m.Url,
		Timestamp:  ts.GetTime(),
		StatusCode: pagina.StatusCode,
	}
	responseSuccess.MappingUrl.Urls = extract_url.OutUrls

	return responseSuccess, nil
}
