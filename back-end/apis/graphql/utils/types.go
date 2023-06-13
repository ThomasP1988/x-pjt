package utils

import "context"

const AuthType_Authenticated string = "authenticated"

type ResolverArgs struct {
	Args   *map[string]interface{}
	UserID *string
	Event  *ResolverEvent
}

type Resolver = func(context.Context, ResolverArgs) (interface{}, error)

// this file will have to be delete at some point, AWS is very late on providing a decent struct for event

// ResolverType is one of "Query", "Mutation", or "Subscription".
type ResolverType string

const (
	// QueryResolver is the resolver type for query operations.
	QueryResolver ResolverType = "Query"
	// MutationResolver is the resolver type for mutation operations.
	MutationResolver ResolverType = "Mutation"
	// SubscriptionResolver is the resolver type for subscription operations.
	SubscriptionResolver ResolverType = "Subscription"
)

// AuthStrategy is one of "ALLOW" or "DENY"
type AuthStrategy string

const (
	// AllowAuth is allowing by default.
	AllowAuth AuthStrategy = "ALLOW"
	// DenyAuth is denying by default.
	DenyAuth AuthStrategy = "DENY"
)

// For full reference expand to see full JSON here. Some items are left out:
// https://github.com/aws/aws-lambda-go/issues/344#issue-767176415

type (

	// ResolverEvent is the object sent from AWS to the resolver.
	ResolverEvent struct {

		// Arguments are whatever is passed into the resolver as an
		// argument. This will be a JSON object with key as the field
		// and value as the value.
		Arguments map[string]interface{} `json:"arguments"`

		Info     ResolverInfo           `json:"info"`
		Identity ResolverIdentity       `json:"identity"`
		Source   map[string]interface{} `json:"source"`
	}

	// ResolverInfo contains information about the request.
	ResolverInfo struct {

		// ParentTypeName is "Query", "Mutation", or "Subscription".
		ParentTypeName ResolverType `json:"parentTypeName"`

		// FieldName is the resolver field within the parent type.
		FieldName string `json:"fieldName"`

		// SelectionSetList is a list of the requested fields to be
		// returned from the resolver.
		SelectionSetList []string `json:"selectionSetList"`
	}

	// ResolverIdentity contains identity information.
	ResolverIdentity struct {
		CognitoIdentityID           string   `json:"cognitoIdentityId"`
		CognitoIdentityPoolID       string   `json:"cognitoIdentityPoolId"`
		CognitoIdentityAuthType     string   `json:"cognitoIdentityAuthType"`
		AccountID                   string   `json:"accountId"`
		CognitoIdentityAuthProvider string   `json:"cognitoIdentityAuthProvider"`
		SourceIP                    []string `json:"sourceIp"`
	}

	// IdentityClaims contains properties of the identity.
	IdentityClaims struct {
		Aud                 string `json:"aud"`
		AuthTime            int64  `json:"auth_time"`
		CognitoUsername     string `json:"cognito:username"`
		Email               string `json:"email"`
		EmailVerified       bool   `json:"email_verified"`
		EventID             string `json:"event_id"`
		Exp                 int64  `json:"exp"`
		IAT                 int64  `json:"iat"`
		ISS                 string `json:"iss"`
		PhoneNumber         string `json:"phone_number"`
		PhoneNumberVerified bool   `json:"phone_number_verified"`
		Sub                 string `json:"sub"`
		TokenUse            string `json:"token_use"`
	}
)
