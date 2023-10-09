CREATE TABLE IF NOT EXISTS `konsumen`
(
    `nik` varchar(100) NOT NULL,
    `full_name` varchar(100) NOT NULL,
    `legal_name` varchar(100) NOT NULL,
    `tempat_lahir` varchar(100) NOT NULL,
    `tanggal_lahir` date NOT NULL,
    `gaji` decimal(15,2) unsigned NOT NULL,
    `foto_ktp` varchar(100) NOT NULL,
    `foto_selfie` varchar(100) NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`nik`)
)