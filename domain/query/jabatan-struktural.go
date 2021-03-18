package query

var ListAllJabatanStruktural = `SELECT * FROM jabatan_struktural ORDER BY id;`

var ReadJabatanStrukturalByID = `SELECT * FROM jabatan_struktural WHERE id = $1;`
