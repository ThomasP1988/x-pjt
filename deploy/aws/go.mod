module aws

go 1.18

require (
	NFTM/shared v0.0.0-00010101000000-000000000000
	github.com/aws/aws-cdk-go/awscdk/v2 v2.28.1
	github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2 v2.28.1-alpha.0
	github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2 v2.28.1-alpha.0
	github.com/aws/constructs-go/constructs/v10 v10.1.42
	github.com/aws/jsii-runtime-go v1.61.0
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/aws/aws-sdk-go-v2 v1.14.0 // indirect
	github.com/aws/smithy-go v1.11.0 // indirect
)

replace NFTM/shared => ../../back-end/shared
