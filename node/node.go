package node

import (
	"blockchain/blockchain"
	"blockchain/wallet"
	"blockchain/txpool"
)

type Node struct {
	Blockchain blockchain.Blockchain
	Wallet wallet.Wallet
	TxPool txpool.TxPool
}

func New() Node {
	node := Node {
		Blockchain: blockchain.New(),
		Wallet: wallet.New(),
		TxPool: txpool.New(),
	}

	return node
}