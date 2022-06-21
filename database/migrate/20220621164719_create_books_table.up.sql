CREATE TABLE IF NOT EXISTS `books`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(255),
    `author` varchar(255),
    `desc` text,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
