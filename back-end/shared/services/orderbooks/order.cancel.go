package orderbooks

import (
	orderbook "NFTM/shared/orderbook/grpc"
	"NFTM/shared/repositories/order"
	"context"
	"errors"
)

func (os *OrderbooksService) CancelOrder(ctx context.Context, symbol string, orderID string) error {

	//TODO: check if orderbook is up, if not, then do nothing

	result, err := os.OrderbookClients[symbol].Client.CancelOrder(ctx, &orderbook.CancelOrderArgs{
		Pair:    symbol,
		OrderId: orderID,
	})

	if err != nil {
		return err
	}

	_, err = order.Cancel(ctx, orderID)

	if err != nil {
		return err
	}

	if !result.Success {
		return errors.New("error canceling")
	}

	return nil
}
