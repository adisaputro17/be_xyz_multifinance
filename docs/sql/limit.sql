CREATE TABLE IF NOT EXISTS `limit_konsumen`
(
    `limit_id` varchar(100) NOT NULL,
    `konsumen_id` varchar(100) NOT NULL,
    `tenor` int unsigned NOT NULL,
    `limit_pinjam` decimal(15,2) unsigned NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`limit_id`),
    UNIQUE KEY `uq_konsumen_id_tenor` (`konsumen_id`,`tenor`)
)