-- name: ListProducts :many
SELECT * FROM products;

-- name: ProductById :one
SELECT * FROM products
    WHERE id = $1;