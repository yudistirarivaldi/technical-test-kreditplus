CREATE TABLE `consumer_limits` (
  `id` int NOT NULL AUTO_INCREMENT,
  `consumer_id` int NOT NULL,
  `tenor` int NOT NULL,
  `limit_amount` double NOT NULL,
  `used_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `consumers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nik` varchar(191) NOT NULL,
  `full_name` longtext NOT NULL,
  `legal_name` longtext NOT NULL,
  `birth_place` longtext,
  `birth_date` date DEFAULT NULL,
  `salary` double DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `ktp_photo` longtext,
  `selfie_photo` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_consumers_nik` (`nik`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `contract_number` varchar(191) NOT NULL,
  `consumer_id` bigint unsigned NOT NULL,
  `otr` double NOT NULL,
  `admin_fee` double DEFAULT NULL,
  `installment` double DEFAULT NULL,
  `interest` double DEFAULT NULL,
  `asset_name` longtext,
  `source_channel` varchar(100) NOT NULL DEFAULT 'web',
  `tenor` int NOT NULL DEFAULT '1',
  `down_payment` decimal(18,2) DEFAULT '0.00',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_transactions_contract_number` (`contract_number`),
  KEY `fk_consumers_transactions` (`consumer_id`),
  CONSTRAINT `fk_consumers_transactions` FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;