// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: stock.sql

package db

import (
	"context"
)

const createOrUpdateStock = `-- name: CreateOrUpdateStock :one
INSERT INTO stocks (product_id, country_id, quantity) 
VALUES ($1, $2, $3)
ON CONFLICT (product_id, country_id) DO UPDATE 
SET quantity = (SELECT quantity from stocks WHERE product_id = $1 and country_id = $2 LIMIT 1 FOR NO KEY UPDATE) + 1
RETURNING id
`

type CreateOrUpdateStockParams struct {
	ProductID int64 `json:"product_id"`
	CountryID int64 `json:"country_id"`
	Quantity  int64 `json:"quantity"`
}

func (q *Queries) CreateOrUpdateStock(ctx context.Context, arg CreateOrUpdateStockParams) (int64, error) {
	row := q.queryRow(ctx, q.createOrUpdateStockStmt, createOrUpdateStock, arg.ProductID, arg.CountryID, arg.Quantity)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createStock = `-- name: CreateStock :one
INSERT INTO stocks (
  product_id,
  country_id,
  quantity
) VALUES (
    $1, $2, $3
) RETURNING stocks.id, stocks.quantity
`

type CreateStockParams struct {
	ProductID int64 `json:"product_id"`
	CountryID int64 `json:"country_id"`
	Quantity  int64 `json:"quantity"`
}

type CreateStockRow struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

func (q *Queries) CreateStock(ctx context.Context, arg CreateStockParams) (CreateStockRow, error) {
	row := q.queryRow(ctx, q.createStockStmt, createStock, arg.ProductID, arg.CountryID, arg.Quantity)
	var i CreateStockRow
	err := row.Scan(&i.ID, &i.Quantity)
	return i, err
}

const getStockByProductAndCountryeForUpdate = `-- name: GetStockByProductAndCountryeForUpdate :one
SELECT stocks.id, stocks.quantity 
FROM stocks
INNER JOIN products
ON stocks.product_id = products.id
AND products.id = $1
INNER JOIN countries
ON stocks.country_id = countries.id
AND countries.id = $2
LIMIT 1
FOR NO KEY UPDATE
`

type GetStockByProductAndCountryeForUpdateParams struct {
	ProductID int64 `json:"product_id"`
	CountryID int64 `json:"country_id"`
}

type GetStockByProductAndCountryeForUpdateRow struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

func (q *Queries) GetStockByProductAndCountryeForUpdate(ctx context.Context, arg GetStockByProductAndCountryeForUpdateParams) (GetStockByProductAndCountryeForUpdateRow, error) {
	row := q.queryRow(ctx, q.getStockByProductAndCountryeForUpdateStmt, getStockByProductAndCountryeForUpdate, arg.ProductID, arg.CountryID)
	var i GetStockByProductAndCountryeForUpdateRow
	err := row.Scan(&i.ID, &i.Quantity)
	return i, err
}

const updateStock = `-- name: UpdateStock :exec
UPDATE stocks
SET quantity = $1
WHERE id = $2
`

type UpdateStockParams struct {
	Quantity int64 `json:"quantity"`
	ID       int64 `json:"id"`
}

func (q *Queries) UpdateStock(ctx context.Context, arg UpdateStockParams) error {
	_, err := q.exec(ctx, q.updateStockStmt, updateStock, arg.Quantity, arg.ID)
	return err
}
