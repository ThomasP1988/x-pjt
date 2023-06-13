import React, { useState, useEffect } from 'react';
import { Orderbook } from "shared/components/Orderbook";
import { OrderHistory } from "../../containers/OrderHistory";
import Grid from '@mui/material/Grid';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Container from '@mui/material/Container';
import Stack from '@mui/material/Stack';
import { MyOrders } from '../MyOrders';
import { Candlestick } from 'shared/components/Chart/Candlestick';
import { useParams } from "react-router-dom";
import { ErrorPage } from 'shared/components/ErrorPage';
import { useNavigate } from "react-router-dom";

enum HighTabs {
    Market = "market",
    Orders = "myOrders",
    History = "history"
}

export const Market = () => {
    const { symbol, tab } = useParams();
    const [selectedTab, selectTab] = useState<HighTabs>(tab && Object.values(HighTabs).includes(tab as HighTabs) ? tab as HighTabs : HighTabs.Market)
    const navigate = useNavigate();

    useEffect(() => {
        console.log("useEffect")
        navigate(`/m/${symbol}/${selectedTab}`);
    }, [navigate, selectedTab, symbol]);

    return (
        <Container>
            {
                !symbol ? <ErrorPage text="Oops! wrong URL, nothing to see here." hideTryAgain={true} /> : <>
                    <Stack justifyContent="flex-end" sx={{ padding: 1 }}>
                        <Tabs
                            value={selectedTab}
                            onChange={(_, v) => selectTab(v)}
                            textColor="secondary"
                            indicatorColor="secondary"

                        >
                            <Tab value={HighTabs.Market} label="Market" />
                            <Tab value={HighTabs.Orders} label="My Orders" />
                            <Tab value={HighTabs.History} label="My History" />
                        </Tabs>
                    </Stack>
                    <Grid container>
                        {
                            selectedTab === HighTabs.Market && <>
                                <Grid item md={9}>
                                    <Candlestick symbol={symbol} />
                                </Grid>
                                <Grid item md={3}>
                                    <Orderbook symbol={symbol} />
                                </Grid>
                            </>
                        }
                        {
                            selectedTab === HighTabs.Orders && <>
                                <MyOrders symbol={symbol} />
                            </>
                        }
                        {
                            selectedTab === HighTabs.History && <>
                                <OrderHistory symbol={symbol} />
                            </>
                        }
                    </Grid>
                </>
            }
        </Container>
    );
}


