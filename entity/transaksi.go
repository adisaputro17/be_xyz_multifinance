package entity

type Transaksi struct {
	NomorKontrak  string  `json:"nomor_kontrak"`
	KonsumenID    string  `json:"konsumen_id"`
	NamaAsset     string  `json:"nama_asset"`
	OTR           float64 `json:"otr"`
	AdminFee      float64 `json:"admin_fee"`
	JumlahCicilan float64 `json:"jumlah_cicilan"`
	JumlahBunga   float64 `json:"jumlah_bunga"`
	CommonField
}

type InsertTransaksiRequest struct {
	NomorKontrak  string  `json:"nomor_kontrak" validate:"required"`
	NIK           string  `json:"nik" validate:"required"`
	NamaAsset     string  `json:"nama_asset" validate:"required"`
	OTR           float64 `json:"otr" validate:"required"`
	AdminFee      float64 `json:"admin_fee" validate:"required"`
	JumlahCicilan float64 `json:"jumlah_cicilan" validate:"required"`
	JumlahBunga   float64 `json:"jumlah_bunga" validate:"required"`
}
