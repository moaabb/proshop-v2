package userdb

var GetUserByEmail = "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users WHERE email = $1"
