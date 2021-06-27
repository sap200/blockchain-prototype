package txpool

// In Blockchain the state of the blockchain is changed by transactions
// But single block containing single transactions leads to unnecesary increase of the chain
// Hence the transactions are grouped together in TransactionPool before being written to the forged block
// Transaction pool  basically contains a slice of transactions that can hold unlimited transactions

import (
	"blockchain/transaction"
	"encoding/json"
)

// The Transaction Pool structure
type TxPool struct {
	// The slice that is supposed to contain all the transactions
	TxList []transaction.Transaction
	// Length of the slice indicating the amounut of transactions currently in the pool
	// It's size can be limited by number of transaction
	// But practically it is limited by the block time i.e. the time period between forging of 2 successive blocks
	Length int
}

// This Method is used to create an empty transaction pool
// Later transaction can be added to the list and the pool can be modified
// The transactions are added using AddTransaction method
func New() TxPool {
	txPool := TxPool{
		TxList: []transaction.Transaction{},
		Length: 0,
	}

	return txPool
}

// Add transaction method
// Adds the transaction provided in its argument to the transaction pool list
func (txPool *TxPool) AddTransaction(tx transaction.Transaction) {

	if !txPool.isDuplicate(tx) { // This statement checks for any duplicates, duplicate transactions aren't added
		txPool.TxList = append(txPool.TxList, tx)
		txPool.Length = len(txPool.TxList)
	}
}

// This method helps in visualizing the transaction pool
//  This is actually used to view the current status of the transaction pool
// This will be further used in API
func (txPool TxPool) ToJson() string {
	jsonFormatted, _ := json.Marshal(txPool)
	return string(jsonFormatted)

}

// Checks for the duplicate transaction
// This method calls the Equals method in transaction package to establish the equality between 2 transactions
// It is a binary method
// Returns true if duplicate found in transaction pool from list of pre-existing transactions
// Otherwise Returns false
func (txPool TxPool) isDuplicate(tx transaction.Transaction) bool {
	for _, poolTx := range txPool.TxList {
		if poolTx.Equals(tx) {
			return true
		}
	}

	return false
}

// Removes the list of transactions from the transaction pool
// This function should be called everytime a new block is created
// So that the transactions are not added to the subsequent blocks
// And there is no double spending problem
func (txPool *TxPool) RemoveFromPool(txList []transaction.Transaction) {
	// If not tx in pool transactions then keep 
	// else if it is in then delete the transaction
	for i, tx := range txList {
		if txPool.contains(tx) {
			txPool.TxList = append(txPool.TxList[:i], txPool.TxList[i+1:]...)
		}
	}
	// Set the length
	txPool.Length = len(txPool.TxList)
}

// Helper function to check if a transaction is in transaction pool
func (txPool TxPool) contains(tx_ transaction.Transaction) bool {
	for _, tx := range txPool.TxList {
		if tx.Equals(tx_) {
			return true
		}
	}

	return false
}
