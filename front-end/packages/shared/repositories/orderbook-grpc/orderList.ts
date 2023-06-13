
import { Metadata } from 'grpc-web';
import { OrderListResult, OrderListArgs } from '../../grpc/orderbook';


import { Auth } from 'aws-amplify';
import { CognitoUserSession } from 'amazon-cognito-identity-js';

import { GetOrderbookClient } from ".";

type RequestOrderListArgs = {
    symbol?: string,
    from?: string,
    limit?: number,
    isOpen?: boolean
}

export const RequestOrderList = async ({ symbol, from, limit, isOpen }: RequestOrderListArgs): Promise<OrderListResult> => {
    let session: CognitoUserSession | undefined;
    try {
        session = await Auth.currentSession()
    } catch (e) {
        return Promise.reject(e)
    }
    if (!session) {
        return Promise.reject("unauthenticated user")
    }

    const args: Partial<OrderListArgs> = {};

    if (symbol) {
        args.symbol = symbol;
    }

    if (limit) {
        args.limit = limit;
    }

    if (from) {
        args.from = from;
    }

    if (isOpen) {
        args.isOpen = isOpen;
    }

    const metadata: Metadata = {
        "jwt": session.getAccessToken().getJwtToken()
    };
    
    try {
        const unaryResponse = await GetOrderbookClient().orderList(args as OrderListArgs, metadata);
        return unaryResponse.response
    } catch(e) {
        console.log(e);
        return Promise.reject(e);
    }
}