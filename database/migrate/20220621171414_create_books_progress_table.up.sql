CREATE TABLE IF NOT EXISTS `books_progresses`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `books_id` BIGINT UNSIGNED NOT NULL,
    `progress` int(10) UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;