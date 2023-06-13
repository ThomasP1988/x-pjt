package wallet

import (
	entity "NFTM/shared/entities/wallet"
	"context"
	"errors"
	"fmt"
	"sync"
)

func DeleteByUser(ctx context.Context, userID string) error {
	var next *string
	var assets *[]entity.WalletAsset
	var err error
	var wg sync.WaitGroup = sync.WaitGroup{}

	errorsList := []error{}
	errorsDelete := []error{}

	for {
		assets, next, err = List(ctx, ListArgs{
			UserID: userID,
			From:   next,
		})

		if err != nil {
			fmt.Printf("error DeleteByUser: %v\n", err)
			errorsList = append(errorsList, err)
		}

		for i := range *assets {
			wg.Add(1)
			go func(assetID string) {

				err := Delete(ctx, userID, assetID)

				if err != nil {
					fmt.Printf("err DeleteByUser delete item: %v\n", err)
					errorsDelete = append(errorsDelete, err)
				}

				defer wg.Done()
			}((*assets)[i].AssetID)
		}

		if next == nil {
			break
		}
	}

	wg.Wait()

	if len(errorsList) > 0 {
		fmt.Printf("error listing wallet asset in delete by user")
		return errors.New("error deleting wallet asset")
	}

	if len(errorsDelete) > 0 {
		fmt.Printf("error deleting one or more single asset in delete by user")
		return errors.New("error deleting wallet asset")
	}

	return nil
}
