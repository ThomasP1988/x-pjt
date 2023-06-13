package cognito

type JWKKey struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
}

type JWK struct {
	Keys []JWKKey `json:"keys"`
}
