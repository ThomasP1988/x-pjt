import { InjectedConnector } from '@web3-react/injected-connector';
import config from '../../config/config';
import { Blockchains } from '../../config/types';

console.log("supported chains", Object.keys(config.blockchain).map((blockchain: string) => config.blockchain[blockchain as Blockchains].chainId))

export const injected = new InjectedConnector({
  supportedChainIds: Object.keys(config.blockchain).map((blockchain: string) => config.blockchain[blockchain as Blockchains].chainId),
});