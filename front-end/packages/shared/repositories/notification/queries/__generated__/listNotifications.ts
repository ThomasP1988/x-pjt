/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { NotificatonType } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL query operation: listNotifications
// ====================================================

export interface listNotifications_listNotifications_notifications {
  __typename: "Notification";
  id: string;
  userId: string;
  type: NotificatonType;
  createdAt: string;
  message: string;
  read: boolean;
}

export interface listNotifications_listNotifications {
  __typename: "ListNotification";
  notifications: (listNotifications_listNotifications_notifications | null)[] | null;
  next: string | null;
}

export interface listNotifications {
  listNotifications: listNotifications_listNotifications | null;
}

export interface listNotificationsVariables {
  from?: string | null;
  limit?: number | null;
}
