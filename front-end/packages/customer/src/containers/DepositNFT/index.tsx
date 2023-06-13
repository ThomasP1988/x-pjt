import { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import Web3 from "web3";
import { AbiItem } from 'web3-utils';
import { useWeb3React } from "@web3-react/core";
import { injected } from "shared/components/Web3Connector";
import { Container, Typography, Stack, Button, Grid, Tooltip, Badge, AppBar } from '@mui/material';
import { SelectCollection } from 'shared/components/SelectCollection';
import NFT_ABI from 'shared/lib/abi/ERC721.json';
import { SwitchNetwork } from 'shared/components/SwitchNetwork';
import { Blockchains } from 'shared/config/types';
import { LoadingButton } from '@mui/lab';
import CableIcon from '@mui/icons-material/Cable';
import { AddCollectionDialog } from 'shared/components/AddCollectionDialog';
import { Collection } from 'shared/repositories/collection/__generated__/Collection';
import { Contract } from 'web3-eth-contract';
import { GetBatchOfToken } from "./helper";
import AddIcon from '@mui/icons-material/Add';
import { ListNFTItem } from "shared/components/ListNFTItem";
import ShoppingBasketIcon from '@mui/icons-material/ShoppingBasket';
import { useDepositNFTContext } from '../../context/deposit';
import { Colors } from '../../styles/theme';
import { NFTSelectionActions } from 'shared/reducers/nft-selection/types';
import { generateStateDepositKey } from 'shared/reducers/nft-selection/reducer';
import { getConnectedBlockchain } from 'shared/lib/blockchain/getConnected';

const batchSize: number = 10;
export const DepositNFT = () => {
    const [totalNFT, setTotalNFT] = useState<string | null>(null);
    const navigate = useNavigate()
    const [NFTContract, setNFTContract] = useState<Contract | null>(null);
    const [connectedChain, setConnectedChain] = useState<Blockchains | null>(null);
    const [selectedCollection, setSelectedCollection] = useState<Collection | null>(null);
    const [addCollectionDialogOpen, setAddCollectionDialogOpen] = useState<boolean>(false);
    const [page, setPage] = useState<number>(0);
    const [tokenIds, setTokenIds] = useState<number[]>([]);
    const [loadingTokens, setLoadingTokens] = useState<boolean>(false);
    const { stateDeposit, dispatchDeposit } = useDepositNFTContext();

    const { active, account, library, activate, connector, deactivate, chainId } = useWeb3React()
    console.log("chainId", chainId);
    console.log("connector", connector);
    const fetchNFTsByCollection = useCallback(
        async () => {
            if (selectedCollection && account) {
                setLoadingTokens(true);
                const contract = new (library as Web3).eth.Contract(NFT_ABI as AbiItem[], selectedCollection.id);
                try {

                    const balanceOf = await contract.methods.balanceOf(account).call();
                    console.log("selectedCollection", selectedCollection);
                    console.log("balanceOf", balanceOf);

                    const firstTokenIds: number[] = await GetBatchOfToken({
                        account,
                        batchSize,
                        contract,
                        total: balanceOf,
                        page: 0
                    });

                    console.log("firstTokenIds", firstTokenIds);

                    setTokenIds(firstTokenIds)
                    setTotalNFT(balanceOf);
                    setNFTContract(contract);
                } catch (e) {
                    console.log(e);
                }
                setLoadingTokens(false);
            }
        },
        [selectedCollection, account, library, setNFTContract],
    );

    useEffect(() => {
        setConnectedChain(getConnectedBlockchain());
    }, [setConnectedChain]);

    useEffect(() => {
        fetchNFTsByCollection();
    }, [selectedCollection, connectedChain, fetchNFTsByCollection]);

    async function connect() {
        console.log("account", account);
        console.log("library", library);
        console.log("activate", activate);
        console.log("connector", connector);

        try {
            await activate(injected);
        } catch (e) {
            console.log(e)
        }
    }

    async function disconnect() {
        try {
            deactivate()
        } catch (e) {
            console.log(e)
        }
    }

    async function onCollectionChange(collection: Collection) {
        setSelectedCollection(collection);
        setPage(0);
    }

    async function onChainChange(chain: Blockchains) {
        setConnectedChain(chain);
    }

    async function onAddedCollection(collection: Collection) {
        setAddCollectionDialogOpen(false)
        setSelectedCollection(collection);
    }

    async function loadMore(): Promise<void> {
        const newPage = page + 1;
        setPage(newPage);
        setLoadingTokens(true);
        if (account && NFTContract && totalNFT) {
            try {
                const moreTokenIds: number[] = await GetBatchOfToken({
                    account,
                    batchSize,
                    contract: NFTContract,
                    total: Number(totalNFT),
                    page: newPage
                });

                setTokenIds([
                    ...tokenIds,
                    ...moreTokenIds
                ])

            } catch (e) {
                console.log(e);
            }
        }
        setLoadingTokens(false);
    }

    return (
        <><Container style={{ marginBottom: 40 }}>
            {
                !active || !connectedChain ? <>
                    <Stack justifyContent="center" alignContent="center" alignItems="center" sx={{ paddingTop: 1 }}>
                        <Typography variant="h4" gutterBottom>Deposit NFT</Typography>
                        <Typography variant="h5" gutterBottom>In order to deposit NFT to your account, connect your wallet</Typography>
                        <LoadingButton variant="contained" color="secondary" onClick={connect} startIcon={<CableIcon />}>Connect your wallet</LoadingButton>
                    </Stack>
                </> :
                    <>
                        <Stack justifyContent="right" alignContent="center" alignItems="center" direction="row">
                            <Tooltip title="Connected wallet" placement="left">
                                <Typography variant="subtitle2" gutterBottom>{account}</Typography>
                            </Tooltip>
                        </Stack>
                        <Grid container alignContent="stretch" alignItems="center" direction="row">
                            <Grid item md={6} >
                                <SelectCollection onSelect={onCollectionChange} sx={{ width: 300 }} />
                                <Button style={{ marginLeft: 5 }} startIcon={<AddIcon />} onClick={() => setAddCollectionDialogOpen(true)}> Or add  a collection</Button>
                            </Grid>
                            <Grid item md={6} container justifyContent="right" alignContent="center" alignItems="center" direction="row">
                                <SwitchNetwork chainPlaceHolder={Blockchains.ETHEREUM} switchWallet={true} onChainSelect={onChainChange} />
                                <Button onClick={disconnect} variant="outlined">Disconnect</Button>
                            </Grid>
                        </Grid>
                        <Stack justifyContent="right" alignContent="center" alignItems="center" direction="row">
                            <Typography variant="subtitle2" gutterBottom>{totalNFT ? `showing 1 - ${tokenIds.length} of ${totalNFT}` : ""}</Typography>
                        </Stack>
                        <Grid container alignContent="stretch" alignItems="center" direction="row" columnSpacing={2}>
                            {
                                selectedCollection && tokenIds.map((tokenId: number) => {
                                    return <Grid item md={4} key={tokenId}>
                                        <ListNFTItem collectionAddress={selectedCollection.id} tokenId={tokenId} actions={
                                            <>
                                                <Button size="small" onClick={() => dispatchDeposit({ type: NFTSelectionActions.Toggle, collectionAddress: selectedCollection.id, tokenId: String(tokenId), chain: connectedChain })}>
                                                    {
                                                        stateDeposit.itemsByChain[connectedChain][generateStateDepositKey(selectedCollection.id, String(tokenId))] ? "unselect" : "select"
                                                    }
                                                </Button>
                                            </>
                                        } />
                                    </Grid>
                                })
                            }
                        </Grid>
                        {
                            tokenIds?.length !== Number(totalNFT) && <Stack justifyContent="center" alignContent="center" alignItems="center" direction="row">
                                <LoadingButton startIcon={<AddIcon />} loading={loadingTokens} onClick={loadMore}>Load more</LoadingButton>
                            </Stack>
                        }

                    </>
            }
            <AddCollectionDialog
                open={addCollectionDialogOpen}
                close={() => setAddCollectionDialogOpen(false)}
                onSuccess={onAddedCollection}
            />
        </Container>
            {
                connectedChain && <AppBar position="fixed" color="inherit" style={{ bottom: 0, top: "inherit" }}>
                    <Container>
                        <Stack justifyContent="right" alignContent="center" alignItems="center" direction="row">
                            <Tooltip title={`${stateDeposit.itemsByChainLength[connectedChain]} NFTs selected`} placement="left">
                                <Button
                                    onClick={() => navigate("/deposit/nft/checkout")}
                                    // style={{ marginRight: 10 }}
                                    endIcon={
                                        <Badge badgeContent={stateDeposit.itemsByChainLength[connectedChain]} color="error" style={{ marginRight: 5 }}>
                                            <ShoppingBasketIcon style={{ color: stateDeposit.itemsByChainLength[connectedChain] ? Colors.black : Colors.grey }} />
                                        </Badge>
                                    }
                                    disabled={!Boolean(stateDeposit.itemsByChainLength[connectedChain])}
                                >
                                    Review &amp; Deposit
                                </Button>
                            </Tooltip>
                        </Stack>
                    </Container>
                </AppBar>
            }
        </>
    )
}