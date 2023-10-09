CREATE TABLE IF NOT EXISTS `transaksi`
(
    `nomor_kontrak` varchar(100) NOT NULL,
    `konsumen_id` varchar(100) NOT NULL,
    `nama_asset` varchar(100) NOT NULL,
    `otr` decimal(15,2) unsigned NOT NULL,
    `admin_fee` decimal(15,2) unsigned NOT NULL,
    `jumlah_cicilan` decimal(15,2) unsigned NOT NULL,
    `jumlah_bunga` decimal(15,2) unsigned NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`nomor_kontrak`)
)