package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions.
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}


// execTx executes a function within a database transaction.
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams is the input parameters for TransferTx
type TransferTxParams struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccount_ID   int64  `json:"to_account_id"`
	Amount int64  `json:"amount"`
}

// TransferTxResult is the output parameters for TransferTx
type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount   Account `json:"to_account"`
	FromEntry   Entry `json:"from_entry"`
	ToEntry   Entry `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another.
// It creates a transfer record, and account entries, and update account balances within a single database transaction


