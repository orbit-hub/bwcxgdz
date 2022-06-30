CREATE TABLE `video`  (
                          `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) NULL DEFAULT NULL,
                          `updated_at` datetime(3) NULL DEFAULT NULL,
                          `deleted_at` datetime(3) NULL DEFAULT NULL,
                          `author_id` int NOT NULL,
                          `play_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `cover_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `favorite_count` int NOT NULL DEFAULT 0,
                          `comment_count` int NOT NULL DEFAULT 0,
                          `publish_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
                          `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
