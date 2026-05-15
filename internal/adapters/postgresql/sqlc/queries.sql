-- name: ListProducts :many
SELECT * FROM products;

-- name: ProductById :one
SELECT * FROM products
    WHERE id = $1;

-- name: AddProduct :exec
INSERT INTO products (name, price, quantity) 
VALUES ($1, $2, $3);

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $2, price = $3, quantity = $4
WHERE id = $1;

-- name: UpdateProductQuantity :exec
UPDATE products
SET quantity = $2
WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO "order_new" (customer_id, created_at, updated_at)
VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO "order_items_new" (order_id, product_id, quantity, price, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: AddOrderItem :one
INSERT INTO "order_items_new" (order_id, product_id, quantity, price, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: GetOrderById :one
SELECT o.id, o.customer_id, o.created_at, o.updated_at,
       oi.id AS order_item_id, oi.product_id, oi.quantity, oi.price, oi.created_at AS order_item_created_at, oi.updated_at AS order_item_updated_at
FROM "order_new" o
LEFT JOIN "order_items_new" oi ON o.id = oi.order_id
WHERE o.id = $1;


-- name: ListOrders :many
SELECT o.id, o.customer_id, o.created_at, o.updated_at,
       oi.id AS order_item_id, oi.product_id, oi.quantity, oi.price, oi.created_at AS order_item_created_at, oi.updated_at AS order_item_updated_at
FROM "order_new" o
LEFT JOIN "order_items_new" oi ON o.id = oi.order_id;
