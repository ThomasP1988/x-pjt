/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { ItemKeys, CollectionStatus } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL query operation: listCollectionsAndItems
// ====================================================

export interface listCollectionsAndItems_listCollectionsByIds {
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

export interface listCollectionsAndItems_listItemsByIds_attributes {
  __typename: "ItemAttribute";
  display_type: string | null;
  trait_type: string | null;
  value: string | null;
  color: string | null;
}

export interface listCollectionsAndItems_listItemsByIds {
  __typename: "Item";
  id: string;
  tokenURI: string | null;
  isFetching: boolean | null;
  collectionAddress: string | null;
  image_data: string | null;
  name: string | null;
  image: string | null;
  imagePath: string | null;
  thumbnailPath: string | null;
  description: string | null;
  external_url: string | null;
  background_color: string | null;
  animation_url: string | null;
  youtube_url: string | null;
  attributes: (listCollectionsAndItems_listItemsByIds_attributes | null)[] | null;
}

export interface listCollectionsAndItems {
  listCollectionsByIds: listCollectionsAndItems_listCollectionsByIds[];
  listItemsByIds: listCollectionsAndItems_listItemsByIds[];
}

export interface listCollectionsAndItemsVariables {
  ids: string[];
  itemKeys: ItemKeys[];
}
