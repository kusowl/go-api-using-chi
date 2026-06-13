-- name: ListProducts :many
SELECT * from products;

-- name: FindProductById :one
SELECT * from products WHERE id = $1;


