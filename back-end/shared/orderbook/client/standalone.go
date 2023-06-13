package orderbook_client

import (
	pair_component "NFTM/shared/components/pair"
	ob_client "NFTM/shared/orderbook/grpc"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderbookStandAloneClient struct {
	conn   *grpc.ClientConn
	Client ob_client.OrderbookClient
	Pair   *pair_component.Pair
}

func NewOrderbookStandAloneClient(url string, pair pair_component.Pair) (*OrderbookStandAloneClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	newClient := OrderbookStandAloneClient{
		conn:   conn,
		Client: ob_client.NewOrderbookClient(conn),
		Pair:   &pair,
	}

	for map[string]bool{
		"IDLE":              true,
		"CONNECTING":        true,
		"TRANSIENT_FAILURE": true,
	}[conn.GetState().String()] {
		fmt.Printf("current status: %v\n", conn.GetState().String())
		conn.WaitForStateChange(context.Background(), conn.GetState())
	}
	fmt.Printf("orderbook connection status: %v\n", conn.GetState())

	return &newClient, nil
}

func (obc *OrderbookStandAloneClient) StartL2() (*ob_client.Orderbook_SubscribeL2Client, error) {
	// get data from orderbook micro service
	stream, err := obc.Client.SubscribeL2(context.Background(), &ob_client.SubscribeArgs{
		Pair: obc.Pair.Symbol(),
	})
	if err != nil {
		return nil, err
	}

	return &stream, nil
}

func (obc *OrderbookStandAloneClient) Stop() {
	obc.conn.Close()
}
