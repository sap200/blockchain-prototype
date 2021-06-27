package pos

import (
	"math/big"
	"blockchain/utils"
	"encoding/hex"
)

type ProofOfStake struct {
	Stakers map[string]int
}

func NewProofOfStake() ProofOfStake {
	pos := ProofOfStake {
		Stakers: map[string]int{},
	}

	return pos
}

func (pos *ProofOfStake) UpdateStake(publickKeyString string, amount int) {
	_, found := pos.Stakers[publickKeyString]
	if !found {
		pos.Stakers[publickKeyString] = amount
	} else {
		pos.Stakers[publickKeyString] += amount
	}
}

func (pos ProofOfStake) GetStakes(publickKeyString string) int {
	val, _ := pos.Stakers[publickKeyString]
	return val
}

func (pos ProofOfStake) ValidatorLots(seed string) []Lot {
	lots := []Lot{}
	for pubKey, stakes := range pos.Stakers {
		for i := 0; i < stakes; i++ {
			lot := NewLot(pubKey, i+1, seed)
			lots = append(lots, lot)
		}
	}

	return lots
}

func (pos ProofOfStake) WinnerLot(lots []Lot, seed string) Lot {
	var winnerLot Lot
	leastOffset := new(big.Int)
	offset := new(big.Int)
	referenceLot := new(big.Int)
	referenceLot.SetString(hex.EncodeToString(utils.Hash(seed)), 16)

	for index, lot := range lots {
		hashData := lot.LotHash()
		hashDataInt := new(big.Int)
		hashDataInt.SetString(hashData, 16)
		offset.Sub(hashDataInt, referenceLot)

		if offset.CmpAbs(leastOffset) == -1 || index == 0 {
			leastOffset = leastOffset.Set(offset)
			winnerLot = lot
		}
	}
	return winnerLot
}

func (pos ProofOfStake) Forger(lastBlockHash string) string {
	validatorLots := pos.ValidatorLots(lastBlockHash)
	winnerLot := pos.WinnerLot(validatorLots, lastBlockHash)
	return winnerLot.PublicKeyString 
}