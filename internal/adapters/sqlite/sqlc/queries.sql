-- name: ListProducts :many
SELECT * FROM products;

-- name: GetProductByID :one
SELECT * FROM products WHERE id = ?;

-- name: CreateOrder :one
INSERT INTO orders (
    customer_id
) VALUES (?) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    price_in_cents
) VALUES (?, ?, ?, ?) RETURNING *;

-- name: UpdateProductQuantity :exec
UPDATE products
SET quantity = ?
WHERE id = ?;
