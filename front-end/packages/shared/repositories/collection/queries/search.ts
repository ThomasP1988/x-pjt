import gql from 'graphql-tag';
import { collectionFragment } from '../fragments';

export const SEARCH_COLLECTION = gql`
    query searchCollections($text: String!, $from: String,  $limit: Int) {
        searchCollections(text: $text, from: $from, limit: $limit) {
            results {
                ...Collection
            }
            total
        }
    }
    ${collectionFragment}
`;