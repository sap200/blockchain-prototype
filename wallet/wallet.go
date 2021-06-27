package wallet

// Wallet package is one of the fundamental package
// It is the class that generates address for the Accounut Model in the blockchain
// This package is used for signing the transaction and handles the private and public key of the user
// It also has some other helper methods to verify the transaction using signatures and
// Also returns the public key which is the public Address of the wallet in string hexadecimal string format

import (
	"blockchain/transaction"
	"blockchain/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

// The wallet structure
// The keypair of the user, private key is named because it is a struct and it has an attribute public key
// That can be used to get public key from private key
// This is basically the wallet address
type Wallet struct {
	PrivateKey *rsa.PrivateKey
}

// This Method is used to createNew Wallet
// It calls the GenerateKey Method from the rsa package
// rand.Reader is sequence of random bytes to make the private key tough and
// 2048 is the modulo
func New() Wallet {

	// Generate the RSA public private keypair
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	wallet := Wallet{
		PrivateKey: privateKey,
	}

	return wallet
}

// This method returns the public key in PEM format
// It has several steps that are
// First retrieve the public key from the instance of the private key
// Then convert the key into bytes format usinf x059 package's Marshal method
// Then create the PEM block wih a message
// This Message marks the beginning and ending of the public key
// Usually the message is RSA PUBLIC KEY and hence i have used it
// Then encode the block into PEM format and return the string format
func (wallet Wallet) PublicKeyString() string {
	// Retrieve the private key and get the public key from it
	privateKey := wallet.PrivateKey
	publicKey := &privateKey.PublicKey

	// Convert the public key into bytes
	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKey)

	// Make the public key block for PEM encoding
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// Encode the key
	publicKeyString := pem.EncodeToMemory(publicKeyBlock)
	return string(publicKeyString)
}

// This method is used to sign the transaction
// This method first hashes the data
// And then uses a func from rsa package to sign in PKCS1v15 format
// The signature is of type []byte. i.e. slice of byte
// And hence for brief representation is converted to hexadecimal format using hex encoding
func (wallet Wallet) Sign(data string) string {
	dataHash := utils.Hash(data)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, wallet.PrivateKey, crypto.SHA256, dataHash)
	signatureHex := hex.EncodeToString(signature)
	return signatureHex
}

// This method is used to create a signed transaction
// The transaction will be signed by the owner of the wallet
// And hence no senderPublicKey in the argument
func (wallet Wallet) NewSignedTransaction(receiverPublicKey string, amount float64, txType transaction.TransactionType) transaction.Transaction {
	tx := transaction.New(wallet.PublicKeyString(), receiverPublicKey, amount, txType)
	signature := wallet.Sign(tx.Payload())
	tx.Sign(signature)
	return tx
}

// This method is used to check whether the transaction is a valid transaction or not
// it does so by verifying the signature
// The dataHash and the public key is used to recreate the signature and hene verification happens
// using the verify method of rsa package
// This is PKCS1v15 signature, there are other kind of signatures too
func CheckSignature(data string, publicKeyString string, signatureString string) bool {
	// convert the publickeystring back into rsa public key
	rsaPublicKey := importPubKey(publicKeyString)
	dataHash := utils.Hash(data)
	signatureInByte, _ := hex.DecodeString(signatureString)
	err := rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, dataHash, signatureInByte)
	if err == nil {
		return true
	} else {
		return false
	}
}

// This is a helper method and it is used to
// reconstruct the publick key in *rsa.PublicKey format from the string
// This isa reverse process
// The string is converted into bytes and pem.Decode returns the PEM block
// The block is parsed usinf x509 package's parse method
// This returns rsa Public key and if any error, which I ignored here
// As this methods wont be using any invalid keys as all keys are system generated.
func importPubKey(pubKeyStr string) *rsa.PublicKey {
	pemInByte := []byte(pubKeyStr)
	block, _ := pem.Decode(pemInByte)
	rsaPubKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	return rsaPubKey.(*rsa.PublicKey)
}
