import React from 'react';
import { Stack, Typography } from '@mui/material';
import FmdBadIcon from '@mui/icons-material/FmdBad';
import { useColors } from "../../lib/hooks/colors";

type Props = {
    bottom?: React.ReactNode
    text?: string,
    hideTryAgain?: boolean
}

export const ErrorPage = ({ bottom, text, hideTryAgain }: Props) => {
    const colors = useColors();
    return (
        <Stack justifyContent="center" alignContent="center" alignItems="center" sx={{ paddingTop: 10 }}>
            <FmdBadIcon sx={{ color: colors.orange, fontSize: 72 }} />
            <Typography variant="h3">{text || "Oops! An error happened."}</Typography>
            <Typography variant="body1">{!hideTryAgain && "Our teams have been warned, please try again later."}</Typography>
            {bottom}
        </Stack>
    )
}