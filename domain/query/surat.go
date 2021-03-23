package query

var ListSurat = `SELECT * FROM tbl_surat OREDER BY id;`

var ReadSuratByID = `SELECT * FROM tbl_surat WHERE id = $1;`

var DeleteSurat = `DELETE FROM tbl_surat WHERE id= $1;`

var CreateSurat = `
INSERT INTO
tbl_surat (
	tanggal,
	nomor,
	id_penerima,
	id_pengirim,
	perihal,
	id_jenis,
	keterangan,
	created_at
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;`

var UpdateSurat = `
UPDATE 
	tbl_surat
SET 
	tanggal = $1, 
	nomor = $2, 
	id_penerima = $3, 
	id_pengirim = $4, 
	perihal = $5,
	id_jenis = $6,
	keterangan = $7,
	updated_at = $8
WHERE
	id = $9 
RETURNING *;`

var IsExistSurat = `SELECT COUNT(*) FROM tbl_surat WHERE id= $1;`
