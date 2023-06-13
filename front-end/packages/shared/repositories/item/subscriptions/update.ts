import gql from 'graphql-tag';
import { itemFragment } from '../fragments';

export const onUpdatedItem = gql`
    subscription onUpdatedItem($collectionAddress: String!, $tokenId: Int!) {
        onUpdatedItem(collectionAddress: $collectionAddress, tokenId: $tokenId) {
            ...Item
        }
    }
    ${itemFragment}
`;