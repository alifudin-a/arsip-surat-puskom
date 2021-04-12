package query

var ListSuratKeluar = `
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

var ListSuratKeluarAsc = `
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
	ORDER BY id ASC OFFSET $1
)t;`

var ListSuratKeluarByIDPengirim = `
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
	WHERE ts.id_pengirim = $1 ORDER BY id DESC
)t;`

var ReadSuratKeluar = `
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
	WHERE ts.id = $1
)t;`
