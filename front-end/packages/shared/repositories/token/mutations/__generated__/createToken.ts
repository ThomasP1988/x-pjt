/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: createToken
// ====================================================

export interface createToken_createToken {
  __typename: "CreateNonceToken";
  token: string | null;
  nonce: string | null;
}

export interface createToken {
  createToken: createToken_createToken | null;
}
