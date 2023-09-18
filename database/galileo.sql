DROP TABLE IF EXISTS `project`;
CREATE TABLE `project`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `identifier` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_by` int(11) unsigned NOT NULL,
    `created_at` datetime                                NOT NULL,
    `update_at`  datetime                                NOT NULL,
    `update_by`  int(11) unsigned NOT NULL,
    `deleted_at` datetime,
    `deleted_by` int(11) unsigned,
    `status`     int (2) unsigned NOT NULL,
    `remark`     longtext COLLATE utf8mb4_unicode_ci,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`)
)

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`     varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
    `chinese_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `nickname`     varchar(255)                           DEFAULT NULL,
    `role`         int(2) DEFAULT NULL,
    `phone`        varchar(255)                           NOT NULL,
    `password`     varchar(255)                           NOT NULL,
    `avatar`       varchar(255)                           DEFAULT NULL,
    `email`        varchar(255)                           NOT NULL,
    `status`       boolean                                DEFAULT true,
    `update_at`    datetime                               NOT NULL,
    `created_at`   datetime                               NOT NULL,
    `deleted_at`   datetime,
    `deleted_by`   int(11) unsigned,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;