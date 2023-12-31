package core

import (
	"ledger/common"
	"math/big"

	"github.com/rs/xid"
)

type Transaction struct {
	id      string
	entries []Entries
}

type Entries struct {
	id        string
	Account   *Account
	Amount    *big.Int
	Direction common.Direction
}

// newEntries creates a new Entries with a unique id.
func NewEntry(Account *Account, amount *big.Int, direction common.Direction, status common.Status) *Entries {
	return &Entries{
		id:        xid.New().String(),
		Account:   Account,
		Amount:    amount,
		Direction: direction,
	}
}

// NewTransaction creates a new Transaction with a unique id and a slice of entries.
func NewTransaction(entries ...Entries) *Transaction {
	return &Transaction{
		id:      xid.New().String(),
		entries: entries,
	}
}

//convert transactions into double ledger Transaction
//

// type Transaction struct {
// 	signer Account //TODO: change to hash
// 	to     Account //TODO: change to hash
// 	value  *big.Int
// 	status string
// }

// type Entry struct {
// 	account Account
// 	value   *big.Int
// }

// type TxData interface {
// 	txType() byte
// }
