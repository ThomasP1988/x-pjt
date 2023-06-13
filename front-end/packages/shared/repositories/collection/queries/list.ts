import gql from 'graphql-tag';
import { collectionFragment } from '../fragments';

export const LIST_COLLECTIONS = gql`
    query listCollections($limit: Int, $from: String,  $status: CollectionStatus) {
        listCollections(limit: $limit, from: $from, status: $status) {
            collections {
                ...Collection
            }
            next
        }
    }
    ${collectionFragment}
`;