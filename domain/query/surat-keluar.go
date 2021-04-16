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

var ListSuratKeluarByIDPengirimAsc = `
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
	WHERE ts.id_pengirim = $1 ORDER BY id ASC OFFSET $2
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

var DeleteSuratKeluar = `DELETE FROM tbl_surat WHERE id = $1;`

var DeletePenerimaSuratKeluar = `DELETE FROM tbl_penerima WHERE id_surat = $1;`

var IsPenerimaSuratKeluarExist = `SELECT COUNT(*) FROM tbl_penerima WHERE id_surat = $1;`

var IsSuratKeluarExist = `SELECT COUNT(*) FROM tbl_surat WHERE id= $1;`

var CreateSuratKeluar = `
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

var UpdateSuratKeluar = `
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

var CreatePenerimaSuratKeluar = `
	INSERT
		INTO
		tbl_penerima (
			id_surat, id_pengguna, created_at
		) 
		VALUES
	`

var UpdatePenerimaSuratKeluar = `
	INSERT
		INTO
		tbl_penerima (
			id_surat, id_pengguna, created_at, updated_at
		) 
		VALUES
	`
