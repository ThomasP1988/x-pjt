package trade

import (
	price_helper "NFTM/shared/libs/price"
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2API "github.com/influxdata/influxdb-client-go/v2/api"
	domain "github.com/influxdata/influxdb-client-go/v2/domain"
)

var WriteClient *influxdb2API.WriteAPI
var QueryAPI *influxdb2API.QueryAPI

var orgName string = "nftm"
var client influxdb2.Client

type Bucket string

var tradeBucket Bucket = "trades"
var candle1mBucket Bucket = "candle1m"
var candle5mBucket Bucket = "candle5m"
var candle15mBucket Bucket = "candle15m"
var candle1hBucket Bucket = "candle1h"
var candle6hBucket Bucket = "candle6h"
var candle7dBucket Bucket = "candle7d"

func getClientAndOrganisation() (*influxdb2.Client, *domain.Organization, error) {
	client = influxdb2.NewClientWithOptions("http://127.0.0.1:51147", "salut", influxdb2.DefaultOptions().SetFlushInterval(uint(time.Second)))
	// return nil
	org, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), orgName)
	return &client, org, err
}

func SetClient() error {

	client, org, err := getClientAndOrganisation()

	if err != nil {
		fmt.Printf("err setting client: %v\n", err)
		return err
	}

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:             context.Background(),
		org:             org,
		bucketName:      string(tradeBucket),
		execIfNotExists: func() {},
	})

	writeAPI := (*client).WriteAPI(orgName, string(tradeBucket))
	queryAPI := (*client).QueryAPI(orgName)
	WriteClient = &writeAPI
	QueryAPI = &queryAPI

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:        context.Background(),
		org:        org,
		bucketName: string(candle1mBucket),
		execIfNotExists: func() {
			task, err := (*client).TasksAPI().CreateTaskWithEvery(context.Background(), "trade_to_candles_1m", GetContinuousQueryCandleStick("-5m", "1m", tradeBucket, candle1mBucket), "10s", *org.Id)
			if err != nil {
				fmt.Printf("err creating task: %v\n", err)
			}
			fmt.Printf("task: %v\n", task)
		},
	})

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:        context.Background(),
		org:        org,
		bucketName: string(candle5mBucket),
		execIfNotExists: func() {
			task, err := (*client).TasksAPI().CreateTaskWithEvery(context.Background(), "trade_to_candles_5m", GetContinuousQueryCandleStick("-10m", "5m", candle1mBucket, candle5mBucket), "1m", *org.Id)
			if err != nil {
				fmt.Printf("err creating task: %v\n", err)
			}
			fmt.Printf("task: %v\n", task)
		},
	})

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:        context.Background(),
		org:        org,
		bucketName: string(candle15mBucket),
		execIfNotExists: func() {
			task, err := (*client).TasksAPI().CreateTaskWithEvery(context.Background(), "trade_to_candles_15m", GetContinuousQueryCandleStick("-20m", "15m", candle5mBucket, candle15mBucket), "1m", *org.Id)
			if err != nil {
				fmt.Printf("err creating task: %v\n", err)
			}
			fmt.Printf("task: %v\n", task)
		},
	})

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:        context.Background(),
		org:        org,
		bucketName: string(candle1hBucket),
		execIfNotExists: func() {
			task, err := (*client).TasksAPI().CreateTaskWithEvery(context.Background(), "trade_to_candles_1h", GetContinuousQueryCandleStick("-2h", "1h", candle15mBucket, candle1hBucket), "1m", *org.Id)
			if err != nil {
				fmt.Printf("err creating task: %v\n", err)
			}
			fmt.Printf("task: %v\n", task)
		},
	})

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:        context.Background(),
		org:        org,
		bucketName: string(candle6hBucket),
		execIfNotExists: func() {
			task, err := (*client).TasksAPI().CreateTaskWithEvery(context.Background(), "trade_to_candles_6h", GetContinuousQueryCandleStick("-7h", "6h", candle1hBucket, candle6hBucket), "1m", *org.Id)
			if err != nil {
				fmt.Printf("err creating task: %v\n", err)
			}
			fmt.Printf("task: %v\n", task)
		},
	})

	CreateBucketIfNotExists(CreateBucketArgs{
		ctx:        context.Background(),
		org:        org,
		bucketName: string(candle7dBucket),
		execIfNotExists: func() {
			task, err := (*client).TasksAPI().CreateTaskWithEvery(context.Background(), "trade_to_candles_7d", GetContinuousQueryCandleStick("8d", "7d", candle6hBucket, candle7dBucket), "1m", *org.Id)
			if err != nil {
				fmt.Printf("err creating task: %v\n", err)
			}
			fmt.Printf("task: %v\n", task)
		},
	})

	return nil
}

