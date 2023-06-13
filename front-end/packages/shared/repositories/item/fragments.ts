import gql from 'graphql-tag';

export const itemFragment = gql`
    fragment Item on Item {
        id,
        tokenURI,
        isFetching,
        collectionAddress,
        image_data,
        name,
        image,
        imagePath,
        thumbnailPath,
        description,
        external_url,
        background_color,
        animation_url,
        youtube_url,
        attributes {
            display_type,
            trait_type,
            value,
            color,
        }
    }
`;
