package temp

import (
	"github.com/d2r2/go-dht"
	"time"
)

type Listener interface {
	Update(r *Response)
	GetGPIO() int
}

type Response struct {
	Temp     float32
	Humidity float32
}

func Listen(l Listener) {
	go func() {
		t := time.NewTicker(time.Second * 30)
		for range t.C {
			t, h, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, l.GetGPIO(), true, 10)
			if err == nil {
				l.Update(&Response{
					Temp:     t,
					Humidity: h,
				})
			}
		}
	}()
}
