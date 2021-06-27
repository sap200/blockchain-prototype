package blockchain

// Implementation using list of blocks

import (
	"blockchain/bank"
	"blockchain/block"
	"blockchain/transaction"
	"blockchain/utils"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Blockchain struct {
	Blocks []block.Block
	Bank   bank.Bank
	Length int64
}

func New() Blockchain {
	bankInstance := bank.New()
	blockchain := Blockchain{
		Blocks: []block.Block{block.GenesisBlock()},
		Bank:   bankInstance,
		Length: 1,
	}

	return blockchain
}

func (blockchain Blockchain) ToJson() string {
	data := struct {
		Blocks []block.Block
		Length int64
	}{
		Blocks: blockchain.Blocks,
		Length: blockchain.Length,
	}

	chainInByteSlice, _ := json.Marshal(data)
	return string(chainInByteSlice)
}

func (blockchain *Blockchain) AddBlock(block_ block.Block) {
	if blockchain.IsValidPrevBlockHash(block_) && blockchain.IsValidBlockCount(block_) {
		blockchain.executeTxs(block_.Transactions)
		blockchain.Blocks = append(blockchain.Blocks, block_)
		blockchain.Length = int64(len(blockchain.Blocks))
	}
}

func (blockchain Blockchain) IsValidPrevBlockHash(currentBlock block.Block) bool {
	lastBlock := blockchain.Blocks[blockchain.Length-1]
	lastBlockHash := hex.EncodeToString(utils.Hash(lastBlock.Payload()))
	return lastBlockHash == currentBlock.PrevBlockHash
}

func (blockchain Blockchain) IsValidBlockCount(currentBlock block.Block) bool {
	lastBlockCount := blockchain.Blocks[blockchain.Length-1].BlockCount
	currentBlockCount := currentBlock.BlockCount

	return (lastBlockCount + 1) == currentBlockCount
}

func (blockchain Blockchain) isTxCovered(tx transaction.Transaction) bool {
	// If transaction is exchange transaction it is always covered
	if tx.TxType == transaction.EXCHANGE {
		return true
	}

	sender := tx.SenderPublicKey
	amount := tx.Amount

	sendersBalance := blockchain.Bank.GetBalance(sender)
	if sendersBalance >= amount {
		return true
	} else {
		return false
	}
}

func (blockchain Blockchain) GetCoveredTxs(txList []transaction.Transaction) []transaction.Transaction {
	coveredTxList := []transaction.Transaction{}

	for _, transaction := range txList {
		txCovered := blockchain.isTxCovered(transaction)
		if txCovered {
			coveredTxList = append(coveredTxList, transaction)
		} else {
			fmt.Println("Transaction not covered", transaction)
		}
	}

	return coveredTxList
}

func (blockchain *Blockchain) executeTx(tx transaction.Transaction) {
	sender := tx.SenderPublicKey
	receiver := tx.ReceiverPublicKey
	amount := tx.Amount

	blockchain.Bank.UpdateBalance(sender, -amount)
	blockchain.Bank.UpdateBalance(receiver, amount)
}

func (blockchain *Blockchain) executeTxs(txList []transaction.Transaction) {
	// Execute exchange tx first
	for _, tx := range txList {
		blockchain.executeTx(tx)
	}
}
