package query

var ListAllJabatanStruktural = `SELECT * FROM jabatan_struktural ORDER BY id;`

var ReadJabatanStrukturalByID = `SELECT * FROM jabatan_struktural WHERE id = $1;`

var DeleteJabatanStruktural = `DELETE FROM jabatan_struktural WHERE id = $1;`

var IsExistJabatanStruktural = `SELECT COUNT(*) FROM jabatan_struktural WHERE id = $1;`

var CreateJabatanStruktural = `INSERT INTO jabatan_struktural (name, fullname) VALUES ($1, $2) RETURNING *;`

var UpdateJabatanStruktural = `UPDATE jabatan_struktural SET name= $1, fullname = $2 WHERE id = $3 RETURNING *;`
