package orderdb

var GetOrderById = `SELECT o.id, o.user_id, o.shipping_address, o.shipping_city, o.shipping_postal_code, o.shipping_country, o.payment_method, o.payment_id, o.payment_status, o.payment_update_time, o.payment_email_address, o.items_price, o.tax_price, o.shipping_price, o.total_price, o.is_paid, o.paid_at, o.is_delivered, o.delivered_at, o.created_at, o.updated_at, u.id, u.name, u.email, u.created_at, u.updated_at FROM orders o JOIN users u ON o.user_id = u.id WHERE o.id = $1`

var GetByUserId = `SELECT o.id, o.user_id, o.shipping_address, o.shipping_city, o.shipping_postal_code, o.shipping_country, o.payment_method, o.payment_id, o.payment_status, o.payment_update_time, o.payment_email_address, o.items_price, o.tax_price, o.shipping_price, o.total_price, o.is_paid, o.paid_at, o.is_delivered, o.delivered_at, o.created_at, o.updated_at, u.id, u.name, u.email, u.created_at, u.updated_at FROM orders o JOIN users u ON o.user_id = u.id WHERE o.user_id = $1`

var GetOrders = `SELECT o.id, o.user_id, o.shipping_address, o.shipping_city, o.shipping_postal_code, o.shipping_country, o.payment_method, o.payment_id, o.payment_status, o.payment_update_time, o.payment_email_address, o.items_price, o.tax_price, o.shipping_price, o.total_price, o.is_paid, o.paid_at, o.is_delivered, o.delivered_at, o.created_at, o.updated_at, u.id, u.name, u.email, u.created_at, u.updated_at FROM orders o JOIN users u ON o.user_id = u.id`
var CreateOrder = `INSERT INTO orders (
	user_id,
	shipping_address,
	shipping_city,
	shipping_postal_code,
	shipping_country,
	payment_method,
	payment_id,
	payment_status,
	payment_update_time,
	payment_email_address,
	items_price,
	tax_price,
	shipping_price,
	total_price,
	is_paid,
	paid_at,
	is_delivered,
	delivered_at,
  created_at,
  updated_at
  ) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
  ) RETURNING
	id
  `

var CreateOrderItem = `INSERT INTO order_items (order_id, image, product_id, quantity, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

var GetOrderItems = `SELECT p.id, p.name, p.image, p.description, oi.price, oi.quantity FROM order_items oi LEFT JOIN products p ON oi.product_id = p.id WHERE order_id = $1`

var UpdateOrderToPaid = `
  UPDATE orders
  SET
    payment_id = $1,
    payment_status = $2,
    payment_update_time = $3,
    payment_email_address = $4,
    is_paid = $5,
    paid_at = $6,
    updated_at = $7
    WHERE id = $8
`
var UpdateOrderToDelivered = `
  UPDATE orders
  SET
    is_delivered = $1,
    delivered_at = $2,
    updated_at = $3
    WHERE id = $4
`

var UpdateOrder = `UPDATE orders
SET
  user_id = $2,
  shipping_address = $3,
  shipping_city = $4,
  shipping_postal_code = $5,
  shipping_country = $6,
  payment_method = $7,
  payment_id = $8,
  payment_status = $9,
  payment_update_time = $10,
  payment_email_address = $11,
  items_price = $12,
  tax_price = $13,
  shipping_price = $14,
  total_price = $15,
  is_paid = $16,
  paid_at = $17,
  is_delivered = $18,
  delivered_at = $19,
  updated_at = $20
WHERE id = $1
RETURNING
  id,
  user_id,
  shipping_address,
  shipping_city,
  shipping_postal_code,
  shipping_country,
  payment_method,
  payment_id,
  payment_status,
  payment_update_time,
  payment_email_address,
  items_price,
  tax_price,
  shipping_price,
  total_price,
  is_paid,
  paid_at,
  is_delivered,
  delivered_at,
  created_at,
  updated_at
`

var DeleteOrder = "DELETE FROM orders WHERE id = $1"
