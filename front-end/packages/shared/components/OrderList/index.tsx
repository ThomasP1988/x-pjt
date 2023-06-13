import React, { useEffect, useState } from 'react';
import { RequestOrderList } from '../../repositories/orderbook-grpc';
import { PrintStatus, ColorStatus } from '../../lib/orders';
import { OrderListResult, SideOrder, TypeOrder, Order } from '../../grpc/orderbook';
import moment from 'moment';
import CircularProgress from '@mui/material/CircularProgress';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Chip from '@mui/material/Chip';

import Stack from '@mui/material/Stack';
import LoadingButton from '@mui/lab/LoadingButton';
import RefreshIcon from '@mui/icons-material/Refresh';
import ArrowDownwardIcon from '@mui/icons-material/ArrowDownward';
import { OrderStatus } from '../../constants';
import { Typography } from '@mui/material';
import { ErrorPage } from '../ErrorPage';
import BubbleChartIcon from '@mui/icons-material/BubbleChart';
import { useColors } from "../../lib/hooks/colors";
import { CancelOrderButton } from './cancel';

type Props = {
    symbol?: string,
    isOpen?: boolean,
    hideCancel?: boolean,
    blue: string
}

const limit: number = 20;

export const OrderList = ({ symbol, isOpen, hideCancel, blue = "blue" }: Props) => {
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<boolean>(false);
    const [orders, setOrders] = useState<Order[]>([]);
    const [next, setNext] = useState<string | undefined>();
    const colors = useColors();

    const loadMore = async (refresh?: boolean) => {
        let result: OrderListResult | undefined;
        setLoading(true);
        setError(false);
        try {
            result = await RequestOrderList({
                symbol,
                limit,
                from: next,
                isOpen
            });
        } catch (e) {
            setLoading(false);
            setError(true);
            console.log(e);
            return;
        }

        if (!refresh) {
            setOrders([...orders, ...result.orders]);
        } else {
            setOrders(result.orders);
        }

        setNext(result.next);
        setLoading(false);
    }

    const updateOrder = (index: number, order: Order) => {
        const newOrders: Order[] = orders.slice();
        newOrders[index] = order;
        setOrders(newOrders);
    }

    useEffect(() => {
        RequestOrderList({
            symbol,
            limit,
            isOpen
        }).then((result: OrderListResult) => {
            console.log(result.orders)
            setOrders(result.orders);
            setNext(result.next);
            setLoading(false);
        }).catch((e) => {
            setError(true);
            setLoading(false);
        })
    }, [setOrders, setNext, setLoading, setError, isOpen, symbol]);

    return (
        <>
            {
                error ? <ErrorPage bottom={
                    <LoadingButton
                        loading={loading}
                        onClick={() => loadMore(true)}
                    >
                        Try Again
                    </LoadingButton>
                } /> :
                    <>
                        <Stack direction="row" spacing={2} justifyContent="flex-end" sx={{ paddingTop: 1, paddingBottom: 1 }}>
                            <LoadingButton
                                loading={loading}
                                loadingPosition="start"
                                startIcon={<RefreshIcon />}
                                variant="outlined"
                                onClick={() => loadMore(true)}
                            >
                                Refresh
                            </LoadingButton>
                        </Stack>
                        {
                            (!loading && !orders.length) ? <Stack justifyContent="center" alignContent="center" alignItems="center" sx={{ paddingTop: 10 }}>
                                <BubbleChartIcon sx={{ color: colors.blue, fontSize: 72 }} />
                                <Typography variant="h3">Nothing to see here yet.</Typography>
                                <Typography variant="body1">Come on, it's time for you to buy or sell some NFTs.</Typography>
                            </Stack> : <TableContainer component={Paper}>
                                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                                    <TableHead>
                                        <TableRow>
                                            <TableCell>Placed</TableCell>
                                            <TableCell>Status</TableCell>
                                            <TableCell align="center">Symbol</TableCell>
                                            <TableCell align="center">Side</TableCell>
                                            <TableCell align="center">Type</TableCell>
                                            <TableCell align="right">Quantity</TableCell>
                                            <TableCell align="right">Filled</TableCell>
                                            <TableCell align="right">Left</TableCell>
                                            <TableCell align="right">Price</TableCell>
                                            {
                                                !hideCancel && <TableCell align="right"></TableCell>
                                            }
                                        </TableRow>
                                    </TableHead>
                                    <TableBody>
                                        {orders.map((row: Order, index: number) => (
                                            <TableRow
                                                key={row.id}
                                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                            >
                                                <TableCell component="th" scope="row">
                                                    {moment(row.createdAt.toString()).format('MMMM Do YYYY, h:mm:ss a')}
                                                </TableCell>
                                                <TableCell>
                                                    <Typography sx={{ color: ColorStatus(row.status as OrderStatus, colors) }} variant="button">
                                                        {PrintStatus(row.status as OrderStatus)}
                                                    </Typography>
                                                </TableCell>
                                                <TableCell align="center">{row.symbol}</TableCell>
                                                <TableCell align="center">
                                                    {row.side === SideOrder.BUY && <Chip label="Buy" color="success" />}
                                                    {row.side === SideOrder.SELL && <Chip label="Sell" color="error" />}
                                                </TableCell>
                                                <TableCell align="center">
                                                    {row.type === TypeOrder.LIMIT && <Chip label="Limit" color="primary" variant="outlined" />}
                                                    {row.type === TypeOrder.MARKET && <Chip label="Market" color="secondary" variant="outlined" />}
                                                </TableCell>
                                                <TableCell align="right">
                                                    {
                                                        row.originalQuantity
                                                    }
                                                </TableCell>
                                                <TableCell align="right">
                                                    {
                                                        row.filledQuantity
                                                    }
                                                </TableCell>
                                                <TableCell align="right">
                                                    {
                                                        row.quantity
                                                    }
                                                </TableCell>
                                                <TableCell align="right">{row.price}</TableCell>
                                                {
                                                    !hideCancel && <TableCell align="right">
                                                        {
                                                            row.status !== OrderStatus.Cancelled && <CancelOrderButton updatedOrder={(order: Order) => updateOrder(index, order)} orderId={row.id} />
                                                        }
                                                    </TableCell>
                                                }
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </TableContainer>
                        }
                    </>
            }
            {
                Boolean(next) && <Stack direction="row" spacing={2} justifyContent="center" sx={{ paddingTop: 2, paddingBottom: 2 }}>
                    <LoadingButton
                        loading={loading}
                        loadingPosition="start"
                        startIcon={<ArrowDownwardIcon />}
                        variant="outlined"
                        onClick={() => loadMore()}
                    >
                        Load More
                    </LoadingButton>
                </Stack>
            }
            {
                loading && !orders?.length && <Stack justifyContent="center" alignContent="center" alignItems="center" sx={{ paddingTop: 10 }}>
                    <CircularProgress />
                </Stack>
            }
        </>
    )
}