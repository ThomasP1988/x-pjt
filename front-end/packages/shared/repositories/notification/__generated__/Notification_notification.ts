/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { NotificatonType } from "./../../../__generated__/globalTypes";

// ====================================================
// GraphQL fragment: Notification_notification
// ====================================================

export interface Notification_notification {
  __typename: "Notification";
  id: string;
  userId: string;
  type: NotificatonType;
  createdAt: string;
  message: string;
  read: boolean;
}
