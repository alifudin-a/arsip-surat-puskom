package query

var ListAllPenerima = `SELECT * FROM tbl_penerima ORDER BY id;`

var ReadPenerimaById = `SELECT * FROM tbl_penerima WHERE id = $1;`

var DeletePenerima = `DELETE FROM tbl_penerima WHERE id= $1;`

var CreatePenerima = `INSERT INTO tbl_penerima (id_surat, id_pengguna, created_at) VALUES ($1, $2, $3) RETURNING *;`

var UpdatePenerima = `UPDATE tbl_penerima SET id_surat = $1, id_pengguna = $2, created_at = $3 WHERE id = $4 RETURNING *;`

var IsExistPenerima = `SELECT COUNT(*) FROM tbl_penerima WHERE id = $1;`
