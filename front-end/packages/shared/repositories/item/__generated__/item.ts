/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL fragment: Item
// ====================================================

export interface Item_attributes {
  __typename: "ItemAttribute";
  display_type: string | null;
  trait_type: string | null;
  value: string | null;
  color: string | null;
}

export interface Item {
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
  attributes: (Item_attributes | null)[] | null;
}
