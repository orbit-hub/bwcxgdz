CREATE TABLE `relation`  (
                             `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                             `created_at` datetime(3) NULL DEFAULT NULL,
                             `updated_at` datetime(3) NULL DEFAULT NULL,
                             `deleted_at` datetime(3) NULL DEFAULT NULL,
                             `user_id` int NULL DEFAULT NULL,
                             `follow` int NULL DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
