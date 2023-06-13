import gql from 'graphql-tag';
import { collectionFragment } from '../fragments';
import { itemFragment } from '../../item/fragments';

export const LIST_COLLECTIONS_AND_ITEMS = gql`
    query listCollectionsAndItems($ids: [String!]!, $itemKeys: [ItemKeys!]!) {
        listCollectionsByIds(ids: $ids) {
            ...Collection
        }
        listItemsByIds(keys: $itemKeys) {
            ...Item
        }
           
    }
    ${collectionFragment}
    ${itemFragment}
`;