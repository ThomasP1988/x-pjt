package price_helper

import "github.com/shopspring/decimal"

var Coef int32 = 8
var MultiplyBy int64 = 100000000

func StringToIntWithAppCoef(input string) (int64, error) {
	inputDec, err := decimal.NewFromString(input)
	if err != nil {
		return 0, nil
	}
	inputMulCoef := inputDec.Mul(decimal.New(MultiplyBy, 0))
	return inputMulCoef.IntPart(), nil
}

func ToIntWithAppCoef(input *decimal.Decimal) int64 {
	inputMulCoef := (*input).Mul(decimal.New(MultiplyBy, 0))
	return inputMulCoef.IntPart()
}

func FromIntWithAppCoef(input int64) decimal.Decimal {
	return decimal.New(input, -Coef)
}

func FromIntToIntWithAppCoef(input int64) int64 {
	return input * MultiplyBy
}

func FromIntToString(input int64) string {
	return FromIntWithAppCoef(input).String()
}

func FromIntToInexactFloat(input int64) float64 {
	return FromIntWithAppCoef(input).InexactFloat64()
}
