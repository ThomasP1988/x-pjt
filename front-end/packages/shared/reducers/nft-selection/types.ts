import { Blockchains } from "../../config/types";

export enum NFTSelectionActions {
  Toggle,
  Select,
  UnSelect,
  Flush
}

export type NFTSelectionState = {
  chainSelected: Blockchains;
  items: Record<string, NFTSelectionStateItem>;
  itemsByChain: Record<Blockchains, Record<string, NFTSelectionStateItem>>;
  itemsByChainAndCollection: Record<Blockchains, Record<string, Record<string, NFTSelectionStateItem>>>;
  itemsByChainLength: Record<Blockchains, number>;
  itemsLength: number;
};

export type NFTSelectionStateItem = {
  tokenId: string;
  collectionAddress: string;
};

export type NFTSelectionActionBase<T = {}> = {
  type: NFTSelectionActions;
  chain: Blockchains;
} & T;

export type NFTSelectionActionSelect = {
  tokenId: string;
  collectionAddress: string;
};

export type NFTSelectionAction = NFTSelectionActionBase<NFTSelectionActionSelect> | NFTSelectionActionBase;