DROP TABLE IF EXISTS `video`;
CREATE DATABASE `video`;
SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for contribute
-- ----------------------------
DROP TABLE IF EXISTS `contribute`;
CREATE TABLE `contribute`
(
    `user_id`     bigint                                                        NOT NULL COMMENT '用户id',
    `video_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频标题',
    `video_id`    varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '视频md5',
    `created_at`  datetime                                                      NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime                                                      NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `id`          int                                                           NOT NULL AUTO_INCREMENT COMMENT '视频id',
    `deleted_at`  tinyint(1) UNSIGNED ZEROFILL NULL DEFAULT NULL COMMENT '逻辑删除',
    PRIMARY KEY (`id`, `user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 105 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`
(
    `md5`       varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '视频唯一标志',
    `video_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频url',
    `fileName`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名字',
    `image_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片url',
    PRIMARY KEY (`md5`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET
FOREIGN_KEY_CHECKS = 1;
