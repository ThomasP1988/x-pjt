package wallet

import (
	entity "NFTM/shared/entities/wallet"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetWallet(ctx context.Context, userID string, assets []string) (*entity.Wallet, error) {

	wallet := entity.Wallet{
		UserID: userID,
		Assets: make(map[string]entity.WalletAsset),
	}
	AssetMapMutex := sync.RWMutex{}
	var wg sync.WaitGroup = sync.WaitGroup{}

	errorsGet := []error{}

	for i := range assets {
		wg.Add(1)
		go func(assetID string) {
			result, err := GetAsset(ctx, userID, assetID)

			if err != nil {
				// CRITICAL
				fmt.Printf("err: %v\n", err)
				errorsGet = append(errorsGet, err)
				return
			}

			if result == nil {
				newAsset := &entity.WalletAsset{
					UserID:  userID,
					AssetID: assetID,
				}
				err = Add(ctx, newAsset)

				if err != nil {
					errorsGet = append(errorsGet, err)
					return
				}
				result = newAsset
			}

			AssetMapMutex.Lock()
			wallet.Assets[assetID] = *result
			AssetMapMutex.Unlock()

			defer wg.Done()
		}(assets[i])
	}
	wg.Wait()

	if len(errorsGet) > 0 {
		return nil, errors.New("error updating wallet")
	}

	return &wallet, nil

}

func GetAsset(ctx context.Context, userID string, assetID string) (*entity.WalletAsset, error) {
	ws := GetWalletService()

	newCond := expression.Key("userId").Equal(expression.Value(userID)).And(expression.Key("assetId").Equal(expression.Value(assetID)))

	expr, err := expression.NewBuilder().WithKeyCondition(newCond).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 &ws.TableName,
	}

	queryOutput, err := ws.Client.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	if len(queryOutput.Items) == 0 {
		return nil, nil
	}

	output := &entity.WalletAsset{}

	err = attributevalue.UnmarshalMap(queryOutput.Items[0], output)

	if err != nil {
		println("failed to unmarshal Items", err.Error())
		return nil, err
	}

	return output, nil
}
