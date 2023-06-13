package cognito

import (
	"NFTM/shared/config"
	"NFTM/shared/entities/auth"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func (service CognitoService) Auth(token string) (*auth.UserAuth, error) {

	jwkURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", *config.Conf.Region, config.Conf.User.UserPool)
	jwkResult, err := GetJWK(jwkURL)

	if err != nil {
		return nil, err
	}

	jwtResult, err := ParseJWT(token, jwkResult.Keys)

	if err != nil {
		return nil, err
	}

	if !jwtResult.Valid {
		return nil, errors.New("invalid JWT")
	}

	return &auth.UserAuth{
		ID: jwtResult.Claims.(jwt.MapClaims)["sub"].(string),
	}, nil

}

func GetJWK(jwkURL string) (*JWK, error) {
	req, err := http.NewRequest("GET", jwkURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jwk := new(JWK)
	err = json.Unmarshal(body, jwk)
	if err != nil {
		return nil, err
	}

	return jwk, nil
}

func ParseJWT(tokenString string, keys []JWKKey) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		key := convertKey(keys[1].E, keys[1].N)
		return key, nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}

func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey
}
