CREATE TABLE `comments`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  `cid` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `product` int UNSIGNED NOT NULL,
  `commentator` int UNSIGNED NOT NULL,
  `response_to` int UNSIGNED NULL DEFAULT NULL,
  `text` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`cid`) USING BTREE,
  UNIQUE INDEX `comments_cid_idx`(`cid`) USING BTREE,
  INDEX `product`(`product`) USING BTREE,
  INDEX `commentator`(`commentator`) USING BTREE,
  INDEX `response_to`(`response_to`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 152 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

CREATE TABLE `images`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  `iid` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `products` int UNSIGNED NOT NULL,
  UNIQUE INDEX `images_iid_idx`(`iid`) USING BTREE,
  INDEX `dddd`(`products`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

CREATE TABLE `messages`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  `mid` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `from` int UNSIGNED NOT NULL,
  `to` int UNSIGNED NOT NULL,
  `text` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`mid`) USING BTREE,
  UNIQUE INDEX `messages_mid_idx`(`mid`) USING BTREE,
  INDEX `from`(`from`) USING BTREE,
  INDEX `to`(`to`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

CREATE TABLE `products`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  `pid` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `price` decimal(6, 2) NOT NULL,
  `owner` int UNSIGNED NOT NULL,
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`pid`) USING BTREE,
  UNIQUE INDEX `products_pid_idx`(`pid`) USING BTREE,
  INDEX `owner`(`owner`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 135 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

CREATE TABLE `users`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  `uid` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `hashed` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码哈希',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像',
  `info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '简历',
  `profile` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '详历',
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '默认地址',
  PRIMARY KEY (`uid`) USING BTREE,
  INDEX `uid`(`uid`, `username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 161 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

ALTER TABLE `comments` ADD CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`product`) REFERENCES `products` (`pid`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `comments` ADD CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`commentator`) REFERENCES `users` (`uid`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `comments` ADD CONSTRAINT `comments_ibfk_3` FOREIGN KEY (`response_to`) REFERENCES `comments` (`cid`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `images` ADD CONSTRAINT `images_ibfk_1` FOREIGN KEY (`products`) REFERENCES `products` (`pid`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `messages` ADD CONSTRAINT `messages_ibfk_1` FOREIGN KEY (`from`) REFERENCES `users` (`uid`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `messages` ADD CONSTRAINT `messages_ibfk_2` FOREIGN KEY (`to`) REFERENCES `users` (`uid`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `products` ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`owner`) REFERENCES `users` (`uid`) ON DELETE RESTRICT ON UPDATE RESTRICT;

