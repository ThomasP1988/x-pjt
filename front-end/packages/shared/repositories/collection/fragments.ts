import gql from 'graphql-tag';

export const collectionFragment = gql`
    fragment Collection on Collection {
        id
        symbol
        description
        name
        supply
        chain
        imagePath
        thumbnailPath
        firstItemId
        status
        openseaSlug
    }
`;
