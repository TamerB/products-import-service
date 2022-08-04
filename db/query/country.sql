-- name: GetCountryForUpdate :one
SELECT * 
FROM countries
WHERE country_code = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateCountry :exec
UPDATE countries
SET country_code = sqlc.arg(country_code)
WHERE id = sqlc.arg(id);

-- name: CreateCountry :one
INSERT INTO countries (country_code) 
VALUES ($1)
RETURNING id;

-- name: CreateOrUpdateCountry :one
INSERT INTO countries (country_code) 
VALUES ($1)
ON CONFLICT (country_code) DO UPDATE 
SET country_code = $1
RETURNING id;