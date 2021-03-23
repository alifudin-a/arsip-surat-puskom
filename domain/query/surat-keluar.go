package query

var ListAllSuratKeluar = `
select
	tsk.id ,
	tsk.tanggal,
	tsk.nomor,
	tp1."name" as pengirim,
	tsk.nomor,
	tsk.perihal,
	tjs."name" as jenis,
	tsk.keterangan
from
	tbl_surat_keluar tsk
join tbl_pengguna tp1 on tp1.id = tsk.id_pengirim
left join tbl_jenis_surat tjs on	tjs.id = tsk.id_jenis;`

var ReadSuratKeluarByID = `
select
	tsk.id ,
	tsk.tanggal,
	tsk.nomor,
	tp1."name" as pengirim,
	tsk.nomor,
	tsk.perihal,
	tjs."name" as jenis,
	tsk.keterangan
from
	tbl_surat_keluar tsk
join tbl_pengguna tp1 on tp1.id = tsk.id_pengirim
left join tbl_jenis_surat tjs on	tjs.id = tsk.id_jenis
where tsk.id = $1;`

var DeletSuratKeluar = `DELETE FROM tbl_surat_keluar WHERE id =$1;`

var UpdateSuratKeluar = `
UPDATE 
	tbl_surat_keluar 
SET 
	tanggal = $1, 
	nomor = $2, 
	id_pengirim = $3, 
	perihal = $4,
	id_jenis = $5,
	keterangan = $6,
	updated_at = $7
WHERE
	id = $8 
RETURNING *;`

var CreateSuratKeluar = `
INSERT INTO
tbl_surat_keluar (
	tanggal,
	nomor,
	id_pengirim,
	perihal,
	keterangan,
	created_at
) VALUES (
	$1, $2, $3, $4, $5, $6
) RETURNING *;`

var IsExistSuratKeluar = `SELECT COUNT(*) FROM tbl_surat_keluar WHERE id = $1;`

var SuratKeluarJoin = `
select
	tsk.id,
	tsk.tanggal,
	tsk.nomor,
	tsk.perihal,
	tp1."name" as pengirim,
	tsk.keterangan,
	tp.id,
	tp.id_surat,
	tp2."name" as penerima
from
	tbl_surat_keluar tsk
join tbl_pengguna tp1 on tp1.id = tsk.id_pengirim
join tbl_penerima tp on tsk.id = tp.id_surat
join tbl_pengguna tp2 on tp.id_pengguna = tp2.id;
`
