package orderbooks

import (
	"NFTM/shared/config"
	wallet_entity "NFTM/shared/entities/wallet"
	price_helper "NFTM/shared/libs/price"
	orderbook "NFTM/shared/orderbook/grpc"
	"NFTM/shared/repositories/order"
	"NFTM/shared/repositories/wallet"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var userIdTest string = "userIdTest"
var assetIDStableCoin = "DAI"
var assetIDNFT = "CRYPTO"

func setup() {
	conf := config.GetConfig(nil)
	conf.Markets[0].DNS = "127.0.0.1"
	conf.Markets[0].Port = 50711
}

func tearDown() {
	config.GetConfig(nil)
	ctx := context.Background()
	err := wallet.DeleteByUser(ctx, userIdTest)
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

func resetWallet(t *testing.T, ctx context.Context) (int64, int64) {
	dai, _ := price_helper.StringToIntWithAppCoef("200")
	crypto, _ := price_helper.StringToIntWithAppCoef("5")

	err := wallet.DeleteByUser(ctx, userIdTest)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = wallet.GetWallet(ctx, userIdTest, []string{
		assetIDStableCoin,
		assetIDNFT,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	err = wallet.Add(ctx, &wallet_entity.WalletAsset{
		UserID:    userIdTest,
		AssetID:   assetIDStableCoin,
		Own:       dai,
		Available: dai,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = wallet.Add(ctx, &wallet_entity.WalletAsset{
		UserID:    userIdTest,
		AssetID:   assetIDNFT,
		Own:       crypto,
		Available: crypto,
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	return crypto, dai
}

func TestCancelOrder(t *testing.T) {
	OrderbooksService := NewOrderbooksService()
	OrderbooksService.Start(config.Conf.Markets[0].Pair.Symbol())

	cancelLimitOrder(t, context.Background(), OrderbooksService, "CRYPTO-DAI", "b556c144-b3bb-4e14-b508-ee5517d901e3")
}

func cancelLimitOrder(t *testing.T, ctx context.Context, os *OrderbooksService, symbol string, orderID string) {

	err := os.CancelOrder(ctx, symbol, orderID)

	if err != nil {
		t.Fail()
	}

}

func TestLimitBuy(t *testing.T) {
	OrderbooksService := NewOrderbooksService()
	OrderbooksService.Start(config.Conf.Markets[0].Pair.Symbol())
	ctx := context.Background()

	// a. prepare wallet

	crypto, dai := resetWallet(t, ctx)

	// b. prepare order

	price := "20"
	priceInt, err := price_helper.StringToIntWithAppCoef(price)
	if err != nil {
		fmt.Printf("err TestLimitBuy priceInt: %v\n", err)
		t.Fail()
	}

	amount := "1"
	amountInt, err := price_helper.StringToIntWithAppCoef(amount)
	if err != nil {
		fmt.Printf("err TestLimitBuy priceInt: %v\n", err)
		t.Fail()
	}

	// c. pass order

	result, err := OrderbooksService.ProcessLimitOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_BUY, amount, price)

	if err != nil {
		fmt.Printf("err  ProcessLimitOrder: %v\n", err)
		t.Fail()
		return
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("result: %v\n", result)

	// 1 Check saved order

	dbOrder, err := order.Get(ctx, result.OrderId)
	if err != nil {
		fmt.Printf("err  GetOrderService().Get: %v\n", err)
		t.Fail()
	}

	cancelLimitOrder(t, ctx, OrderbooksService, result.Symbol, result.OrderId)

	if dbOrder.OriginalQuantity != amountInt {
		fmt.Printf("\"  quantity different\": %v\n", dbOrder.OriginalQuantity)
		t.Fail()
	}

	if dbOrder.Price != priceInt {
		fmt.Printf("\"  price different\": %v\n", dbOrder.OriginalQuantity)
		t.Fail()
	}

	// 2 check wallet update

	updatedWallet, err := wallet.GetWallet(ctx, userIdTest, []string{
		assetIDStableCoin,
		assetIDNFT,
	})

	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}

	if updatedWallet.Assets["DAI"].Available != dai-priceInt {
		fmt.Printf("\"wallet dai result price expected\": %v\n", dai-amountInt)
		fmt.Printf("\"wallet dai result price found\": %v\n", updatedWallet.Assets["DAI"].Available)
		t.Fail()
	}

	if updatedWallet.Assets["CRYPTO"].Available != crypto {
		fmt.Printf("\"wallet crypto result price expected\": %v\n", dai-amountInt)
		fmt.Printf("\"wallet crypto result price found\": %v\n", updatedWallet.Assets["CRYPTO"].Available)
		t.Fail()
	}

}

func TestLimitSell(t *testing.T) {

	OrderbooksService := NewOrderbooksService()
	OrderbooksService.Start(config.Conf.Markets[0].Pair.Symbol())
	ctx := context.Background()

	// a. prepare wallet

	crypto, dai := resetWallet(t, ctx)

	// b. prepare order

	price := "20"
	priceInt, err := price_helper.StringToIntWithAppCoef(price)
	if err != nil {
		fmt.Printf("err TestLimitBuy priceInt: %v\n", err)
		t.Fail()
	}

	amount := "1"
	amountInt, err := price_helper.StringToIntWithAppCoef(amount)
	if err != nil {
		fmt.Printf("err TestLimitBuy priceInt: %v\n", err)
		t.Fail()
	}

	// c. pass order

	result, err := OrderbooksService.ProcessLimitOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_SELL, amount, price)

	if err != nil {
		fmt.Printf("err  ProcessLimitOrder: %v\n", err)
		t.Fail()
	}
	time.Sleep(time.Second * 3)

	fmt.Printf("result: %v\n", result)

	// 1 Check saved order

	dbOrder, err := order.Get(ctx, result.OrderId)
	if err != nil {
		fmt.Printf("err  GetOrderService().Get: %v\n", err)
		t.Fail()
	}

	cancelLimitOrder(t, ctx, OrderbooksService, result.Symbol, result.OrderId)

	if dbOrder.OriginalQuantity != amountInt {
		fmt.Printf("\"  quantity different\": %v\n", dbOrder.OriginalQuantity)
		t.Fail()
	}

	if dbOrder.Price != priceInt {
		fmt.Printf("\"  price different\": %v\n", dbOrder.OriginalQuantity)
		t.Fail()
	}

	// 2 check wallet update

	updatedWallet, err := wallet.GetWallet(ctx, userIdTest, []string{
		assetIDStableCoin,
		assetIDNFT,
	})

	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}

	if updatedWallet.Assets["DAI"].Available != dai {
		fmt.Printf("\"wallet result price expected\": %v\n", dai-amountInt)
		fmt.Printf("\"wallet result price found\": %v\n", updatedWallet.Assets["DAI"].Available)
		t.Fail()
	}
	if updatedWallet.Assets["CRYPTO"].Available != crypto-amountInt {
		fmt.Printf("\"wallet result price expected\": %v\n", dai-amountInt)
		fmt.Printf("\"wallet result price found\": %v\n", updatedWallet.Assets["DAI"].Available)
		t.Fail()
	}
}

func TestMarketBuy(t *testing.T) {
	OrderbooksService := NewOrderbooksService()
	OrderbooksService.Start(config.Conf.Markets[0].Pair.Symbol())
	ctx := context.Background()

	var limitOrder1ID, limitOrder2ID, marketOrderID string
	var symbol string = "CRYPTO-DAI"

	defer func() {
		if limitOrder1ID != "" {
			cancelLimitOrder(t, ctx, OrderbooksService, symbol, limitOrder1ID)
		}
		if limitOrder2ID != "" {
			cancelLimitOrder(t, ctx, OrderbooksService, symbol, limitOrder2ID)
		}
		if marketOrderID != "" {
			cancelLimitOrder(t, ctx, OrderbooksService, symbol, marketOrderID)
		}
	}()

	// a. prepare wallet

	crypto, dai := resetWallet(t, ctx)

	// first limit order
	priceOrder1 := "20"
	priceOrder1Int, _ := price_helper.StringToIntWithAppCoef(priceOrder1)
	amountOrder1 := "1"
	amountOrder1Int, _ := price_helper.StringToIntWithAppCoef(amountOrder1)
	fmt.Printf("amountOrder1Int: %v\n", amountOrder1Int)
	resultLimitOrder1, err := OrderbooksService.ProcessLimitOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_BUY, amountOrder1, priceOrder1)
	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}
	limitOrder1ID = resultLimitOrder1.OrderId
	// second limit order
	priceOrder2 := "30"
	priceOrder2Int, _ := price_helper.StringToIntWithAppCoef(priceOrder2)
	amountOrder2 := "2"
	amountOrder2Int, _ := price_helper.StringToIntWithAppCoef(amountOrder2)
	fmt.Printf("amountOrder2Int: %v\n", amountOrder2Int)

	resultLimitOrder2, err := OrderbooksService.ProcessLimitOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_BUY, amountOrder2, priceOrder2)
	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}
	limitOrder2ID = resultLimitOrder2.OrderId
	// market order, we want to buy fully second order and half of first order with 2 as amount
	amountMarketOrder := "2"
	amountMarketOrderInt, _ := price_helper.StringToIntWithAppCoef(amountMarketOrder)
	fmt.Printf("amountMarketOrderInt: %v\n", amountMarketOrderInt)
	resultMarketOrder, err := OrderbooksService.ProcessMarketOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_SELL, amountMarketOrder)
	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}
	fmt.Printf("resultMarketOrder: %v\n", resultMarketOrder)

	marketOrderID = resultMarketOrder.OrderId

	limitOrder1, err := order.Get(ctx, resultLimitOrder1.OrderId)
	if err != nil {
		fmt.Printf("err  GetOrderService().Get: %v\n", err)
		t.Fail()
	}

	if limitOrder1.Quantity != price_helper.FromIntToIntWithAppCoef(1) {
		fmt.Printf("wrong limitOrder1.Quantity: %v\n", limitOrder1.Quantity)
		t.Fail()
	}

	limitOrder2, err := order.Get(ctx, resultLimitOrder1.OrderId)
	if err != nil {
		fmt.Printf("err  GetOrderService().Get: %v\n", err)
		t.Fail()
	}

	if limitOrder2.Quantity != 0 {
		fmt.Printf("wrong limitOrder2.Quantity: %v\n", limitOrder2.Quantity)
		t.Fail()
	}

	if limitOrder2.FilledQuantity != price_helper.FromIntToIntWithAppCoef(1) {
		fmt.Printf("wrong limitOrder2.FilledQuantity: %v\n", limitOrder2.Quantity)
		t.Fail()
	}

	expectedCumulative := priceOrder1Int + priceOrder2Int

	if *resultMarketOrder.CumulativeQuoteQuantity != expectedCumulative {
		fmt.Printf("expectedCumulative: %v\n", expectedCumulative)
		fmt.Printf("resultMarketOrder.CumulativeQuoteQuantity: %v\n", *resultMarketOrder.CumulativeQuoteQuantity)
		t.Fail()
	}

	if *resultMarketOrder.ExecutedQuantity != price_helper.FromIntToIntWithAppCoef(2) {
		fmt.Printf("resultMarketOrder.ExecutedQuantity: %v\n", resultMarketOrder.ExecutedQuantity)
		fmt.Printf("price_helper.FromIntToIntWithAppCoef(2): %v\n", price_helper.FromIntToIntWithAppCoef(2))
		t.Fail()
	}

	updatedWallet, err := wallet.GetWallet(ctx, userIdTest, []string{
		assetIDStableCoin,
		assetIDNFT,
	})

	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}

	if updatedWallet.Assets["DAI"].Available != dai-(priceOrder1Int+priceOrder2Int) {
		fmt.Printf("\"wallet result dai expected\": %v\n", dai-(priceOrder1Int+priceOrder2Int))
		fmt.Printf("\"wallet result dai found\": %v\n", updatedWallet.Assets["DAI"].Available)
		t.Fail()
	}

	if updatedWallet.Assets["CRYPTO"].Available != crypto-1*price_helper.MultiplyBy {
		fmt.Printf("\"wallet result crypto expected\": %v\n", updatedWallet.Assets["CRYPTO"].Available)
		fmt.Printf("\"wallet result crypto found\": %v\n", crypto-1)
		t.Fail()
	}

}

