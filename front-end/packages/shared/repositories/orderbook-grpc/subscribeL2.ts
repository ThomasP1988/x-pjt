import { SubscribeL2Args, Prices } from '../../grpc/orderbook';
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import { GetOrderbookClient } from ".";

type SubscribeL2FcArgs = {
    symbol: string
}

export const SubscribeL2 = ({ symbol }: SubscribeL2FcArgs): ServerStreamingCall<SubscribeL2Args, Prices> => {
    const args: SubscribeL2Args = {
        symbol
    };

    GetOrderbookClient().subscribeL2(args)
    const stream : ServerStreamingCall<SubscribeL2Args, Prices> = GetOrderbookClient().subscribeL2(args);
    return stream;
}