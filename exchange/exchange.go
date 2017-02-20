package exchange

import (
	"time"
	"github.com/bolsunovskyi/pb_api"
	"errors"
)

type Listener interface {
	Update(r *pb_api.ExchangeRate)
	GetPBID() string
	GetPBSecret() string
}

func Listen(l Listener) {
	go func() {
		t := time.NewTicker(time.Minute * 5)
		for range t.C {
			rsp, err := GetUSD(l.GetPBID(), l.GetPBSecret())
			//TODO: think about this section
			if err == nil {
				l.Update(rsp)
			}
		}
	}()
}

func GetUSD(pbid string, pbsecret string) (*pb_api.ExchangeRate, error) {
	pb_api.Init(pbid, pbsecret)
	ss, err := pb_api.SessionCreate()
	if err != nil {
		return nil, err
	}
	defer pb_api.SessionRemove(ss.ID)

	rates, err := pb_api.GetExchangeRate(pb_api.RATE_PB, ss.ID)
	if err != nil {
		return nil, err
	}

	for _, v := range *rates {
		if v.ExchangeRate.CCY == "USD" {
			return &v.ExchangeRate, nil
		}
	}

	return nil, errors.New("Unexpected error")
}
