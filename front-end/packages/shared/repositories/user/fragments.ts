import gql from 'graphql-tag';

export const userFragment = gql`
    fragment User on User {
        id,
        email,
        createdAt,
        lastSeenNotification
    }
`;
