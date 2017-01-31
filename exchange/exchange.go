package exchange

import (
	"net/http"
	"encoding/json"
	"errors"
)

const url string = "https://obmenka.kharkov.ua/api/rates/020016"

type LatestRates struct {
	WholeSale	float32		`json:"wholeSale"`
	WholeBuy	float32		`json:"wholeBuy"`
}

type BankUnitRates struct {
	RetailSale	float32
	RetailBuy	float32
}

type BankRates struct {
	PrivatBankRates		BankUnitRates
	NBURates		BankUnitRates
}

type Response struct {
	ID		int
	From		string
	To		string
	LatestRates	LatestRates	`json:"latestRates"`
	BankRates	BankRates	`json:"bankRates"`
}

func GetUSD() (*Response, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	resps := make([]Response, 0)

	err = json.NewDecoder(rsp.Body).Decode(&resps)
	if err != nil {
		return nil, err
	}
	if len(resps) > 0 {
		return &resps[0], nil
	}

	return nil, errors.New("Empty response")
}
