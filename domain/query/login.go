package query

var GetUsername = `SELECT * FROM tbl_login WHERE username = $1 AND password = $2;`
