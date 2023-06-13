import gql from 'graphql-tag';

export const CREATE_TOKEN = gql`
    mutation createToken {
        createToken {
           token,
           nonce
        }
    }
`;