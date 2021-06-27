package bank

import (
	"encoding/json"
)

type Bank struct {
	Accounts []string
	Balances map[string]float64
}

func New() Bank {
	bank := Bank{
		Accounts: []string{},
		Balances: map[string]float64{},
	}

	return bank
}

func (b *Bank) AddAccount(accountAddress string) {
	if !contains(b.Accounts, accountAddress) {
		b.Accounts = append(b.Accounts, accountAddress)
		b.Balances[accountAddress] = 0
	}
}

func (b *Bank) UpdateBalance(accountAddress string, amount float64) {
	if !contains(b.Accounts, accountAddress) {
		b.AddAccount(accountAddress)
	}

	b.Balances[accountAddress] += amount
}

func (b *Bank) GetBalance(accountAddress string) float64 {
	if !contains(b.Accounts, accountAddress) {
		b.AddAccount(accountAddress)
	}

	return b.Balances[accountAddress]
}

func contains(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}

	return false
}

func (b Bank) ToJson() string {
	bankInByteSlice,_ := json.Marshal(b)
	return string(bankInByteSlice)
}
