package db

import (
	"context"
	"database/sql"
	"fmt"
)

//provides all functions to execute queries & db transactions run individually
type Store struct {
   *Queries //extending the queries struct on the store struct composition instead of inheritance
   db *sql.DB 
}

//NewStore creates 
func NewStore(db *sql.DB) *Store{
	return &Store{
		db:db,
		Queries: New(db),
	}
}
//execTx executes a function  within adatabase transaction 
func(store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
	tx, err:= store.db.BeginTx(ctx,nil)
   if err!=nil{
	   return err
   }
   q := New(tx)
   err = fn(q)
   if err!=nil{
	   if rbErr := tx.Rollback(); rbErr != nil{
         return fmt.Errorf("tx err: %v ,rb err: %v", err, rbErr)
	   }
	   return err
   }
   return tx.Commit()
}

//transfertxparams contains input parameters of the transfer transaction
type TransferTxParams struct{
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
//TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer Transfer 	 `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}
//TransferTx performs a money transfer from one account to the other
//it creates a transfer record, add account entries and updates accounts balance within asingle transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams)(TransferTxResult, error){
   var result TransferTxResult

   err := store.execTx(ctx ,func(q *Queries) error {
	    var err error
	    result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount:arg.Amount,
		})
		if err!=nil{
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err!=nil{
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err!=nil{
			return err
		}
		//TODO update balance transaction
	   return nil
   })
   return result, err
}