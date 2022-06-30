CREATE TABLE `relation`  (
                             `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                             `created_at` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3),
                             `user_id` int NOT NULL,
                             `follow` int NOT NULL,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
