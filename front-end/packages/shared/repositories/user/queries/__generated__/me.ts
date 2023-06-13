/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: me
// ====================================================

export interface me_me {
  __typename: "User";
  id: string;
  email: string;
  createdAt: string;
  lastSeenNotification: string;
}

export interface me {
  me: me_me | null;
}