func Buy(pair string, tradeID string, price int64, amount int64) {
	Trade(pair, tradeID, price, amount, "buy", time.Now())
}

func Sell(pair string, tradeID string, price int64, amount int64) {
	Trade(pair, tradeID, price, amount, "sell", time.Now())
}

func Trade(pair string, tradeID string, price int64, amount int64, side string, timeTrade time.Time) {
	// Create point using fluent style

	fmt.Printf("price: %v\n", price_helper.FromIntToInexactFloat(price))
	fmt.Printf("amount: %v\n", price_helper.FromIntToInexactFloat(amount))
	p := influxdb2.NewPointWithMeasurement(string(tradeBucket)).
		AddTag("pair", pair).
		AddField("tradeId", tradeID).
		AddField("side", side).
		AddField("price", price_helper.FromIntToInexactFloat(price)).
		AddField("amount", price_helper.FromIntToInexactFloat(amount)).
		SetTime(timeTrade)

	(*WriteClient).WritePoint(p)
	(*WriteClient).Flush()

}

func Find(pair string) (*map[string]interface{}, error) {
	// var taskFlux string = `from(bucket:"` + string(tradeBucket) + `")|> range(start: -1h) |> filter(fn: (r) => r.pair == "` + pair + `")`
	var taskFlux string = `from(bucket:"` + string(tradeBucket) + `") 
|> range(start: -1h)`
	result, err := (*QueryAPI).Query(context.Background(), taskFlux)
	if err != nil {
		return nil, err
	}
	fmt.Printf("result: %v\n", result.Next())
	if !result.Next() {
		return &map[string]interface{}{}, nil
	}

	for result.Next() {
		// Observe when there is new grouping key producing new table
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// read result
		fmt.Printf("row: %s\n", result.Record().Values())
	}

	fmt.Printf("result: %v\n", result)
	return nil, nil
}

func Find1m(pair string) (*map[string]interface{}, error) {
	// var taskFlux string = `from(bucket:"` + string(candle1mBucket) + `") |> range(start: -10h) |> filter(fn: (r) => r.pair == "` + pair + `")`
	var taskFlux string = `from(bucket:"` + string(candle1mBucket) + `") |> range(start: -10h)`
	result, err := (*QueryAPI).Query(context.Background(), taskFlux)
	if err != nil {
		return nil, err
	}
	fmt.Printf("result: %v\n", result.Next())
	if !result.Next() {
		return &map[string]interface{}{}, nil
	}

	for result.Next() {
		// Observe when there is new grouping key producing new table
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// read result
		fmt.Printf("row: %s\n", result.Record().Values())
	}

	fmt.Printf("result: %v\n", result)
	return nil, nil
}

type CreateBucketArgs struct {
	ctx             context.Context
	org             *domain.Organization
	bucketName      string
	rules           []domain.RetentionRule
	execIfNotExists func()
}

func CreateBucketIfNotExists(args CreateBucketArgs) {
	bucket, err := client.BucketsAPI().FindBucketByName(args.ctx, args.bucketName)

	if err == nil {
		fmt.Printf("err CreateBucketIfNotExists: %v\n", err)
		return
	}

	if bucket == nil {
		client.BucketsAPI().CreateBucketWithName(context.Background(), args.org, args.bucketName, args.rules...)
		args.execIfNotExists()
	}

}
