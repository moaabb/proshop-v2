package productdb

var GetProductById = "SELECT id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at FROM products WHERE id = $1"
var GetProducts = "SELECT id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at FROM products LIMIT 10 OFFSET 10 * ($1 - 1)"
var GetTopProducts = "SELECT id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at FROM products ORDER BY rating DESC LIMIT 3"
var CreateProduct = "INSERT INTO products (name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at"
var GetReviewsByProduct = "SELECT id, rating, comment, user_id, product_id, created_at, updated_at FROM reviews WHERE id = $1"
var UpdateProduct = "UPDATE products SET name = $1, description = $2, brand = $3, category = $4, image = $5, num_reviews = $6, rating = $7, price = $8, count_in_stock = $9, updated_at = $10 WHERE id = $11 RETURNING id, name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at"
var DeleteProduct = "DELETE FROM products WHERE id = $1"
var GetTotalRecords = "SELECT COUNT(*) FROM products"
