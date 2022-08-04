-- name: GetProductForUpdate :one
SELECT * 
FROM products
WHERE sku = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateProduct :exec
UPDATE products
SET name = sqlc.arg(name)
WHERE id = sqlc.arg(id);

-- name: CreateOrUpdateProduct :one
INSERT INTO products (sku, name) 
VALUES ($1, $2)
ON CONFLICT (sku) DO UPDATE 
SET name = $2
RETURNING id;