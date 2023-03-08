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
    `deleted_at` datetime NULL,
    `deleted_by` int(11) unsigned NOT NULL,
    `status`     int (2) unsigned NOT NULL,
    `remark`     longtext COLLATE utf8mb4_unicode_ci     NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`)
)