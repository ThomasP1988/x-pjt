import React, { useEffect, useState } from 'react';
import Grid from '@mui/material/Grid';
import { NewOrder } from './NewOrder';
import { PriceLevel, Prices, SubscribeL2Args } from '../../grpc/orderbook';
import { Prices as PricesDiplay } from "./prices";
import { BidAsk } from "../../constants";
import { SubscribeL2 } from '../../repositories/orderbook-grpc/subscribeL2';
import { RpcOutputStream, ServerStreamingCall } from "@protobuf-ts/runtime-rpc";

type Props = {
    symbol: string
}

export const Orderbook = ({ symbol }: Props) => {

    const [client, setClient] = useState<ServerStreamingCall<SubscribeL2Args, Prices> | null>(null);
    const [bids, setBids] = useState<PriceLevel[] | null>(null);
    const [asks, setAsks] = useState<PriceLevel[] | null>(null);

    useEffect(() => {
        if (!client) {
            const output: ServerStreamingCall<SubscribeL2Args, Prices> = SubscribeL2({ symbol });

            const stream: RpcOutputStream<Prices> = output.responses;

            stream.onMessage((response: Prices) => {
                // console.log("responses", response);
                setAsks(response.asks.reverse());
                setBids(response.bids);
            });

            stream.onError((error: Error) => {
                console.log("error", error)
                // setClient(null);
            });
            stream.onComplete(() => {
                setClient(null);
            });
         
            setClient(output);
        }
        return () => {
            if (client) {
                client.trailers
            }
        }
    }, [client, setClient, symbol]);


    return <Grid container>
        <Grid item xs={12}>
            <NewOrder symbol={symbol} />
        </Grid>
        <Grid item xs={12}>
            {asks && <PricesDiplay bidOrAsk={BidAsk.Ask} prices={asks} />}
            {bids && <PricesDiplay bidOrAsk={BidAsk.Bid} prices={bids} hideHeader={true} />}
        </Grid>
    </Grid>
}
