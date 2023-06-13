/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: inviteUser
// ====================================================

export interface inviteUser_inviteUser {
  __typename: "InviteResponse";
  success: boolean | null;
}

export interface inviteUser {
  inviteUser: inviteUser_inviteUser | null;
}

export interface inviteUserVariables {
  email: string;
}
