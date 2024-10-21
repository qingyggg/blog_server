SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment_favorite
-- ----------------------------
DROP TABLE IF EXISTS `comment_favorite`;
CREATE TABLE `comment_favorite` (
                                  `id`         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                                  `article_id` BINARY(32) NOT NULL COMMENT '评论文章ID',
                                  `comment_id` BINARY(32) NOT NULL COMMENT '被点赞或踩的评论 ID',
                                  `user_id`    BINARY(8) NOT NULL COMMENT '用户ID',
                                  `status`     SMALLINT NOT NULL COMMENT '2：踩, 1：点赞',
                                  INDEX `commentIdIdx` (`comment_id`),              -- 查询评论被点赞或者踩的数量
                                  INDEX `user_comment_IdIdx` (`user_id`, `comment_id`), -- 用户对该评论的状态，用户对该评论进行点赞和踩
                                  INDEX `statusIdx` (`status`)                      -- 可选
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '用户对评论的点赞或者踩';

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
                         `id`          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
                         `user_id`     BINARY(8) NOT NULL COMMENT '用户ID',
                         `follower_id` BINARY(8) NOT NULL COMMENT '粉丝ID',
                         INDEX `FollowerIdIdx` (`follower_id`),
                         INDEX `userIdIdx` (`user_id`),
                         INDEX `userIdToFollowerIdIdx` (`user_id`, `follower_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '关注表';

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                       `id`               BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
                       `hash_id`          BINARY(8) NOT NULL UNIQUE COMMENT '用户的hash值',
                       `user_name`        TINYTEXT NOT NULL UNIQUE COMMENT '用户名',
                       `password`         VARCHAR(255) NOT NULL COMMENT '用户密码',
                       `avatar`           VARCHAR(255) NOT NULL COMMENT '用户头像',
                       `background_image` VARCHAR(255) NOT NULL COMMENT '用户个人页顶部大图',
                       `signature`        TINYTEXT NOT NULL COMMENT '个人简介',
                       INDEX `user_name_password_idx` (`user_name`, `password`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
                          `id`            BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '文章ID',
                          `user_id`    BINARY(8) NOT NULL COMMENT '作者ID',
                          `title`         TINYTEXT NOT NULL COMMENT '文章标题',
                          `note`          TINYTEXT NOT NULL COMMENT '文章小记',
                          `cover_url`     TEXT NOT NULL COMMENT '背景图URL',
                          `publish_time`  timestamp NOT NULL COMMENT '发布时间戳',
                          `hash_id`          BINARY(32) NOT NULL UNIQUE COMMENT '文章的hashID值',
                          INDEX `userIdIdx` (`user_id`),
                          INDEX `userIdToHashIdIdx` (`user_id`, `hash_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '文章表';

-- ----------------------------
-- Table structure for article_favorite
-- ----------------------------
DROP TABLE IF EXISTS `article_favorite`;
CREATE TABLE `article_favorite` (
                                  `id`         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                                  `article_id` BINARY(32) NOT NULL COMMENT '评论文章ID',
                                  `user_id`    BINARY(8) NOT NULL COMMENT '用户ID',
                                  `status`     SMALLINT NOT NULL COMMENT '2：踩, 1：点赞',
                                  INDEX `articleIdIdx` (`article_id`),              -- 查询评论被点赞或者踩的数量
                                  INDEX `user_article_IdIdx` (`user_id`, `article_id`), -- 用户对该评论的状态，用户对该评论进行点赞和踩
                                  INDEX `statusIdx` (`status`)                      -- 可选
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '文章点赞表';

-- ----------------------------
-- Table structure for article_collect
-- ----------------------------
DROP TABLE IF EXISTS `article_collect`;
CREATE TABLE `article_collect` (
                                 `id`         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                                 `article_id` BINARY(32) NOT NULL COMMENT '文章ID',
                                 `user_id`    BINARY(8) NOT NULL COMMENT '用户ID',
                                `tag` varchar(255) NOT NULL COMMENT '收藏分类',
                                 INDEX `articleIdIdx` (`article_id`),
                                 INDEX `user_article_IdIdx` (`user_id`, `article_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '收藏表';

DROP TABLE IF EXISTS `article_view`;
CREATE TABLE `article_view`(
                             `id`         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                             `article_id` BINARY(32) NOT NULL UNIQUE COMMENT '文章ID',
                            `view_count` BIGINT DEFAULT 0 COMMENT '文章阅读数',
                            INDEX `articleIdIdx` (`article_id`)
)AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb4 COMMENT = '统计文章阅读数的表';

SET FOREIGN_KEY_CHECKS = 1;
