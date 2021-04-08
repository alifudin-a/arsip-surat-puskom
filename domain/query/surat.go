package query

var ReadSuratByID = `
select row_to_json(t)
from (
	select
		ts.id,
		ts.tanggal,
		ts.nomor,
		ts.id_pengirim,
		tp2."name" as pengirim,
		ts.perihal,
		ts.id_jenis,
		tjs."name" as jenis,
		ts.keterangan,
		ts.created_at,
		ts.updated_at,
		(
			select array_to_json(array_agg(row_to_json(d)))
			from (
				select 
					tp.id as id,
					tp.id_pengguna,
					tp3."name" as penerima,
					tp.id_surat,
					tp.created_at,
					tp.updated_at
				from tbl_penerima tp
				join tbl_pengguna tp3 on tp3.id = tp.id_pengguna 
				where tp.id_surat=ts.id
			)d
		)as penerima
	from tbl_surat ts
	join tbl_pengguna tp2 on tp2.id = ts.id_pengirim
	left join tbl_jenis_surat tjs on tjs.id = ts.id_jenis where ts.id = $1
)t;`

var DeleteSurat = `DELETE FROM tbl_surat WHERE id = $1;`

var CreateSurat = `
	INSERT
		INTO
		tbl_surat ( 
			tanggal,
			nomor,
			id_pengirim,
			perihal,
			id_jenis ,
			keterangan ,
			created_at 
		)
		VALUES ( $1, $2, $3, $4, $5, $6, $7 ) RETURNING *;`

var CreatePenerimaSurat = `
	INSERT
		INTO
		tbl_penerima (
			id_surat, id_pengguna, created_at
		) 
		VALUES
	`

var UpdatePenerimaSurat = `
	INSERT
		INTO
		tbl_penerima (
			id_surat, id_pengguna, created_at, updated_at
		) 
		VALUES
	`

var UpdateSurat = `
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

var IsExistSurat = `SELECT COUNT(*) FROM tbl_surat WHERE id= $1;`

var ListSurat = `
SELECT row_to_json(t)
FROM (
	SELECT
		ts.id,
		ts.tanggal,
		ts.nomor,
		tp2."name" AS pengirim,
		ts.perihal,
		tjs."name" AS jenis,
		ts.keterangan,
		(
			SELECT array_to_json(array_agg(row_to_json(d)))
			FROM (
				SELECT 
					tp.id as id,
					tp3."name" as name
				FROM tbl_penerima tp
				JOIN tbl_pengguna tp3 on tp3.id = tp.id_pengguna 
				WHERE tp.id_surat=ts.id
			)d
		)AS penerima
	FROM tbl_surat ts
	JOIN tbl_pengguna tp2 on tp2.id = ts.id_pengirim
	LEFT JOIN tbl_jenis_surat tjs on tjs.id = ts.id_jenis
	ORDER BY id DESC
)t;`
