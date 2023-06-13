import React, { useState } from 'react';
import Grid from '@mui/material/Grid';

import Tooltip from '@mui/material/Tooltip';
import Button from '@mui/material/Button';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';

import { InputLabelInside } from '../../Theme/Input';
import { OrderArgs, SideOrder, TypeOrder, Order } from '../../../grpc/orderbook';
import { SellBuy, LimitMarket } from "../../../constants";
import { NewOrderRequest } from '../../../repositories/orderbook-grpc/newOrder';
import { ProcessingOrderDialog } from "./processDialog";

type Props = {
    symbol: string
}

export const NewOrder = ({ symbol }: Props) => {

    const [buyOrSell, setBuyOrSell] = useState<SellBuy>(SellBuy.Buy);
    const [dialogOpen, setDialogOpen] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);
    const [orderResult, setOrderResult] = useState<Order | null>(null);
    const [quantity, setQuantity] = useState<string>("");
    const [price, setPrice] = useState<string>("");
    const [limitOrMarket, setLimitOrMarket] = useState<LimitMarket>(LimitMarket.Limit);


    const checkDecimal = (input: string, size: number): boolean => {
        if (input.split(".")?.[1]) {
            return input.split(".")?.[1].length <= size;
        }
        return true;
    }

    const handleChangePrice = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>): void => {
        const newPrice: string = e?.target?.value;
        if (checkDecimal(newPrice, 8)) {
            setPrice(newPrice);
        }
    }

    const handleChangeQuantity = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>): void => {
        const newQuantity: string = e?.target?.value;
        if (checkDecimal(newQuantity, 8)) {
            setQuantity(newQuantity);
        }
    }

    const triggerOrder = async () => {
        console.log("buyOrSell", buyOrSell);
        console.log("limitOrMarket", limitOrMarket);
        setDialogOpen(true);
        setLoading(true);
        setOrderResult(null);
        setError(null);
        const order: OrderArgs = {
            symbol,
            side: buyOrSell as unknown as SideOrder,
            quantity,
            type: limitOrMarket as unknown as TypeOrder,
            price: "0"
        };

        if (limitOrMarket === LimitMarket.Limit) {
            order.price = price;
        }

        console.log("order sent", order)

        try {
            const result: Order = await NewOrderRequest(order);
            setOrderResult(result);
        } catch (e) {
            console.log(e);
            setError((e as Error)?.message || "")
        }
        setLoading(false);
    }

    return (<>
        <Grid container alignContent="center" spacing={{ paddingTop: 8 }} >
            <Grid item xs={12} container justifyContent="center">
                <ToggleButtonGroup sx={{ minWidth: "100%", maxHeight: 40 }}
                    exclusive
                    value={buyOrSell}
                    onChange={(_: React.MouseEvent<HTMLElement, MouseEvent>,
                        value: string | null) => {
                        console.log("value", value)
                        setBuyOrSell(Number(value) as SellBuy)
                    }}
                >
                    <ToggleButton color="success" style={{ minWidth: "50%" }} value={SellBuy.Buy}>Buy</ToggleButton>
                    <ToggleButton color="error" style={{ minWidth: "50%" }} value={SellBuy.Sell}>Sell</ToggleButton>
                </ToggleButtonGroup>
            </Grid>
            <Grid item xs={12}>
                <Tooltip title="A limit order is an order to buy or sell a stock with a restriction on the maximum price to be paid or the minimum price to be received (the “limit price”). If the order is filled, it will only be at the specified limit price or better. However, there is no assurance of execution." placement="top">
                    <Button variant="text" color={limitOrMarket === LimitMarket.Limit ? "info" : "inherit"}
                        onClick={() => setLimitOrMarket(LimitMarket.Limit)}
                    >Limit</Button>
                </Tooltip>
                <Tooltip title="A market order is an order to buy or sell a security immediately. This type of order guarantees that the order will be executed, but does not guarantee the execution price. A market order generally will execute at or near the current bid (for a sell order) or ask (for a buy order) price." placement="top">
                    <Button variant="text" color={limitOrMarket === LimitMarket.Market ? "info" : "inherit"}
                        onClick={() => setLimitOrMarket(LimitMarket.Market)}
                    >Market</Button>
                </Tooltip>
            </Grid>
            <Grid item xs={12}>
                <InputLabelInside label="Price" style={{ marginBottom: 10 }}
                    type={limitOrMarket === LimitMarket.Market ? "text" : "number"}
                    min="0" max="10000000000000" step="0.00000001"
                    disabled={limitOrMarket === LimitMarket.Market}
                    value={limitOrMarket === LimitMarket.Market ? "MARKET" : price}
                    onChange={handleChangePrice}
                    autoComplete="off"
                    textAlign="right"
                />
                <InputLabelInside label="Amount" textAlign="right" value={quantity} onChange={handleChangeQuantity} style={{ marginBottom: 10 }} type="number" min="0" max="10000000000000" step="0.00000001" />
            </Grid>
            <Grid item xs={12}>
                <Button fullWidth variant="contained"
                    color={buyOrSell === SellBuy.Buy ? "success" : "error"}
                    onClick={triggerOrder}
                >
                    {buyOrSell === SellBuy.Buy ? "Buy" : "Sell"}
                </Button>
            </Grid>
        </Grid>
        <ProcessingOrderDialog
            open={dialogOpen}
            loading={loading}
            close={() => setDialogOpen(false)}
            error={error}
            order={orderResult}
        />
    </>
    )
}
