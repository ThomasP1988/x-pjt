import gql from 'graphql-tag';
import { collectionFragment } from "../fragments";

export const VALIDATE_COLLECTION = gql`
    mutation validateCollection($address: String!, $description: String, $name: String, $image: String, $symbol: String, $openseaSlug: String, $status: CollectionStatus) {
        validateCollection(address: $address, description: $description, name: $name, image: $image, symbol: $symbol, openseaSlug: $openseaSlug, status: $status) {
           ...Collection
        }
    }
    ${collectionFragment}
`;