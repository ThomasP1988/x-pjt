package blockchain

import (
	"NFTM/shared/config"
	"context"
	"os"
	"testing"
)

func setup() {
	config.GetConfig(nil)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {
	// "0xaadc2d4261199ce24a4b0a57370c4fcf43bb60aa", // THE CURRENCY
	// "0x3bf99d504e67a977f88b417ab68d34915f3a1209", // EMPRESSES
	// asset: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", // CRYPTO PUNKS
	FetchCollection(context.TODO(), "0x3bf99d504e67a977f88b417ab68d34915f3a1209")

	t.Fail()

}
