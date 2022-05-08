package data

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strconv"
)

type ExchangeRateser interface {
	getRate() error
}

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewExchangeRates(log hclog.Logger) (*ExchangeRates, error) {
	err := &ExchangeRates{log: log, rates: map[string]float64{}}

	return err, nil
}

func (e *ExchangeRates) getRate() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &Cubes{}
	err = xml.NewDecoder(resp.Body).Decode(md)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[c.Currency] = r
	}

	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
