import { useState, useEffect } from "react";
import { FormControl, ListItemText, MenuItem, Stack, Select, SelectChangeEvent, CircularProgress, Grid } from "@mui/material"
import config from "../../config/config";
import { Blockchains } from "../../config/types";
import Web3 from 'web3';
import { getConnectedBlockchain } from "shared/lib/blockchain/getConnected";

type Props = {
  chainPlaceHolder?: Blockchains,
  onChainSelect?: (chain: Blockchains) => void,
  switchWallet?: boolean
}

export const SwitchNetwork = ({ chainPlaceHolder = Blockchains.ETHEREUM, onChainSelect, switchWallet = false }: Props) => {
  const [network, setNetwork] = useState<Blockchains>(chainPlaceHolder);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    if (setNetwork) {
      setNetwork(getConnectedBlockchain())
    }
  }, [setNetwork]);

  const handleChange = async (event: SelectChangeEvent<Blockchains>) => {
    const {
      target: { value },
    } = event;

    if (switchWallet) {
      const provider = (window as WindowChain).ethereum;
      if (provider) {
        setLoading(true);
        try {
          const result = await provider.request({
            method: 'wallet_switchEthereumChain',
            params: [
              {
                chainId: Web3.utils.toHex(config.blockchain[value as Blockchains].chainId)
              },
            ],
          });
          console.log("result", result);
          onChainSelect?.(value as Blockchains);
          setNetwork(value as Blockchains);
        } catch (error) {
          console.error(error);
        }
        setLoading(false);
      }
    } else {
      onChainSelect?.(value as Blockchains);
      setNetwork(value as Blockchains);
    }
  };

  return <>
    <Stack alignContent="center" alignItems="center" direction="row" >
      <Stack sx={{ width: 30 }} alignContent="center" alignItems="center">
        {
          loading && <CircularProgress size={25} />
        }
      </Stack>
      <FormControl sx={{ m: 1, width: 300 }}>
        <Select
          labelId="select-network"
          id="select-network"
          value={network}
          onChange={handleChange}
          size="small"
          sx={{ height: 36.5 }}
        >
          {Object.keys(config.blockchain).map((blockchain: string) => (
            <MenuItem key={blockchain} value={blockchain}>
              <Grid container alignContent="center" alignItems="center" direction="row" >
                <Grid item xs={1} container alignContent="center" justifyContent="center">
                  <img src={config.blockchain[blockchain as Blockchains].svg} alt={blockchain} height={20} />
                </Grid>
                <Grid item xs={10} container alignContent="center">
                  <ListItemText sx={{ marginLeft: 2 }} primary={config.blockchain[blockchain as Blockchains].name} />
                </Grid>
              </Grid>
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </Stack>
  </>
}