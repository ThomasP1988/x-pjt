import gql from 'graphql-tag';
import { userFragment } from '../fragments';

export const ME = gql`
    query me {
        me {
            ...User
        }
    }
    ${userFragment}
`;