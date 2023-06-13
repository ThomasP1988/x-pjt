package helper

import (
	"NFTM/shared/config"
	"fmt"

	"github.com/aws/jsii-runtime-go"
)

var stage config.Stage

func SetStage(newStage config.Stage) {
	stage = newStage
}

func SetName(name string) *string {
	fmt.Printf("stage: %v\n", stage)
	return jsii.String(string(stage) + "-" + name)
}
