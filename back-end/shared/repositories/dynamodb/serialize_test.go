package dynamodb_helper

import (
	"NFTM/shared/config"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func setup() {
	config.GetConfig(nil)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

type SerializedTest struct {
	Animals   int       `dynamodbav:"animals"`
	Country   string    `dynamodbav:"country"`
	CreatedAt time.Time `dynamodbav:"createdAt"`
	Unused    string    `dynamodbav:"unused"`
}

func TestSerializeDeserialized(t *testing.T) {

	input := map[string]types.AttributeValue{
		"animals": &types.AttributeValueMemberN{
			Value: "5",
		},
		"country": &types.AttributeValueMemberS{
			Value: "ukraine",
		},
		"createdAt": &types.AttributeValueMemberS{
			Value: time.Now().Format(time.RFC3339),
		},
	}

	serialisedString, err := Serialize(input, &SerializedTest{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("serialisedString: %v\n", *serialisedString)

	result, err := Deserialize(*serialisedString, SerializedTest{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("result: %v\n", result)

	t.Fail()
}
