package trade

import (
	price_helper "NFTM/shared/libs/price"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var pair string = "CRYPTO-DAI"

func TestTrade(t *testing.T) {
	t.Log("1")
	err := SetClient()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
		return
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("i: %v\n", i)
		// randomTime := rand.Int63n(time.Now().UnixNano()-int64(time.Hour)) + time.Now().Unix()

		// randomNow := time.Unix(randomTime, 0)

		Trade(pair, strconv.Itoa(i), price_helper.FromIntToIntWithAppCoef(500), price_helper.FromIntToIntWithAppCoef(100), strconv.Itoa(rand.Intn(1)), time.Now())
	}

	time.Sleep(time.Second * 10)
	result2, err := Find(pair)
	if err != nil {
		t.Logf("err: %v\n", err)
		t.Fail()
	}
	fmt.Printf("result2: %v\n", result2)
	t.Fail()

}
