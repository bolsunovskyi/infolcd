package exchange

import (
	"testing"
	"fmt"
)

func TestGetUSD(t *testing.T) {
	rsp, err := GetUSD()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v", rsp)
}
