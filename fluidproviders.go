package fluidcurrency

import (
  "fmt"
	"encoding/json"
	"strings"
	"time"
)

// Provider are abstract
type Provider interface {
	GetRatesFromAPI() Rates
}
// APIFixerProvider are abstract
type APIFixerProvider struct {
	Name string
}

type apiFixerResponse struct {
	Base string
	Date string
	Rates map[string]float64
}


// GetRatesFromAPI is the implementation of the fetchmethod
func (provider *APIFixerProvider) GetRatesFromAPI(pms *Parameters) (*Rates, error) {
	var URL = "http://api.fixer.io/"
	var date string
	// setting up parameters
	getParams := make([]string, 0)
	if pms.Millitimestamp == 0 {
		date = "latest"
	} else {
		t := time.Unix(0, pms.Millitimestamp * 1000000)
		date = t.Format("2006-01-02")
	}
	if pms.From != "" {
		getParams = append(getParams, fmt.Sprintf("base=%s", pms.From))
	}
	if len(pms.To) > 0 {
		getParams = append(getParams, fmt.Sprintf("symbols=%s", strings.Join(pms.To, ",")))
	}
	URL = fmt.Sprintf("%s%s", URL, date)
	if len(getParams) > 0 {
		URL += "?" + strings.Join(getParams, "&")
	}
	// actual request
	rawResp, err := Get(URL)
	if err != nil {
		return nil, err
	}
	// response handling
	var resp apiFixerResponse
	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		return nil, err
	}
	rates := Rates{
		Base:  resp.Base,
		// Date:  resp.Date,
		Currencies: resp.Rates,
	}
	// rates := Rates{}
	return &rates, nil
}
