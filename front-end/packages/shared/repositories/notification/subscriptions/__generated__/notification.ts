/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { NotificatonType } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL subscription operation: notification
// ====================================================

export interface notification_notification {
  __typename: "Notification";
  id: string;
  userId: string;
  type: NotificatonType;
  createdAt: string;
  message: string;
  read: boolean;
}

export interface notification {
  notification: notification_notification | null;
}

export interface notificationVariables {
  userId: string;
}
