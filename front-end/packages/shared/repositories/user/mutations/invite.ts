import gql from 'graphql-tag';

export const INVITE_USER = gql`
    mutation inviteUser($email: String!) {
        inviteUser(email: $email) {
           success
        }
    }
`;