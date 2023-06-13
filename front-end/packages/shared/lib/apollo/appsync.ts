
import { AuthOptions, AUTH_TYPE, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";

import {
  ApolloProvider,
  ApolloClient,
  InMemoryCache,
  HttpLink,
  ApolloLink,
  NormalizedCacheObject,
} from "@apollo/client";
import { Auth } from '@aws-amplify/auth'

export const GetClient = (url: string, region: string,): ApolloClient<NormalizedCacheObject> => {
    const auth: AuthOptions = {
        type: AUTH_TYPE.AWS_IAM,
        credentials: async () => Auth.currentCredentials()
      };
      
      const httpLink = new HttpLink({ uri: url });
      
      const link = ApolloLink.from([
        createAuthLink({ url, region, auth }),
        createSubscriptionHandshakeLink({ url, region, auth }, httpLink),
      ]);
      
      const client = new ApolloClient({
        link,
        cache: new InMemoryCache()
      });
      return client;
}