/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { CollectionStatus } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL query operation: searchCollections
// ====================================================

export interface searchCollections_searchCollections_results {
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

export interface searchCollections_searchCollections {
  __typename: "SearchCollection";
  results: searchCollections_searchCollections_results[];
  total: number;
}

export interface searchCollections {
  searchCollections: searchCollections_searchCollections | null;
}

export interface searchCollectionsVariables {
  text: string;
  from?: string | null;
  limit?: number | null;
}
