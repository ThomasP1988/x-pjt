import gql from 'graphql-tag';
import { itemFragment } from '../fragments';

export const GET_ITEM = gql`
    query getItem($collectionAddress: String!, $tokenId: Int!) {
        getItem(collectionAddress: $collectionAddress, tokenId: $tokenId) {
            ...Item
        }
    }
    ${itemFragment}
`;