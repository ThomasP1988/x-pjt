import React from 'react';
import { Typography, Grid } from '@mui/material';
import { useColors } from "shared/lib/hooks/colors";
import { SignUpWithoutPassword } from 'shared/components/auth/SignUpWithoutPassword';
import { ConnectWallet } from 'shared/components/ConnectWallet';
import { AddNetwork } from "shared/components/AddNetwork";
import { Blockchains } from 'shared/config/types';

export const Home = () => {

    const colors = useColors();

    return (<Grid container direction="row" sx={{marginTop: 5}}>
        <Grid item xs={12} md={6} container justifyContent="center" alignItems="center" alignContent="center">
            <Grid item xs={10}>
                <Typography variant="h2" gutterBottom component="div">
                    <Typography variant="h2" component="span" color={colors.green}> <strong>less</strong></Typography> fees
                </Typography>
            </Grid>
            <Grid item xs={10}>
                <Typography variant="h2" gutterBottom component="div">
                    <Typography variant="h2" component="span" color={colors.green}> <strong>less</strong></Typography> wait
                </Typography>
            </Grid>
            <Grid item xs={10}>
                <Typography variant="h2" gutterBottom component="div">
                    <Typography variant="h2" component="span" color={colors.green}> <strong>more</strong></Typography> exchange
                </Typography>
            </Grid>
            <Grid item xs={10}>
                Buy &amp; trade NFT as much as you want on any collections.
            </Grid>
        </Grid>
        <Grid item xs={12} md={6}>
            <SignUpWithoutPassword />
            <ConnectWallet />
            <AddNetwork chain={Blockchains.POLYGON} />
        </Grid>
    </Grid>)
}