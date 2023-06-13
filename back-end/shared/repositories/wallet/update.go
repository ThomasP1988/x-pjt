package wallet

import (
	entity "NFTM/shared/entities/wallet"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UpdateWalletCurrency struct {
	Own       int64
	Available int64
}

type UpdateWalletArgs struct {
	Ctx        context.Context
	UserID     string
	Currencies map[string]UpdateWalletCurrency
	Wallet     *entity.Wallet
}

func UpdateWallet(args UpdateWalletArgs) error {
	var wg sync.WaitGroup = sync.WaitGroup{}

	errorsUpdate := []error{}
	AssetMapMutex := sync.RWMutex{}

	for k := range args.Currencies {
		wg.Add(1)
		go func(assetID string) {
			result, err := AddSubCurrency(args.Ctx, args.UserID, assetID, args.Currencies[assetID])

			if err != nil {
				// CRITICAL
				fmt.Printf("err: %v\n", err)
				errorsUpdate = append(errorsUpdate, err)
			}

			if args.Wallet != nil {
				AssetMapMutex.Lock()
				args.Wallet.Assets[assetID] = *result
				AssetMapMutex.Unlock()
			}
			defer wg.Done()
		}(k)
	}
	wg.Wait()

	if len(errorsUpdate) > 0 {
		return errors.New("error updating wallet")
	}

	return nil
}

type AddSubValuesToCurrency = UpdateWalletCurrency

func AddSubCurrency(ctx context.Context, userID string, assetID string, values AddSubValuesToCurrency) (*entity.WalletAsset, error) {
	ws := GetWalletService()

	var updateBuild expression.UpdateBuilder = expression.UpdateBuilder{}

	if values.Own != 0 {
		updateBuild = updateBuild.Add(expression.Name("own"), expression.Value(values.Own))
	}

	if values.Available != 0 {
		updateBuild = updateBuild.Add(expression.Name("available"), expression.Value(values.Available))
	}
	updateBuild = updateBuild.Set(expression.Name("lastUpdate"), expression.Value(time.Now().Format(time.RFC3339)))
	updateBuilt := expression.NewBuilder().WithUpdate(updateBuild)
	builder, err := updateBuilt.Build()

	if err != nil {
		fmt.Printf("AddSubCurrency error: %v\n", err.Error())
		return nil, err
	}

	result, err := ws.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{
				Value: userID,
			},
			"assetId": &types.AttributeValueMemberS{
				Value: assetID,
			},
		},
		TableName:                 &ws.TableName,
		ReturnValues:              types.ReturnValueAllNew,
		ExpressionAttributeNames:  builder.Names(),
		ExpressionAttributeValues: builder.Values(),
		UpdateExpression:          builder.Update(),
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	output := &entity.WalletAsset{}

	err = attributevalue.UnmarshalMap(result.Attributes, output)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	return output, nil
}
