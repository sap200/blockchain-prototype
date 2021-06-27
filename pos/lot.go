package pos

import (
	"blockchain/utils"
	"encoding/hex"
)

type Lot struct {
	PublicKeyString string
	Iterations       int
	PrevBlockHash    string
}

func NewLot(pubkey string, iter int, prevhash string) Lot {
	lot := Lot {
		PublicKeyString: pubkey,
		Iterations: iter,
		PrevBlockHash: prevhash,
	}

	return lot
}

func (lot Lot) LotHash() string {
	hashData := lot.PublicKeyString + lot.PrevBlockHash
	for i := 0; i < lot.Iterations; i++ {
		hashData = hex.EncodeToString(utils.Hash(hashData))
	}

	return hashData
}
