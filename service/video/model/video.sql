CREATE TABLE `video`  (
                          `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) NULL DEFAULT NULL,
                          `updated_at` datetime(3) NULL DEFAULT NULL,
                          `deleted_at` datetime(3) NULL DEFAULT NULL,
                          `author_id` int NULL DEFAULT NULL,
                          `play_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `cover_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `favorite_count` int NULL DEFAULT 0,
                          `comment_count` int NULL DEFAULT 0,
                          `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
