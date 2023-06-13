/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { CollectionStatus } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL query operation: listCollections
// ====================================================

export interface listCollections_listCollections_collections {
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

export interface listCollections_listCollections {
  __typename: "ListCollection";
  collections: listCollections_listCollections_collections[];
  next: string | null;
}

export interface listCollections {
  listCollections: listCollections_listCollections | null;
}

export interface listCollectionsVariables {
  limit?: number | null;
  from?: string | null;
  status?: CollectionStatus | null;
}
