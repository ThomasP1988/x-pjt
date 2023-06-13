import { useState, useEffect } from 'react';
import { useMutation } from '@apollo/client';
import { LoadingButton } from '@mui/lab';
import { CREATE_TOKEN } from '../../repositories/token/mutations/createToken';
import { CONNECT_WALLET } from '../../repositories/wallet-blockchain/mutations/connect';
import { connectWalletVariables, connectWallet } from '../../repositories/wallet-blockchain/mutations/__generated__/connectWallet';
import { useWeb3React } from "@web3-react/core";
import { createToken } from '../../repositories/token/mutations/__generated__/createToken';
import { injected } from "./../../components/Web3Connector";

export function ConnectWallet() {
    const [loading, setLoading] = useState<boolean>(false);
    const [connecting, setConnecting] = useState<boolean>(false);
    const [createToken] = useMutation<createToken>(CREATE_TOKEN);
    const [triggerConnectWallet] = useMutation<connectWallet, connectWalletVariables>(CONNECT_WALLET);
    const { library, activate, error } = useWeb3React();

    // console.log("library", library);

    useEffect(() => {
        if (connecting) {
            (async () => {
                const newAccounts = await library.eth.getAccounts();

                const { data } = await createToken();
                const nonce = data?.createToken?.nonce;
                const tokenId = data?.createToken?.token;

                if (nonce && tokenId) {
                    const msg = nonce;

                    await library.currentProvider.sendAsync(
                        {
                            method: 'personal_sign',
                            params: [msg, newAccounts?.[0]],
                            from: newAccounts?.[0],
                        },
                        async (err: any, { result }: { result: string }) => {
                            if (result) {
                                await triggerConnectWallet({
                                    variables: {
                                        tokenId,
                                        signature: result,
                                    },
                                });
                            }
                        }
                    );
                }
                setLoading(false);
                setConnecting(false);
            })()
        }
    }, [library, connecting, setConnecting, createToken, triggerConnectWallet]);

    const connect = async () => {
        setLoading(true);
        try {
            await activate(injected);
            setConnecting(true);
        } catch (e) {
            console.log(e);
        }
    };

    return (<>
        <LoadingButton loading={loading} onClick={connect}>Connect Wallet</LoadingButton>
        {
            error && <>{error}</>
        }
    </>
    )
}
