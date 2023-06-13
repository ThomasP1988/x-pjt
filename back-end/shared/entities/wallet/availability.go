package wallet

import "fmt"

func (w *Wallet) HasEnough(currency string, amount *int64) bool {
	if asset, ok := w.Assets[currency]; ok {
		return asset.Own >= *amount
	}

	return false
}

func (w *Wallet) HasEnoughAvailability(currency string, amount *int64) bool {

	fmt.Printf("currency: %v\n", currency)
	fmt.Printf("amount: %v\n", *amount)

	if asset, ok := w.Assets[currency]; ok {
		return asset.Available >= *amount
	}
	return false
}
