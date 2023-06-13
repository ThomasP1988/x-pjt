import { useState } from 'react';
import { LoadingButton } from '@mui/lab';
import config from '../../config/config';
import { Blockchains } from '../../config/types';
import Web3 from 'web3';

type Props = {
    chain: Blockchains
}

export const AddNetwork = ({ chain }: Props) => {
    const [loading, setLoading] = useState<boolean>(false);

    const onClick = async () => {
        console.log("add network")
        const provider = (window as WindowChain).ethereum;
        console.log("provider", provider);
        if (provider) {
            setLoading(true)
            try {
               const result = await provider.request({
                    method: 'wallet_addEthereumChain',
                    params: [
                        {
                            chainId: Web3.utils.toHex(config.blockchain[chain].chainId),
                            chainName: config.blockchain[chain].name,
                            nativeCurrency: {
                                name: config.blockchain[chain].currency?.name,
                                symbol: config.blockchain[chain].currency?.symbol,
                                decimals: config.blockchain[chain].currency?.decimals,
                            },
                            rpcUrls: [config.blockchain[chain].network],
                            blockExplorerUrls: [config.blockchain[chain].explorer],
                        },
                    ],
                });
                console.log("result", result);
            } catch (error) {
                console.error(error);
            }
            setLoading(false)
        }
    }

    return <LoadingButton onClick={onClick} loading={loading}>Add</LoadingButton>
      
}