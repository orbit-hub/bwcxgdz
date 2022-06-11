CREATE TABLE `user`  (
                         `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `follow_count` bigint NULL DEFAULT 0,
                         `follower_count` bigint NULL DEFAULT 0,
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `name`(`name`) USING BTREE,
                         INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE,
                         INDEX `idx_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
