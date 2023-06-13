import gql from 'graphql-tag';
import { collectionFragment } from "../fragments";

export const SUBMIT_COLLECTION = gql`
    mutation submitCollection($address: String!, $description: String) {
        submitCollection(address: $address, description: $description) {
           ...Collection
        }
    }
    ${collectionFragment}
`;