package query

var ListAllJabatanStruktural = `SELECT * FROM jabatan_struktural ORDER BY id;`

var GetJabatanStrukturalByID = `SELECT * FROM jabatan_struktural WHERE id = $1;`
