package query

var ListSuratMasukDesc = `
select 
	ts.id, 
	ts.tanggal, 
	ts.nomor, 
	ts.id_pengirim, 
	tp."name" as pengirim,
	ts.id_jenis,
	tjs."name" as jenis,
	ts.perihal, 
	ts.keterangan 
from 
	tbl_surat ts
left join tbl_pengguna tp on tp.id = ts.id_pengirim
left join tbl_jenis_surat tjs on tjs.id = ts.id_jenis 
order by id desc;`

var ListSuratMasukAsc = `
select 
	ts.id, 
	ts.tanggal, 
	ts.nomor, 
	ts.id_pengirim, 
	tp."name" as pengirim,
	ts.id_jenis,
	tjs."name" as jenis,
	ts.perihal, 
	ts.keterangan 
from 
	tbl_surat ts
left join tbl_pengguna tp on tp.id = ts.id_pengirim
left join tbl_jenis_surat tjs on tjs.id = ts.id_jenis 
order by id asc offset $1;`

var GetSuratMasukByID = `
select 
	ts.id, 
	ts.tanggal, 
	ts.nomor, 
	ts.id_pengirim, 
	tp."name" as pengirim,
	ts.id_jenis,
	tjs."name" as jenis,
	ts.perihal, 
	ts.keterangan 
from 
	tbl_surat ts
left join tbl_pengguna tp on tp.id = ts.id_pengirim
left join tbl_jenis_surat tjs on tjs.id = ts.id_jenis
where ts.id = $1;`

var DeletePenerimaSurat = `DELETE FROM tbl_penerima WHERE id_surat = $1;`

var DeleteSuratMasuk = `DELETE FROM tbl_surat WHERE id = $1;`

var IsPenerimaSuratExist = `SELECT COUNT(*) FROM tbl_penerima WHERE id_surat = $1;`

var IsSuratMasukExist = `SELECT COUNT(*) FROM tbl_surat WHERE id= $1;`

var ListSuratMasukByIDPenerima = `
select 
	ts.id, 
	ts.tanggal, 
	ts.nomor, 
	ts.id_pengirim, 
	tp."name" as pengirim,
	ts.id_jenis,
	tjs."name" as jenis,
	ts.perihal, 
	ts.keterangan
from 
	tbl_surat ts
left join tbl_pengguna tp on tp.id = ts.id_pengirim
left join tbl_jenis_surat tjs on tjs.id = ts.id_jenis
left join tbl_penerima tp2 on tp2.id_surat = ts.id
where tp2.id_pengguna = $1 order by id DESC;`

var CreateSuratMasuk = `
	insert 
		into 
			tbl_surat (
				tanggal, 
				nomor, 
				id_pengirim, 
				perihal, 
				id_jenis, 
				keterangan, 
				created_at
			) values(
				$1,$2,$3,$4,$5,$6,$7
			) RETURNING *;`

var UpdateSuratMasuk = `
UPDATE 
	tbl_surat
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
