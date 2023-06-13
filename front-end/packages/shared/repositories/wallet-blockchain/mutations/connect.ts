import gql from 'graphql-tag';

export const CONNECT_WALLET = gql`
    mutation connectWallet($signature: String, $tokenId: String) {
        connectWallet(signature: $signature, tokenId: $tokenId) {
            success
        }
    }
`;