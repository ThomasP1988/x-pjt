import { Blockchains } from "../../config/types";
import config from "../../config/config";
import Web3 from 'web3';

export const getConnectedBlockchain = (): Blockchains => {
    const provider = (window as any).ethereum;
    console.log("chainId", provider?.chainId);
    if (provider?.chainId) {
      for (let bc in config.blockchain) {
        const chainId: number = Web3.utils.hexToNumber(provider.chainId)
        if (config.blockchain[bc as Blockchains].chainId === chainId) {
          return bc as Blockchains;
        }
      }
    }
    return Blockchains.ETHEREUM
}