import React from 'react';
import { BidAsk } from "../../constants";
import { PriceLevel } from './../../grpc/orderbook';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { Typography, TypographyProps } from '@mui/material';
import { useColors } from "../../lib/hooks/colors";

type Props = {
    bidOrAsk: BidAsk
    prices: PriceLevel[]
    hideHeader?: boolean
}

export const Prices = ({ bidOrAsk, prices, hideHeader }: Props) => {
    const colors = useColors();

    const color: TypographyProps["color"] = bidOrAsk === BidAsk.Bid ? colors.green : colors.red;

    const tableCellStyle: React.CSSProperties = {
        paddingTop: 0,
        paddingBottom: 0,
    }

    return (
        <Table size="small">
            {!hideHeader && <TableHead>
                <TableRow>
                    <TableCell>Price</TableCell>
                    <TableCell align="right">Quantity</TableCell>
                </TableRow>
            </TableHead>}
            <TableBody>
                {
                    prices?.map((item: PriceLevel, index: number) => {
                        return <TableRow hover key={index}>
                            <TableCell style={tableCellStyle}>
                                <Typography color={color} variant="caption">
                                    {item.price}
                                </Typography>
                            </TableCell>
                            <TableCell style={tableCellStyle} align="right">
                                <Typography color={colors.blue} variant="caption">
                                    {item.price}
                                </Typography>
                            </TableCell>
                        </TableRow>
                    })
                }
            </TableBody>
        </Table>
    )

}