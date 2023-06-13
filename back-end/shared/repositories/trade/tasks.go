package trade

func GetContinuousQueryCandleStick(rangeTime string, window string, from Bucket, to Bucket) string {
	return `from(bucket:"` + string(from) + `") 
|> range(start: ` + rangeTime + `)
|> window(every: ` + window + `)
|> first("price")
|> last("price")
|> min("price")
|> max("price")
|> mean("price")
|> cumulativeSum("amount")
|> group(columns: ["pair"])
|> to(bucket: "` + string(to) + `", org: "` + orgName + `")`
}

// from(bucket:"trades")
// |> range(start: -5m)
// |> window(every: 1m)
// |> map(fn: (r) => r["price"] == "jobs")
// |> first("result.price")
// |> last("result.price")
// |> min("result.price")
// |> max("result.price")
// |> mean("result.price")
// |> cumulativeSum("result.amount")
// |> group(columns: ["pair"])
// |> to(bucket: "candle1m", org: "nftm")
