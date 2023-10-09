package domain

const (
	insertKonsumen = `INSERT INTO konsumen (
		nik,
		full_name,
		legal_name,
		tempat_lahir,
		tanggal_lahir,
		gaji,
		foto_ktp,
		foto_selfie,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	readKonsumenByID = `SELECT
		nik,
		full_name,
		legal_name,
		tempat_lahir,
		tanggal_lahir,
		gaji,
		foto_ktp,
		foto_selfie,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM konsumen WHERE nik = ?`
)
