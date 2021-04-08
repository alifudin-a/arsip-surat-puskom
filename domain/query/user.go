package query

var ListAllUser = `SELECT * FROM tbl_pengguna ORDER BY id DESC;`

var ReadUserByID = `SELECT * FROM tbl_pengguna WHERE id = $1;`

var DeleteUser = `DELETE FROM tbl_pengguna WHERE id = $1;`

var IsExistUser = `SELECT COUNT(*) FROM tbl_pengguna WHERE id = $1;`

var CreateUser = `INSERT INTO tbl_pengguna (name, fullname, created_at) VALUES ($1, $2, $3) RETURNING *;`

var UpdateUser = `UPDATE tbl_pengguna SET name= $1, fullname = $2, updated_at = $3 WHERE id = $4 RETURNING *;`
