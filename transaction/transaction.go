package transaction

// This package is the heart of the blockchain
// Transactions are the atomic entities that changes the state of the blockchain
// transaction package includes some constants and Transaction types and some other helper methods

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// This is the transaction type
// It depicts with integer whether the transaction is of type
// TRANSFER, EXCHANGE or STAKE
type TransactionType int

// These are the three constants encoding the 3 types of
// transactions described above
const (
	TRANSFER TransactionType = 0
	EXCHANGE TransactionType = 1
	STAKE    TransactionType = 2
)

// The transaction type
// It has 3 key components
// From , To and Amount
// The TxType and timestamp and Id are used to distinguish among different transactions
// And of course the signature is used to verify the sender of this transaction
type Transaction struct {
	SenderPublicKey   string
	ReceiverPublicKey string
	Amount            float64
	TxType            TransactionType
	Id                uuid.UUID
	Timestamp         time.Time
	Signature         string
}

// This is a Marshalling method
// It marshalls the whole Transaction struct type
// and returns its string representation
// for better visualization
func (tx Transaction) ToJson() string {
	txInByteSlice, _ := json.Marshal(tx)
	return string(txInByteSlice)
}

// This method is used to create the new transaction instance
// uuid.New() generates a new UUID , that uniquely identifies a transaction
// time.now() returns the current time
func New(senderPublicKey string, receiverPublicKey string, amount float64, txType TransactionType) Transaction {
	tx := Transaction{
		SenderPublicKey:   senderPublicKey,
		ReceiverPublicKey: receiverPublicKey,
		Amount:            amount,
		TxType:            txType,
		Id:                uuid.New(),
		Timestamp:         time.Now(),
		Signature:         "",
	}
	return tx
}

// This method is used to set the signature parameter of the struct
// The signature is generated by wallet class
// The pointer is used so that the change is directly reflected in the data structure
func (tx *Transaction) Sign(signature string) {
	tx.Signature = signature
}

// The payload method is used for consistent representation of the transaction
// This is basically whole transaction without signature
// It is necessary because setting the signature changes the checksum of the transaction
// And hence consistent representation lets us verify the transaction sender
func (tx Transaction) Payload() string {
	tx.Signature = ""
	return tx.ToJson()
}

// Equals method basically is a helper method
// It is used to prove the 2 transactions are different by comparing their uuids
// It first marshals the uuid into a byte slice
// Then the byte slice is compared and the equality is established
// It is an implementation from github.com/google/uuid
func (tx Transaction) Equals(newTx Transaction) bool {
	idTx, _ := tx.Id.MarshalBinary()
	idNewTx, _ := newTx.Id.MarshalBinary()

	result := bytes.Compare(idTx, idNewTx)

	if result == 0 {
		return true
	}

	return false
}
