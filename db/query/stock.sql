-- name: GetStockByProductAndCountryeForUpdate :one
SELECT stocks.id, stocks.quantity 
FROM stocks
INNER JOIN products
ON stocks.product_id = products.id
AND products.id = sqlc.arg(product_id)
INNER JOIN countries
ON stocks.country_id = countries.id
AND countries.id = sqlc.arg(country_id)
LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateStock :exec
UPDATE stocks
SET quantity = sqlc.arg(quantity)
WHERE id = sqlc.arg(id);

-- name: CreateStock :one
INSERT INTO stocks (
  product_id,
  country_id,
  quantity
) VALUES (
    $1, $2, $3
) RETURNING stocks.id, stocks.quantity;

-- name: CreateOrUpdateStock :one
INSERT INTO stocks (product_id, country_id, quantity) 
VALUES ($1, $2, $3)
ON CONFLICT (product_id, country_id) DO UPDATE 
SET quantity = (SELECT quantity from stocks WHERE product_id = $1 and country_id = $2 LIMIT 1 FOR NO KEY UPDATE) + 1
RETURNING id;