import gql from 'graphql-tag';
import { notificationFragment } from '../fragments';

export const LIST_NOTIFICATIONS = gql`
    query listNotifications($from: String, $limit: Int) {
        listNotifications(from: $from, limit: $limit) {
            notifications {
               ...Notification_notification
            }
            next
        }
    }
    ${notificationFragment}
`;