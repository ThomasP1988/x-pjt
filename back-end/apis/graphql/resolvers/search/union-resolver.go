package search

import (
	"NFTM/shared/entities/nft"
	"encoding/json"
	"errors"
	"fmt"
)

type SearchResult interface {
	nft.Collection | nft.Item
}

const typename string = "__typename"

func ResolveArray[T SearchResult](input []T) []interface{} {
	result := []any{}

	for _, v := range input {
		resolvedEntity, err := ResolverUnion(v)

		if err != nil {
			fmt.Printf("ResolveArray err: %v\n", err)
			result = append(result, err)
			continue
		}
		result = append(result, resolvedEntity)
	}

	return result
}

func ResolverUnion[T SearchResult](input T) (map[string]interface{}, error) {
	switch v := any(input).(type) {
	default:
		fmt.Printf("unexpected type %T", v)
		return nil, errors.New("unknown type for union resolver")
	case nft.Collection:
		return SetTypename(input, "Collection")
	}
}

func SetTypename[T SearchResult](input T, name string) (map[string]interface{}, error) {
	var outInterface map[string]interface{}
	inrec, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &outInterface)
	if err != nil {
		return nil, err
	}
	outInterface[typename] = name
	return outInterface, nil
}
