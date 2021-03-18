package query

var ListAllJabatanStruktural = `SELECT * FROM jabatan_struktural ORDER BY id;`

var ReadJabatanStrukturalByID = `SELECT * FROM jabatan_struktural WHERE id = $1;`

var DeleteJabatanStruktural = `DELETE FROM jabatan_struktural WHERE id = $1;`

var IsExistJabatanStruktural = `SELECT COUNT(*) FROM jabatan_struktural WHERE id = $1;`
