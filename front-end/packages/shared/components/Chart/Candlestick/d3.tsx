import { useRef, useEffect } from "react";
import * as d3 from "d3";
import { D3ZoomEvent } from "d3";

export type CandleStickInput = {
    Date: string,
    Open: number,
    High: number,
    Low: number,
    Close: number,
    "Adj Close": number,
    Volume: number
}

export type CandleStickOptions = {
    input: CandleStickInput[];
    width: number | string;
    height: number | string;
    hoverItem?: (item: CandleStickInput | null) => void;
    colorUp: string,
    colorStill: string,
    colorDown: string,
    colorTooltip: string
}

const margin: Record<string, number> = {
    bottom: 20,
    top: 20,
    left: 10,
    right: 90
}

export function CandlestickChart({
    input,
    width,
    height,
    colorUp,
    colorStill,
    colorDown,
    colorTooltip,
    hoverItem
}: CandleStickOptions) {
    const svgRef = useRef<SVGSVGElement>(null);

    useEffect(() => {
        console.log("useEffect")
        const svgEl = d3.select<SVGSVGElement, unknown>(svgRef.current as SVGSVGElement);
        svgEl.selectAll("*").remove(); // Clear svg content before adding new elements 
        const widthSVG = svgEl.property("width").baseVal.value;
        const heightSVG = svgEl.property("height").baseVal.value;

        const svg = svgEl
            .append("g")
            .attr("transform", `translate(${margin.left},${margin.top})`);

        // format data 

        let minPrice: number = 0;
        let maxPrice: number = 0;
        let minDateTimestamp: number = 0;
        let minDate: Date | undefined;
        let maxDateTimestamp: number = 0;
        let maxDate: Date | undefined;

        const inputByTimestamp: Record<number, CandleStickInput> = {};
        const allDates: Date[] = [];
        const allLows: number[] = [];
        const allHighs: number[] = [];
        const allOpens: number[] = [];
        const allCloses: number[] = [];
        const allDatesRange = d3.range(input.length);
        const pseudoBandwith = (((widthSVG - margin.left - margin.right) / input.length)) / 2;

        for (let i = 0; i < input.length; i++) {
            const dateDate: Date = new Date(input[i].Date);
            const dateTimestamp: number = +dateDate;

            allDates.push(dateDate);
            allLows.push(input[i].Low);
            allHighs.push(input[i].High);
            allOpens.push(input[i].Open);
            allCloses.push(input[i].Close);

            inputByTimestamp[dateTimestamp] = input[i];

            if (input[i].Low < minPrice || minPrice === 0) {
                minPrice = input[i].Low;
            }
            if (input[i].High > maxPrice) {
                maxPrice = input[i].High;
            }
            if (dateTimestamp < minDateTimestamp || minDateTimestamp === 0) {
                minDateTimestamp = dateTimestamp;
                minDate = dateDate;
            }
            if (dateTimestamp > maxDateTimestamp) {
                maxDateTimestamp = dateTimestamp;
                maxDate = dateDate;
            }

        }
        // end format data


        // Y line
        const yScale = d3.scaleLinear()
            .domain([Math.floor(minPrice), Math.ceil(maxPrice)])
            // .range([Math.floor(min), Math.ceil(max)]);
            .range([heightSVG - margin.top - margin.bottom, 0]);

        svg.append("g")
            .call(d3.axisRight(yScale).ticks(8))
            .attr("transform", "translate(" + (widthSVG - margin.left - margin.right) + ",0)")
            .call(g => g.select(".domain").remove()) // hide native axis line
            .call(g => g.selectAll(".tick line").clone()
                .attr("stroke-opacity", 0.2)
                .attr("x2", margin.right + margin.left - widthSVG))

        // X line

        const xScale: d3.ScaleTime<number, number, never> = d3.scaleTime()
            .domain([maxDate as Date, minDate as Date]) // reverse, why?
            .range([widthSVG - margin.right - margin.left, 0]);


        const tickRange: Date[] = d3.timeTicks(minDate as Date, maxDate as Date, 8)


        const gX = svg.append("g")
            .attr("width", widthSVG)
            .attr("transform", "translate(0," + (heightSVG - margin.top - margin.bottom) + ")")
            .call(d3.axisBottom(xScale)
                .tickValues(tickRange.map(i => i))
                // .tickFormat((d) => d3.timeFormat("%Y-%m-%d")(d as Date))
                // .ticks(8)
            )
            .call(g => g.select(".domain").remove());

        // end X line

        // prevent zoom from getting out of selected area
        // Add a clipPath: everything out of this area won't be drawn.
        svg.append("defs").append("SVG:clipPath")
            .attr("id", "clip")
            .append("SVG:rect")
            .attr("width", widthSVG - margin.left - margin.right)
            .attr("height", heightSVG - margin.top - margin.bottom)
            .attr("x", 0)
            .attr("y", 0);

        // Create the scatter variable: where both the circles and the brush take place
        const drawableArea = svg.append('g')
            .attr("clip-path", "url(#clip)")

        // draw candles
        const candles = drawableArea.append("g")
            .attr("stroke", "black")
            .attr("stroke-linecap", "grey")
            .selectAll("g")
            .data(allDatesRange)
            .join("g")
            .attr("transform", (i) => {
                return `translate(${xScale(allDates[i])},0)`
            });

        const lowHighLine = candles.append("line")
            .attr("y1", i => yScale(allLows[i]))
            .attr("y2", i => yScale(allHighs[i]))
            .attr("stroke-width", pseudoBandwith / 5);

        const openCloseLine = candles.append("line")
            .attr("y1", i => yScale(allOpens[i]))
            .attr("y2", i => yScale(allCloses[i]))
            .attr("stroke-width", pseudoBandwith)
            // .attr("stroke-width", xScale.bandwidth())
            .attr("stroke", i => [colorUp, colorStill, colorDown][1 + Math.sign(allOpens[i] - allCloses[i])]);
        // end draw candles

        // zoom

        const extent: [[number, number], [number, number]] = [[0, 0], [widthSVG - margin.left - margin.right, heightSVG - margin.top - margin.bottom]];
        let xScaleZ: d3.ScaleTime<number, number, never> | undefined;
        const zoomSvg = d3.zoom<SVGSVGElement, unknown>()
            .scaleExtent([1, 100])
            .translateExtent(extent)
            .extent(extent)
            .on("zoom", (event: D3ZoomEvent<SVGSVGElement, unknown>) => {
                // var t = transform;

                xScaleZ = event.transform.rescaleX(xScale);

                gX.call(
                    d3.axisBottom(xScaleZ)
                )

                const newBandwith = pseudoBandwith * event.transform.k
                candles.attr("transform", (d, i) => `translate(${xScaleZ?.(allDates[i])},0)`);
                openCloseLine.attr("stroke-width", (d, i) => newBandwith);
                lowHighLine.attr("stroke-width", newBandwith / 5 > 4 ? 4 : newBandwith / 5);

            })
            .on('zoom.end', () => {

            });

        svgEl.call(zoomSvg);
        // end zoom

        // mouse over drawing lines and showing legend

        const legendPriceContainer = svg.append("g")
            .style("opacity", "0")
            .attr("x", "-6")

        legendPriceContainer.append("rect")
            .attr('width', margin.right) // can't catch mouse events on a g element
            .attr('height', 16)
            .attr("fill", colorTooltip)

        const legendPrice = legendPriceContainer.append("text").text("test")
            .style("font-size", "10px")
            .attr("x", "3")
            .attr("y", "11");


        const legendDateContainer = svg.append("g")
            .style("opacity", "0");

        legendDateContainer.append("rect")
            .attr('width', 110)
            .attr('height', 16)
            .attr("fill", colorTooltip)

        const legendDate = legendDateContainer.append("text").text("test")
            .style("font-size", "10px")
            .attr("x", "3")
            .attr("y", "11");


        const mouseG = drawableArea.append("g")
            .attr("class", "mouse-over-effects")
            .style("cursor", "crosshair");

        const verticalLine = mouseG.append("path") // this is the black vertical line to follow mouse
            .attr("class", "mouse-line")
            .style("stroke", "black")
            .style("stroke-dasharray", "5")
            .style("stroke-width", "1px")
            .style("opacity", "0");

        const horizontalLine = mouseG.append("path") // this is the black vertical line to follow mouse
            .attr("class", "mouse-line")
            .style("stroke", "black")
            .style("stroke-dasharray", "5")
            .style("stroke-width", "1px")
            .style("opacity", "0");

        const bisect = d3.bisector((input: CandleStickInput) => {
            return +new Date(input.Date);
        }).center;

        mouseG.append('svg:rect') // append a rect to catch mouse movements on canvas
            .attr('width', width) // can't catch mouse events on a g element
            .attr('height', height)
            .attr('fill', 'none')
            .attr('pointer-events', 'all')
            .on('mouseout', function () { // on mouse out hide line, circles and text
                verticalLine.style("opacity", "0");
                horizontalLine.style("opacity", "0");
                legendPriceContainer.style("opacity", "0");
                legendDateContainer.style("opacity", "0");
                hoverItem?.(null);
            })
            .on('mouseover', function () { // on mouse in show line, circles and text
                verticalLine.style("opacity", "1");
                horizontalLine.style("opacity", "1");
                legendPriceContainer.style("opacity", "1");
                legendDateContainer.style("opacity", "1");
            })
            .on('mousemove', function (event: MouseEvent) { // mouse moving over canvas

                let coords = d3.pointer(event);

                const xPointed: Date = (xScaleZ ? xScaleZ : xScale).invert(coords[0]);
                const yPointed: number = yScale.invert(coords[1]);

                verticalLine.attr("d", () => {
                    var d = "M" + (coords[0]) + "," + heightSVG;
                    d += " " + (coords[0]) + "," + 0;
                    return d;
                });
                horizontalLine.attr("d", () => {
                    var d = "M" + widthSVG + "," + (coords[1]);
                    d += " " + 0 + "," + (coords[1]);
                    return d;
                });

                const widthWithoutMargin: number = widthSVG - margin.left - margin.right;

                legendPriceContainer
                    .attr("transform", `translate(${widthWithoutMargin + 10},${coords[1] - 8})`);

                legendPrice
                    .text(yPointed.toFixed(8));

                let legendDateX: number = coords[0] - 55;

                if (legendDateX < 0) {
                    legendDateX = 0;
                } else if (legendDateX > widthWithoutMargin - 110) {
                    legendDateX = widthWithoutMargin - 110;
                }

                legendDateContainer
                    .attr("transform", `translate(${legendDateX},${heightSVG - margin.bottom - margin.top})`)

                legendDate
                    .text(xPointed.toLocaleString())

                // console.log("xPointed", xPointed);
                const index = bisect(input, +xPointed);
                hoverItem?.(input[index]);
            });

    }, [input, width, height, hoverItem,  colorUp,colorStill,colorDown,colorTooltip]);

    return <svg ref={svgRef} width={width} height={height} style={{ display: "block" }} />;
}
