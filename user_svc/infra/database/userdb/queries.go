package userdb

var GetUserById = "SELECT id, name, email, is_admin, created_at, updated_at FROM users WHERE id = $1"
var GetUsers = "SELECT id, name, email, is_admin, created_at, updated_at FROM users"
var CreateUser = "INSERT INTO users (name, email, password, is_admin, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, password, is_admin, created_at, updated_at"
var UpdateUser = "UPDATE users SET name = $1, email = $2, password = $3, is_admin = $4, updated_at = $5 WHERE id = $6 RETURNING id, name, email, password, is_admin, created_at, updated_at"
var DeleteUser = "DELETE FROM users WHERE id = $1"
