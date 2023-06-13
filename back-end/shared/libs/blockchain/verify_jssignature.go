package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type EIP191 struct {
	msg       string
	signature string
}

func VerifySignatureFromJS(signature string, nonce string) (*string, error) {
	fmt.Printf("signature: %v\n", signature)
	fmt.Printf("nonce: %v\n", nonce)
	address, err := decodePersonal(EIP191{
		signature: signature,
		msg:       nonce,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return address, nil
}

func hasValidLastByte(sig []byte) bool {
	return sig[64] == 0 || sig[64] == 1
}

func signEIP191(message string) common.Hash {
	msg := []byte(message)
	formattedMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	return crypto.Keccak256Hash([]byte(formattedMsg))
}

func decodePersonal(eipChallenge EIP191) (*string, error) {
	decodedSig, err := hexutil.Decode(eipChallenge.signature)
	if err != nil {
		return nil, err
	}

	if decodedSig[64] < 27 {
		if !hasValidLastByte(decodedSig) {
			panic("Invalid last byte")
		}
	} else {
		decodedSig[64] -= 27 // shift byte?
	}

	hash := signEIP191(eipChallenge.msg)

	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	if err != nil {
		return nil, err
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		return nil, err
	}

	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()

	return &recoveredAddress, nil
}
