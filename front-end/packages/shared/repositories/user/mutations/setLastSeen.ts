import gql from 'graphql-tag';
import { userFragment } from '../fragments';

export const SET_LAST_SEEN_NOTIFICATION = gql`
    mutation setLastSeenNotification {
        setLastSeenNotification {
            ...User
        }
    }
    ${userFragment}
`;