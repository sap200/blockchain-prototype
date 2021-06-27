package block

import (
	"blockchain/transaction"
	"encoding/json"
	"time"
)

type Block struct {
	Transactions  []transaction.Transaction
	PrevBlockHash string
	BlockCount    int64
	Forger        string
	Timestamp     time.Time
	Signature     string
}

func New(transactions []transaction.Transaction, prevBlockHash string, blockCount int64, forger string) Block {
	block := Block{
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		BlockCount:    blockCount,
		Forger:        forger,
	}

	return block
}

func (b Block) ToJson() string {
	blockInByteSlice, _ := json.Marshal(b)
	return string(blockInByteSlice)
}

func (b Block) Payload() string {
	b.Signature = ""
	return b.ToJson()
}

func (b *Block) Sign(signature string) {
	b.Signature = signature
}

func GenesisBlock() Block {
	genesis := New([]transaction.Transaction{}, "genesisHash", 0, "genesis")
	genesis.Timestamp, _ = time.Parse(time.UnixDate, time.UnixDate)
	return genesis
}
