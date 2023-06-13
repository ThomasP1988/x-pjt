import { Blockchains } from "../../config/types";
import { NFTSelectionAction, NFTSelectionState, NFTSelectionActions, NFTSelectionActionBase, NFTSelectionActionSelect, NFTSelectionStateItem } from "./types";

export const initNFTSelectionState = (): NFTSelectionState => {

    const itemsByChain: NFTSelectionState["itemsByChain"] = {} as NFTSelectionState["itemsByChain"];
    const itemsByChainAndCollection: NFTSelectionState["itemsByChainAndCollection"] = {} as NFTSelectionState["itemsByChainAndCollection"];
    const itemsByChainLength: NFTSelectionState["itemsByChainLength"] = {} as NFTSelectionState["itemsByChainLength"];
    const chains: Blockchains[] = Object.values(Blockchains);

    for (let i = 0; i < chains.length; i++) {
        itemsByChain[chains[i]] = {};
        itemsByChainAndCollection[chains[i]] = {};
        itemsByChainLength[chains[i]] = 0;
    }

    return {
        chainSelected: Blockchains.ETHEREUM,
        itemsByChain,
        itemsByChainAndCollection,
        itemsByChainLength,
        items: {},
        itemsLength: 0,
    }
}

export const generateStateDepositKey = (collectionAddress: string, tokenId: string) => {
    return `${collectionAddress}#${tokenId}`
}

export const NFTSelectionReducer = (state: NFTSelectionState, action: NFTSelectionAction): NFTSelectionState => {
    const { type, chain } = action;
    switch (type) {
        case NFTSelectionActions.Toggle:
            const { collectionAddress, tokenId } = action as NFTSelectionActionBase<NFTSelectionActionSelect>;
            if (!state.items[generateStateDepositKey(collectionAddress, tokenId)]) {
                action.type = NFTSelectionActions.Select;
                return NFTSelectionReducer(state, action);
            } else {
                action.type = NFTSelectionActions.UnSelect;
                return NFTSelectionReducer(state, action);
            }
        case NFTSelectionActions.Select:
            {
                const { collectionAddress, tokenId } = action as NFTSelectionActionBase<NFTSelectionActionSelect>;
                const newState = Object.assign({}, state);
                const key: string = generateStateDepositKey(collectionAddress, tokenId);
                const item = {
                    collectionAddress,
                    tokenId
                }
                newState.items[key] = item;
                newState.itemsByChain[chain][key] = item;

                if (!newState.itemsByChainAndCollection[chain]?.[item.collectionAddress]) {
                    newState.itemsByChainAndCollection[chain][item.collectionAddress] = {}
                }
                newState.itemsByChainAndCollection[chain][item.collectionAddress][tokenId] = item;

                newState.itemsByChainLength[chain]++;
                newState.itemsLength++;

                return newState;
            }
        case NFTSelectionActions.UnSelect:
            {
                const { collectionAddress, tokenId } = action as NFTSelectionActionBase<NFTSelectionActionSelect>;
                const newState = Object.assign({}, state);
                const key: string = generateStateDepositKey(collectionAddress, tokenId);
                delete newState.itemsByChain[chain][key];
                delete newState.items[key];

                if (newState.itemsByChainAndCollection[chain][collectionAddress]) {
                    delete newState.itemsByChainAndCollection[chain][collectionAddress][tokenId];
                    if (!Object.keys(newState.itemsByChainAndCollection[chain][collectionAddress]).length) {
                        delete newState.itemsByChainAndCollection[chain][collectionAddress]
                    }
                }

                newState.itemsByChainLength[chain]--;
                newState.itemsLength--;
                return newState;
            }
        case NFTSelectionActions.Flush:
            {
                return initNFTSelectionState();
            }
        default:
            return state;
    }
};

export const GetCollectionsIds = (input: NFTSelectionState): string[] => {
    const collections: Record<string, boolean> = {};
    const items: NFTSelectionStateItem[] = Object.values(input.items) || [];

    for (let i = 0; i < items.length; i++) {
        if (!collections[items[i].collectionAddress]) {
            collections[items[i].collectionAddress] = true;
        }
    }

    return Object.keys(collections);
}