package orderbook

import (
	ob_client "NFTM/shared/orderbook/grpc"
	"fmt"
)

type Broker struct {
	stopCh    chan struct{}
	publishCh chan *ob_client.Prices
	subCh     chan chan *ob_client.Prices
	unsubCh   chan chan *ob_client.Prices
	pair      string
}

// TODO:GENERIC
func NewBroker(pair string) *Broker {
	return &Broker{
		stopCh:    make(chan struct{}),
		publishCh: make(chan *ob_client.Prices, 1),
		subCh:     make(chan chan *ob_client.Prices, 1),
		unsubCh:   make(chan chan *ob_client.Prices, 1),
		pair:      pair,
	}
}

func (b *Broker) Start() {
	subs := map[chan *ob_client.Prices]ob_client.Prices{}
	for {
		select {
		case <-b.stopCh:
			return
		case msgCh := <-b.subCh:
			subs[msgCh] = ob_client.Prices{}
		case msgCh := <-b.unsubCh:
			delete(subs, msgCh)
			fmt.Printf("\"unsubscribe\": %v\n", len(subs))
			if len(subs) == 0 { // stop subscribing to ob
				b.Stop()
			}
		case msg := <-b.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	fmt.Printf("stop price broker: %v\n", b.pair)

	OBServices.StopL2(b.pair)
	close(b.stopCh)
	delete((*brokersL2), b.pair)
}

func (b *Broker) Subscribe() chan *ob_client.Prices {
	msgCh := make(chan *ob_client.Prices, 1)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan *ob_client.Prices) {
	b.unsubCh <- msgCh
}

func (b *Broker) Publish(msg *ob_client.Prices) {
	b.publishCh <- msg
}
