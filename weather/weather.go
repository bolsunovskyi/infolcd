package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Listener interface {
	Update(r *Response)
	GetCityID() int
	GetAPIKey() string
}

const url string = "http://api.openweathermap.org/data/2.5/weather?id=%d&apikey=%s&units=metric"

type Weather struct {
	Main        string
	Description string
}

type Main struct {
	Temp     float32
	Pressure float32
	Humidity float32
}

type Response struct {
	Weather []Weather
	Main    Main
}

func Listen(l Listener) {
	go func() {
		t := time.NewTicker(time.Minute)
		for range t.C {
			rsp, err := GetWeather(l.GetCityID(), l.GetAPIKey())
			if err == nil {
				l.Update(rsp)
			}
		}
	}()
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
