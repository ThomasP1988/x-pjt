import React, { Dispatch } from 'react';
import { NFTSelectionState, NFTSelectionAction } from "shared/reducers/nft-selection/types";

interface IContextProps {
    stateDeposit: NFTSelectionState;
    dispatchDeposit: Dispatch<NFTSelectionAction>
}

export const DepositNFTContext = React.createContext({} as IContextProps);

export function useDepositNFTContext() {
    return React.useContext(DepositNFTContext);
}

