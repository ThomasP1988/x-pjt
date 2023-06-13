import { useEffect, useState, useCallback } from "react";

import { Paper, Stack, ToggleButtonGroup, ToggleButton, Grid, Typography } from "@mui/material";
import CandlestickChartIcon from '@mui/icons-material/CandlestickChart';
import ShowChartIcon from '@mui/icons-material/ShowChart';

import { CandlestickChart, CandleStickInput } from "./d3";
import aapl from "./aapl.json"
import { useColors, Color } from "../../../lib/hooks/colors";

type Props = {
    symbol: string
}

export const Candlestick = ({ symbol }: Props) => {

    const colors = useColors();

    const [typeChart, setTypeChart] = useState<string>("candlestick")
    const [hoveredInput, setHoveredInput] = useState<CandleStickInput | null>(null);
    const [direction, setDirection] = useState<number>(0);
    const [colorsDirection] = useState<Color[]>([colors.green, colors.grey, colors.red]);

    useEffect(() => {
        // fetchTestData().then((result) => {
        //     console.log("result", result);
        //     drawChart(result);
        // }).catch((e) => {
        //     console.log(e);
        // });
    }, [])

    const hoverInput = useCallback((item: CandleStickInput | null) => {
        if (item) {
            setHoveredInput(item);
            setDirection(Math.sign(item.Open - item.Close));
        }
    }, [setHoveredInput])

    const topIndicator = (label: string, value: string | number, color?: Color) => {
        return <Grid item md={2}>
            <Typography variant="subtitle2">{label}</Typography> <Typography
                variant="subtitle2"
                color={color}
            >{value}</Typography>
        </Grid>
    }

    return (<>
        <Stack direction="row" justifyContent="flex-end" sx={{ margin: 1 }}>
            <ToggleButtonGroup value={typeChart}
                size="small"
                onChange={(_: React.MouseEvent<HTMLElement, MouseEvent>,
                    value: string | null) => {
                    console.log("value", value)
                    setTypeChart(value as string);
                }} >
                <ToggleButton value="candlestick"><CandlestickChartIcon /></ToggleButton>
                <ToggleButton value="area"><ShowChartIcon /></ToggleButton>
            </ToggleButtonGroup>
        </Stack>
        <Paper sx={{ padding: 2, margin: 1 }}>
            <Stack sx={{ height: 50 }}>
                {
                    hoveredInput && <Grid container justifyContent="center" spacing={2}>
                        {topIndicator("Date", (new Date(hoveredInput.Date)).toLocaleDateString())}
                        {topIndicator("Open", hoveredInput.Open, colorsDirection[1 + direction])}
                        {topIndicator("Close", hoveredInput.Close, colorsDirection[1 + direction])}
                        {topIndicator("High", hoveredInput.High, colorsDirection[1 + direction])}
                        {topIndicator("Low", hoveredInput.Low, colorsDirection[1 + direction])}
                    </Grid>
                }
            </Stack>
            <Stack sx={{ height: 400 }}>
                <CandlestickChart input={aapl} width={"100%"} height={"100%"} hoverItem={hoverInput}
                    colorUp={colors.green}
                    colorStill={colors.grey}
                    colorDown={colors.red}
                    colorTooltip={colors.lightgrey}
                />
            </Stack>
        </Paper>
    </>
    )
}