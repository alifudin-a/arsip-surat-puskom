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

var DeleteSuratMasuk = `DELETE FROM tbl_surat WHERE id = $1;`

var IsSuratMasukExist = `SELECT COUNT(*) FROM tbl_surat WHERE id= $1;`
