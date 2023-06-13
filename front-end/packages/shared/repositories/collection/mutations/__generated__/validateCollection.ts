/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { CollectionStatus } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL mutation operation: validateCollection
// ====================================================

export interface validateCollection_validateCollection {
  __typename: "Collection";
  id: string;
  symbol: string;
  description: string | null;
  name: string;
  supply: string;
  chain: string;
  imagePath: string | null;
  thumbnailPath: string | null;
  firstItemId: number | null;
  status: CollectionStatus;
  openseaSlug: string | null;
}

export interface validateCollection {
  validateCollection: validateCollection_validateCollection | null;
}

export interface validateCollectionVariables {
  address: string;
  description?: string | null;
  name?: string | null;
  image?: string | null;
  symbol?: string | null;
  openseaSlug?: string | null;
  status?: CollectionStatus | null;
}
