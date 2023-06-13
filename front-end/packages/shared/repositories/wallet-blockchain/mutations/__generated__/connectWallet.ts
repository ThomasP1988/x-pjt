/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: connectWallet
// ====================================================

export interface connectWallet_connectWallet {
  __typename: "ConnectWalletResponse";
  success: boolean | null;
}

export interface connectWallet {
  connectWallet: connectWallet_connectWallet | null;
}

export interface connectWalletVariables {
  signature?: string | null;
  tokenId?: string | null;
}
