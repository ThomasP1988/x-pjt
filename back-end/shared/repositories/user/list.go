package user

import (
	entity "NFTM/shared/entities/user"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UserListArgs struct {
	Limit *int32
	From  *string
}

func List(ctx context.Context, args UserListArgs) (*[]entity.User, *string, error) {
	us := GetUserService()
	users := &[]entity.User{}

	input := dynamodb.ScanInput{
		TableName: &us.TableName,
	}

	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}
	// fmt.Printf("args.From: %v\n", *args.From)
	if args.From != nil && *args.From != "" {
		// input.ExclusiveStartKey = *DeserializeLastKeyUserSymbolIsOpenIndex(*args.From)
		fmt.Printf("input.ExclusiveStartKey: %v\n", input.ExclusiveStartKey)
	}

	output, err := us.Client.Scan(ctx, &input)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("output: %v\n", output)
	err = attributevalue.UnmarshalListOfMaps(output.Items, users)
	if err != nil {
		return nil, nil, err
	}

	var next string

	if len(output.LastEvaluatedKey) > 0 {
		lastKey := &entity.User{}
		err = attributevalue.UnmarshalMap(output.LastEvaluatedKey, lastKey)
		if err != nil {
			return nil, nil, err
		}
		// next = SerializeLastKeyUserSymbolIsOpenIndex(lastKey.LastModified.Format(time.RFC3339), lastKey.ID, lastKey.UserIDSymbolIsOpen)
	}

	return users, &next, nil
}
