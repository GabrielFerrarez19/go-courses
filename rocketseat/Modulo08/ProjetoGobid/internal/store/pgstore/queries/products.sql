-- name: CreateProduct :one
INSERT INTO products (
        seller_id,
        product_name,
        description,
        baseprice,
        auction_end
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
-- name: GetProduct :many
SELECT *
FROM products;
-- name: GetProductById :one
SELECT *
FROM products
WHERE id = $1;