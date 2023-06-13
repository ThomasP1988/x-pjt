package dynamodb_helper

import (
	"NFTM/shared/constants"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	Delimitator_VALUE    string = "+"
	Deliminator_PROPERTY string = "/"
)

// func Serialize(input interface{}) *string {
func Serialize(lastEvaluatedKey map[string]types.AttributeValue, typeEntity interface{}) (*string, error) {

	if len(lastEvaluatedKey) == 0 {
		return nil, nil
	}

	var output []string = []string{}
	err := attributevalue.UnmarshalMap(lastEvaluatedKey, typeEntity)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	v := reflect.Indirect(reflect.ValueOf(typeEntity))
	typesCollection := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			typeField := typesCollection.Field(i)
			nameField := typeField.Name

			if reflect.TypeOf(time.Now()).String() == typeField.Type.String() {
				fmt.Printf("field.Interface().(time.Time).Format(time.RFC3339): %v\n", field.Interface().(time.Time).Format(time.RFC3339Nano))
				output = append(output, url.PathEscape(nameField)+Delimitator_VALUE+url.PathEscape(field.Interface().(time.Time).Format(time.RFC3339Nano)))
			} else {
				output = append(output, url.PathEscape(nameField)+Delimitator_VALUE+url.PathEscape(fmt.Sprint(field.Interface())))
			}

		}
	}

	result := strings.Join(output, Deliminator_PROPERTY)

	return &result, nil
}

func Deserialize(input string, typeEntity interface{}) (*map[string]types.AttributeValue, error) {

	typesEntity := reflect.ValueOf(typeEntity).Type()

	output := map[string]types.AttributeValue{}
	properties := strings.Split(input, Deliminator_PROPERTY)

	for _, v := range properties {
		keyValue := strings.Split(v, Delimitator_VALUE)

		field, exists := typesEntity.FieldByName(keyValue[0])
		if !exists {
			return nil, ErrDeserializedUnknownField
		}
		key := field.Tag.Get(constants.Tag_DynamoDB)
		if field.Type.Kind() == reflect.String {
			unescapedValue, err := url.PathUnescape(keyValue[1])

			if err != nil {
				fmt.Printf("err: %v\n", err)
				return nil, err
			}

			output[key] = &types.AttributeValueMemberS{
				Value: unescapedValue,
			}
		} else if field.Type.Kind() >= reflect.Int && field.Type.Kind() <= reflect.Complex128 {
			output[key] = &types.AttributeValueMemberN{
				Value: keyValue[1],
			}
		} else if field.Type.Kind() < reflect.Bool {

			boolValue, err := strconv.ParseBool(keyValue[1])

			if err != nil {
				fmt.Printf("err: %v\n", err)
				return nil, ErrDeserializedUnparsableBool
			}

			output[key] = &types.AttributeValueMemberBOOL{
				Value: boolValue,
			}
		} else if field.Type.Kind() == reflect.Struct {

			switch field.Type.String() {
			case reflect.TypeOf(time.Now()).String():
				unescapedValue, err := url.PathUnescape(keyValue[1])

				if err != nil {
					fmt.Printf("err: %v\n", err)
					return nil, err
				}

				output[key] = &types.AttributeValueMemberS{
					Value: unescapedValue,
				}
			default:
				return nil, ErrDeserializedUnknownTypeField
			}
		} else {
			return nil, ErrDeserializedUnknownTypeField
		}

	}

	return &output, nil
}
