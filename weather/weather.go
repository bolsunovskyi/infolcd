package weather

import (
	"encoding/json"
	"net/http"
	"fmt"
)

const key string = "26f430f11e19862ae62c3b11fbea4c37"
const city string = "703448"
const url string = "http://api.openweathermap.org/data/2.5/weather?id=%d&apikey=%s&units=metric"

type Weather struct {
	Main		string
	Description	string
}

type Main struct {
	Temp		float32
	Pressure	float32
	Humidity	float32
}

type Response struct {
	Weather		[]Weather
	Main		Main
}

func GetWeather(cityID int, apiKey string) (*Response, error) {
	rsp, err := http.Get(fmt.Sprintf(url, cityID, apiKey))
	if err != nil {
		return nil, err
	}
	resp := Response{}

	err = json.NewDecoder(rsp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}