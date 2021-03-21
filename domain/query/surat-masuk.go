package query

var ListAllSuratMasuk = `
select
	tsm.id ,
	tsm.tanggal,
	tsm.nomor,
	tp1."name" as pengirim,
	tp2."name" as penerima,
	tsm.nomor,
	tsm.perihal,
	tjs."name" as jenis,
	tsm.keterangan
from
	tbl_surat_masuk tsm
join tbl_pengguna tp1 on tp1.id = tsm.id_pengirim
join tbl_pengguna tp2 on tp2.id = tsm.id_penerima 
left join tbl_jenis_surat tjs on tjs.id = tsm.id_jenis;`

var ReadSuratMasukByID = `
select
	tsm.id ,
	tsm.tanggal,
	tsm.nomor,
	tp1."name" as pengirim,
	tp2."name" as penerima,
	tsm.nomor,
	tsm.perihal,
	tjs."name" as jenis,
	tsm.keterangan
from
	tbl_surat_masuk tsm
join tbl_pengguna tp1 on tp1.id = tsm.id_pengirim
join tbl_pengguna tp2 on tp2.id = tsm.id_penerima 
left join tbl_jenis_surat tjs on	tjs.id = tsm.id_jenis
where tsm.id = $1;`

var DeletSuratMasuk = `DELETE FROM tbl_surat_masuk WHERE id =$1;`

var UpdateSuratMasuk = `
UPDATE 
	tbl_surat_masuk 
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

var CreateSuratMasuk = `
INSERT INTO
tbl_surat_masuk (
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

var IsExistSuratMasuk = `SELECT COUNT(*) FROM tbl_surat_masuk WHERE id = $1;`
