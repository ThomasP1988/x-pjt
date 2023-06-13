import gql from 'graphql-tag';
import { notificationFragment } from '../fragments';

export const onAddedNotification = gql`
    subscription notification($userId: String!) {
        notification(userId: $userId) {
            ...Notification_notification
        }
    }
    ${notificationFragment}
`;