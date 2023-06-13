
import { Metadata, RpcError } from 'grpc-web';
import { OrderArgs, Order } from '../../grpc/orderbook_pb';

import { Auth } from 'aws-amplify';
import { CognitoUserSession } from 'amazon-cognito-identity-js';

import { GetOrderbookClient } from ".";

export const NewOrderRequest = async (order: OrderArgs): Promise<Order> => {
    let session: CognitoUserSession | undefined;
    try {
        session = await Auth.currentSession()
    } catch (e) {
        return Promise.reject(e)
    }
    if (!session) {
        return Promise.reject("unauthenticated user")
    }

    const metadata: Metadata = {
        "jwt": session.getAccessToken().getJwtToken()
    };

    return new Promise((resolve, reject) => {
        GetOrderbookClient().processOrder(order, metadata, (error: RpcError | null, responseMessage: Order | null): void => {
            if (error) {
                console.log("error", error);
                reject(error);
            } else {
                resolve(responseMessage as Order);
            }
        });
    })
}