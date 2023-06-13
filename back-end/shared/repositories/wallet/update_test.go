package wallet

import (
	"NFTM/shared/config"
	"NFTM/shared/entities/wallet"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

var userIdTest string = "userIdTest"

var assetIDStableCoin = "DAI"
var assetIDNFT = "CRYPTO"

func setup() {
	config.GetConfig(nil)
	ctx := context.Background()

	err := Add(ctx, &wallet.WalletAsset{
		UserID:  userIdTest,
		AssetID: assetIDStableCoin,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = Add(ctx, &wallet.WalletAsset{
		UserID:  userIdTest,
		AssetID: assetIDNFT,
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

}

func tearDown() {
	config.GetConfig(nil)
	ctx := context.Background()
	err := Delete(ctx, userIdTest, assetIDStableCoin)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = Delete(ctx, userIdTest, assetIDNFT)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestUpdate(t *testing.T) {
	ctx := context.TODO()

	wallet, err := GetWallet(ctx, userIdTest, []string{assetIDStableCoin, assetIDNFT})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	if wallet.Assets[assetIDStableCoin].Own != 0 || wallet.Assets[assetIDNFT].Own != 0 ||
		wallet.Assets[assetIDStableCoin].Available != 0 || wallet.Assets[assetIDNFT].Available != 0 {
		t.Fail()
	}

	err = UpdateWallet(UpdateWalletArgs{
		Ctx:    ctx,
		UserID: userIdTest,
		Currencies: map[string]UpdateWalletCurrency{
			assetIDNFT: {
				Own:       5,
				Available: 5,
			},
			assetIDStableCoin: {
				Own:       5,
				Available: 5,
			},
		},
		Wallet: wallet,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	if wallet.Assets[assetIDStableCoin].Own != 5 || wallet.Assets[assetIDNFT].Own != 5 ||
		wallet.Assets[assetIDStableCoin].Available != 5 || wallet.Assets[assetIDNFT].Available != 5 {
		t.Fail()
	}

	err = UpdateWallet(UpdateWalletArgs{
		Ctx:    ctx,
		UserID: userIdTest,
		Currencies: map[string]UpdateWalletCurrency{
			assetIDNFT: {
				Own:       5,
				Available: 5,
			},
			assetIDStableCoin: {
				Own:       5,
				Available: 5,
			},
		},
		Wallet: wallet,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	wallet, err = GetWallet(ctx, userIdTest, []string{assetIDStableCoin, assetIDNFT})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	if wallet.Assets[assetIDStableCoin].Own != 10 || wallet.Assets[assetIDNFT].Own != 10 ||
		wallet.Assets[assetIDStableCoin].Available != 10 || wallet.Assets[assetIDNFT].Available != 10 {
		t.Fail()
	}

}
