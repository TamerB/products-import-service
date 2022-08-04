package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

var DBStore *SQLStore

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateOrUpdateStockTx(ctx context.Context, arg UpdateStockTxParams) error
}

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function withing a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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

// UpdateStockTxParams contains the input parameter of the stock update transaction
type UpdateStockTxParams struct {
	ProductSKU  string `json:"product_sku"`
	ProductName string `json:"product_name"`
	CountryCode string `json:"country_code"`
	Quantity    int64  `json:"quantity"`
}

// UpdateStockTx performs a stock update.
// It validates that the stock exists and will not be negative after the update.
func (store *SQLStore) CreateOrUpdateStockTx(ctx context.Context, arg UpdateStockTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		productID, err := q.CreateOrUpdateProduct(ctx, CreateOrUpdateProductParams{
			Sku:  arg.ProductSKU,
			Name: arg.ProductName,
		})

		if err != nil {
			return err
		}

		countryID, err := q.CreateOrUpdateCountry(ctx, arg.CountryCode)

		if err != nil {
			if !strings.Contains(err.Error(), "duplicate") {
				return err
			}
		}

		_, err = q.CreateOrUpdateStock(ctx, CreateOrUpdateStockParams{
			ProductID: productID,
			CountryID: countryID,
			Quantity:  arg.Quantity,
		})

		if err != nil {
			if strings.Contains(err.Error(), "violates check constraint \"stocks_quantity_check\"") {
				return fmt.Errorf("quantity value would make quantity in minus: %d", arg.Quantity)
			}
			return err
		}

		return nil
	})

	return err
}
