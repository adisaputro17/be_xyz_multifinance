package domain

const (
	insertTransaksi = `INSERT INTO transaksi (
		nomor_kontrak,
		konsumen_id,
		nama_asset,
		otr,
		admin_fee,
		jumlah_cicilan,
		jumlah_bunga,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	readTransaksiByKonsumenID = `SELECT 
		nomor_kontrak,
		konsumen_id,
		nama_asset,
		otr,
		admin_fee,
		jumlah_cicilan,
		jumlah_bunga,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM transaksi WHERE konsumen_id = ?`
)
