package contract

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
)

func CreateAddress(sEoa string, nonce uint64) (string) {
	aEoa := common.HexToAddress(sEoa)
	return crypto.CreateAddress(aEoa, nonce).String()
}
