CREATE TABLE `testing` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `money` FLOAT(20,2) NOT NULL,
  `created_at` DATETIME DEFAULT current_timestamp(),
  `updated_at` DATETIME DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;