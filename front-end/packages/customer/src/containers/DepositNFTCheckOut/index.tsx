import { useState, useEffect, Fragment } from "react";
import { Stack, Paper, Typography, CircularProgress, Alert, Grid, Link, Button } from "@mui/material";
import { useDepositNFTContext } from "../../context/deposit";
import { Blockchains } from "shared/config/types";
import { LIST_COLLECTIONS_AND_ITEMS } from "shared/repositories/collection/queries/list-and-list-items";
import { listCollectionsAndItems, listCollectionsAndItemsVariables } from "shared/repositories/collection/queries/__generated__/listCollectionsAndItems";
import { NFTSelectionActions, NFTSelectionStateItem } from "shared/reducers/nft-selection/types";
import config from "shared/config/config";
import { useQuery } from "@apollo/client";
import { Collection } from "shared/repositories/collection/__generated__/Collection";
import { Item } from "shared/repositories/item/__generated__/item";
import { GetCollectionsIds } from "shared/reducers/nft-selection/reducer";
import { Colors } from "../../styles/theme";
import { useWeb3React } from "@web3-react/core";
import ERC1155 from 'shared/lib/abi/ERC1155.json';
import Web3 from "web3";
import { AbiItem } from 'web3-utils';



export const DepositNFTCheckOut = () => {
    const { stateDeposit } = useDepositNFTContext();

    return <>
        {
            stateDeposit?.itemsLength ? <DisplayCheckoutDetails /> : <>No NFTs Selected, please go back to deposit.</>
        }
    </>
}

export const DisplayCheckoutDetails = () => {
    const { stateDeposit, dispatchDeposit } = useDepositNFTContext();
    console.log("stateDeposit", stateDeposit);
    const [collectionsByIds, setCollectionsById] = useState<Record<string, Collection>>({});
    const [nftItemsByIds, setNftItemsByIds] = useState<Record<string, Item>>({});
    const { active, account, library, activate, connector, deactivate, chainId } = useWeb3React()

    // useEffect(() => {

    // }, [setOrderedItemsByCollections, stateDeposit.allItems])

    const { loading, error } = useQuery<listCollectionsAndItems, listCollectionsAndItemsVariables>(LIST_COLLECTIONS_AND_ITEMS, {
        variables: {
            ids: GetCollectionsIds(stateDeposit),
            itemKeys: Object.values(stateDeposit.items)
        },
        onCompleted: (data: listCollectionsAndItems): void => {
            console.log("data", data)
            if (data.listCollectionsByIds) {
                setCollectionsById(Object.assign({}, ...data.listCollectionsByIds.map((item: Collection) => ({ [item.id]: item }))));
            }
            if (data.listItemsByIds) {
                setNftItemsByIds(Object.assign({}, ...data.listItemsByIds.map((item: Item) => ({ [item.id]: item }))))
            }
        },
    });

    const deposit = (chain: Blockchains) => {
        // check blockchain id and switch network ?

        // retrieve tokenIds and batch send them to a specific wallet
        // const contract = new (library as Web3).eth.Contract(ERC1155 as AbiItem[], selectedCollection.id);


    }

    return <>
        {
            loading && <Stack direction="row" justifyContent="center"><CircularProgress /></Stack>
        }
        {
            error && <Alert severity="error">{error.message}</Alert>
        }
        {
            Object.keys(stateDeposit.itemsByChainAndCollection)
                .filter((chain: string) => Boolean(stateDeposit.itemsByChainLength[chain as Blockchains]))
                .map((chain: string) => {
                    return <Paper key={chain} sx={{ marginBottom: 2, padding: 1 }}>
                        <Stack sx={{ padding: 1 }}>
                            <Typography variant="h6">
                                <img src={config.blockchain[chain as Blockchains].svg} alt={chain}
                                    style={{
                                        verticalAlign: "middle",
                                        marginRight: 10,
                                        maxHeight: 20,
                                        maxWidth: 20
                                    }}
                                />
                                {config.blockchain[chain as Blockchains].name}
                            </Typography>
                        </Stack>
                        {
                            Object.keys(stateDeposit.itemsByChainAndCollection[chain as Blockchains]).map((collectionId: string) => {
                                return <Fragment key={collectionId} ><Typography variant="body1" sx={{ marginLeft: 2 }}>
                                    {
                                        collectionsByIds[collectionId]?.name || "no collection name"
                                    }

                                </Typography>
                                    <Grid container>
                                        {
                                            Object.values(stateDeposit.itemsByChainAndCollection[chain as Blockchains][collectionId]).map((value: NFTSelectionStateItem) => {
                                                return <Fragment key={value.tokenId}>
                                                    <Grid item xs={6}>
                                                        <Typography variant="body2" sx={{ marginLeft: 4 }}>
                                                            {
                                                                `${nftItemsByIds[value.tokenId]?.id}. ${nftItemsByIds[value.tokenId]?.name}`
                                                            }
                                                        </Typography>
                                                    </Grid>
                                                    <Grid item xs={6} container justifyContent="flex-end">
                                                        <Link variant="body2" color={Colors.grey} sx={{ cursor: "pointer" }}
                                                            onClick={() => dispatchDeposit({
                                                                type: NFTSelectionActions.UnSelect,
                                                                collectionAddress: collectionId,
                                                                tokenId: value.tokenId,
                                                                chain: chain as Blockchains
                                                            })}
                                                        >
                                                            unselect
                                                        </Link>
                                                    </Grid>
                                                </Fragment>
                                            })
                                        }
                                    </Grid>
                                </Fragment>
                            })
                        }
                        <Stack direction="row" justifyContent="flex-end" sx={{ paddingTop: 2 }}>
                            <Button variant="contained" disableElevation size="medium" color="success" onClick={() => deposit(chain as Blockchains)}>deposit</Button>
                        </Stack>
                    </Paper>
                })
        }
    </>
}