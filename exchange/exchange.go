package exchange

import "net/http"

const url string = "https://obmenka.kharkov.ua/api/rates/020016"

type LatestRates struct {
	WholeSale	float32
	WholeBuy	float32
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
	LatestRates	LatestRates
	BankRates	BankRates
}

func GetUSD() (*string, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

}
