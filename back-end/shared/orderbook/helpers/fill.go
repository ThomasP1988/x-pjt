package ob_helpers

import (
	"NFTM/shared/constants"
	"NFTM/shared/entities/order"
	price_helper "NFTM/shared/libs/price"
	order_repo "NFTM/shared/repositories/order"
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	ob "github.com/i25959341/orderbook"
)

func FillOrderbook(obook *ob.OrderBook, symbol string) {
	var wg sync.WaitGroup = sync.WaitGroup{}
	var from *string
	for {
		wg.Add(1)
		fmt.Printf("\"ici\": %v\n", "ici")
		orders, next, err := order_repo.OrderListBySymbol(context.Background(), order_repo.OrderListBySymbolArgs{
			Symbol: aws.String("CRYPTO-DAI"),
			IsOpen: aws.Int8(1),
			Limit:  aws.Int32(20),
			From:   from,
		})

		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Printf("\"la\": %v\n", "la")
		from = next

		go func() {
			defer wg.Done()
			OrdersFromIntWithAppCoefToOB(obook, orders)
		}()
		if next == nil || *next == "" {
			break
		}
		fmt.Printf("next: %v\n", *next)
	}

	wg.Wait()
}

func OrdersFromIntWithAppCoefToOB(obook *ob.OrderBook, orders *[]order.Order) {
	for _, item := range *orders {
		if item.Type == constants.Order_LIMIT {
			price := price_helper.FromIntWithAppCoef(item.Price)
			quantity := price_helper.FromIntWithAppCoef(item.Quantity)

			side := FormatSide(item.Side)
			fmt.Printf("item.ID: %v\n", item.ID)
			fmt.Printf("item.Side: %v\n", item.Side)
			fmt.Printf("side: %v\n", side)
			fmt.Printf("quantity: %v\n", quantity.String())
			fmt.Printf("price: %v\n", price.String())
			obook.ProcessLimitOrder(
				side,
				item.ID,
				quantity,
				price,
			)
		}
	}
}
