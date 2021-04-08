package query

var ListAllJenisSurat = `SELECT * FROM tbl_jenis_surat ORDER BY id;`

var ReadJenisSuratById = `SELECT * FROM tbl_jenis_surat WHERE id = $1;`

var DeleteJenisSurat = `DELETE FROM tbl_jenis_surat WHERE id = $1;`

var IsExistJenisSurat = `SELECT COUNT(*) FROM tbl_jenis_surat WHERE id = $1;`

var CreateJenisSurat = `INSERT INTO tbl_jenis_surat (name, created_at) VALUES ($1, $2, $3) RETURNING *;`

var UpdateJenisSurat = `UPDATE tbl_jenis_surat SET name = $1, updated_at = $2 WHERE id = $3 RETURNING *;`
