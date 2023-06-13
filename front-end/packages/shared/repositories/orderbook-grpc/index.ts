import { OrderbooksClient } from '../../grpc/orderbook.client';
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
export * from "./orderList";
export * from "./cancel";

export const GetOrderbookClient = (): OrderbooksClient => {
    // TODO : check if https is enough for securisation
    // 'http://example.api/markets'

    let transport = new GrpcWebFetchTransport({
        baseUrl: "http://example.api/markets"
    });

    return new OrderbooksClient(transport);
}

