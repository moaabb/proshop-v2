package productdb

var GetProductById = "SELECT id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at FROM products WHERE id = $1"
var GetProducts = "SELECT id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at FROM products LIMIT 10 OFFSET 10 * ($1 - 1)"
var GetTopProducts = "SELECT id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at FROM products ORDER BY rating DESC LIMIT 3"
var CreateProduct = "INSERT INTO products (name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at"
var GetReviewsByProduct = "SELECT r.id, r.rating, r.comment, r.product_id, r.created_at, r.updated_at, u.name, u.email FROM reviews AS r JOIN users AS u ON u.id = r.user_id WHERE r.product_id = $1"
var UpdateProduct = "UPDATE products SET name = $1, description = $2, brand = $3, category = $4, image = $5, num_reviews = $6, rating = $7, price = $8, count_in_stock = $9, updated_at = $10 WHERE id = $11 RETURNING id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at"
var DeleteProduct = "DELETE FROM products WHERE id = $1"
var GetTotalRecords = "SELECT COUNT(*) FROM products"
var CreateReview = `
WITH inserted_rows AS (INSERT INTO reviews
(rating, comment, user_id, product_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, rating, comment, product_id, user_id, created_at, updated_at )
SELECT r.id, r.rating, r.comment, r.product_id, r.created_at, r.updated_at, u.name, u.email FROM inserted_rows AS r JOIN users AS u ON r.user_id = u.id`
var UpdateRating = "UPDATE products SET rating = $1, num_reviews = $2 WHERE id = $3"
