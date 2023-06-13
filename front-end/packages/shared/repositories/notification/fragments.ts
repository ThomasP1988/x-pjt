import gql from 'graphql-tag';

export const notificationFragment = gql`
    fragment Notification_notification on Notification {
        id,
        userId,
        type,
        createdAt,
        message,
        read
    }
`;
