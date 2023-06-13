/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { CollectionStatus } from "./../../../__generated__/globalTypes";

// ====================================================
// GraphQL fragment: Collection
// ====================================================

export interface Collection {
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
