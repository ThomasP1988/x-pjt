package price_helper

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestToIntWithAppCoef(t *testing.T) {

	newdec := decimal.New(1, 0)
	result := ToIntWithAppCoef(&newdec)

	if result != 100000000 {
		t.Fail()
	}
}
func TestFromIntWithAppCoef(t *testing.T) {

	var newint int64 = 100003454
	result := FromIntWithAppCoef(newint)
	fmt.Printf("result.String(): %v\n", result.String())
	if result.Equal(decimal.New(1, 0)) {
		t.Fail()
	}
}
func TestToFromIntWithAppCoef(t *testing.T) {
	var original int64 = 2353463634
	newdec := decimal.New(original, 0)
	intResult := ToIntWithAppCoef(&newdec)
	result := FromIntWithAppCoef(intResult)
	fmt.Printf("result: %v\n", result)
	if !result.Equal(decimal.New(original, 0)) {
		t.Fail()
	}
}