func TestMarketSell(t *testing.T) {
	OrderbooksService := NewOrderbooksService()
	OrderbooksService.Start(config.Conf.Markets[0].Pair.Symbol())
	ctx := context.Background()

	var limitOrder1ID, limitOrder2ID, marketOrderID string
	var symbol string = "CRYPTO-DAI"

	defer func() {
		if limitOrder1ID != "" {
			cancelLimitOrder(t, ctx, OrderbooksService, symbol, limitOrder1ID)
		}
		if limitOrder2ID != "" {
			cancelLimitOrder(t, ctx, OrderbooksService, symbol, limitOrder2ID)
		}
		if marketOrderID != "" {
			cancelLimitOrder(t, ctx, OrderbooksService, symbol, marketOrderID)
		}
	}()

	// a. prepare wallet

	crypto, dai := resetWallet(t, ctx)

	// first limit order
	priceOrder1 := "20"
	priceOrder1Int, _ := price_helper.StringToIntWithAppCoef(priceOrder1)
	amountOrder1 := "1"
	amountOrder1Int, _ := price_helper.StringToIntWithAppCoef(amountOrder1)
	fmt.Printf("amountOrder1Int: %v\n", amountOrder1Int)
	resultLimitOrder1, err := OrderbooksService.ProcessLimitOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_SELL, amountOrder1, priceOrder1)
	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}
	limitOrder1ID = resultLimitOrder1.OrderId
	// second limit order
	priceOrder2 := "30"
	priceOrder2Int, _ := price_helper.StringToIntWithAppCoef(priceOrder2)
	amountOrder2 := "2"
	amountOrder2Int, _ := price_helper.StringToIntWithAppCoef(amountOrder2)
	fmt.Printf("amountOrder2Int: %v\n", amountOrder2Int)

	resultLimitOrder2, err := OrderbooksService.ProcessLimitOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_SELL, amountOrder2, priceOrder2)
	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}
	limitOrder2ID = resultLimitOrder2.OrderId
	// market order, we want to buy fully second order and half of first order with 2 as amount
	amountMarketOrder := "2"
	amountMarketOrderInt, _ := price_helper.StringToIntWithAppCoef(amountMarketOrder)
	fmt.Printf("amountMarketOrderInt: %v\n", amountMarketOrderInt)
	resultMarketOrder, err := OrderbooksService.ProcessMarketOrder(ctx, userIdTest, &config.Conf.Markets[0].Pair, orderbook.SideOrder_BUY, amountMarketOrder)
	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}
	fmt.Printf("resultMarketOrder: %v\n", resultMarketOrder)

	marketOrderID = resultMarketOrder.OrderId

	limitOrder1, err := order.Get(ctx, resultLimitOrder1.OrderId)
	if err != nil {
		fmt.Printf("err  GetOrderService().Get: %v\n", err)
		t.Fail()
	}

	if limitOrder1.Quantity != price_helper.FromIntToIntWithAppCoef(1) {
		fmt.Printf("wrong limitOrder1.Quantity: %v\n", limitOrder1.Quantity)
		t.Fail()
	}

	limitOrder2, err := order.Get(ctx, resultLimitOrder1.OrderId)
	if err != nil {
		fmt.Printf("err  GetOrderService().Get: %v\n", err)
		t.Fail()
	}

	if limitOrder2.Quantity != 0 {
		fmt.Printf("wrong limitOrder2.Quantity: %v\n", limitOrder2.Quantity)
		t.Fail()
	}

	if limitOrder2.FilledQuantity != price_helper.FromIntToIntWithAppCoef(1) {
		fmt.Printf("wrong limitOrder2.Quantity: %v\n", limitOrder2.Quantity)
		t.Fail()
	}

	expectedCumulative := priceOrder1Int + priceOrder2Int

	if *resultMarketOrder.CumulativeQuoteQuantity != expectedCumulative {
		fmt.Printf("expectedCumulative: %v\n", expectedCumulative)
		fmt.Printf("resultMarketOrder.CumulativeQuoteQuantity: %v\n", *resultMarketOrder.CumulativeQuoteQuantity)
		t.Fail()
	}

	if *resultMarketOrder.ExecutedQuantity != price_helper.FromIntToIntWithAppCoef(2) {
		fmt.Printf("resultMarketOrder.ExecutedQuantity: %v\n", resultMarketOrder.ExecutedQuantity)
		fmt.Printf("price_helper.FromIntToIntWithAppCoef(2): %v\n", price_helper.FromIntToIntWithAppCoef(2))
		t.Fail()
	}

	updatedWallet, err := wallet.GetWallet(ctx, userIdTest, []string{
		assetIDStableCoin,
		assetIDNFT,
	})

	if err != nil {
		fmt.Printf("\"  error finding wallet \": %v\n", err)
		t.Fail()
	}

	if updatedWallet.Assets["DAI"].Available != dai-(priceOrder1Int+priceOrder2Int) {
		fmt.Printf("\"wallet result dai expected\": %v\n", dai-(priceOrder1Int+priceOrder2Int))
		fmt.Printf("\"wallet result dai found\": %v\n", updatedWallet.Assets["DAI"].Available)
		t.Fail()
	}

	if updatedWallet.Assets["CRYPTO"].Available != crypto-1*price_helper.MultiplyBy {
		fmt.Printf("\"wallet result crypto expected\": %v\n", updatedWallet.Assets["CRYPTO"].Available)
		fmt.Printf("\"wallet result crypto found\": %v\n", crypto-1)
		t.Fail()
	}
}
