-- name: GetOneProduct :one
SELECT * FROM products
WHERE id = ? LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: CreateProduct :execresult
INSERT INTO products (
  title, description
) VALUES (
  ?, ?
);

-- name: UpdateProduct :exec
UPDATE products
SET title = ?, description = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;