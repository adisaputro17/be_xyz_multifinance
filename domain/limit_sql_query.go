package domain

const (
	insertLimit = `INSERT INTO limit_konsumen (
		limit_id,
		konsumen_id,
		tenor,
		limit_pinjam,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	readLimitByKonsumenID = `SELECT
		limit_id,
		konsumen_id,
		tenor,
		limit_pinjam,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM limit_konsumen WHERE konsumen_id = ?`

	updateLimitByLimitID = `UPDATE limit_konsumen SET
	limit_pinjam = ?,
		updated_at = ?,
		updated_by = ?
	WHERE limit_id = ?`
)
