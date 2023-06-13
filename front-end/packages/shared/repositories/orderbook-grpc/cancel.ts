
import { Metadata, RpcError } from 'grpc-web';
import { CancelOrderArgs, Order } from '../../grpc/orderbook_pb';

import { Auth } from 'aws-amplify';
import { CognitoUserSession } from 'amazon-cognito-identity-js';

import { GetOrderbookClient } from ".";

export const CancelOrder = async (args: CancelOrderArgs): Promise<Order> => {

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
        GetOrderbookClient().cancelOrder(args, metadata, (error: RpcError | null, responseMessage: Order | null): void => {
            if (error) {
                console.log("error", error);
                reject(error);
            } else {
                resolve(responseMessage as Order);
            }
        });
    })

}
