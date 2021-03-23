package query

var Login = `SELECT * FROM tbl_login WHERE username = $1 AND password = $2;`
