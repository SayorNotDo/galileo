DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`        varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
    `chinese_name`    varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `nickname`        varchar(255)                           DEFAULT NULL,
    `hashed_password` tinyblob                               NOT NULL,
    `avatar`          varchar(255)                           DEFAULT NULL,
    `email`           varchar(255)                           NOT NULL,
    `status`          int(11) NOT NULL,
    `update_at`       datetime                               NOT NULL,
    `created_at`      datetime                               NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`username`)

) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;