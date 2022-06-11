CREATE TABLE `comment`  (
                            `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) NULL DEFAULT NULL,
                            `updated_at` datetime(3) NULL DEFAULT NULL,
                            `deleted_at` datetime(3) NULL DEFAULT NULL,
                            `user_id` int UNSIGNED NULL DEFAULT NULL,
                            `video_id` int UNSIGNED NULL DEFAULT NULL,
                            `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                                INDEX `idx_comment_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;