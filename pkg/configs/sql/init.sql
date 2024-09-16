SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
                          `id`                 BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '评论 ID',
                          `user_id`            BIGINT NOT NULL COMMENT '评论发布用户ID',
                          `article_id`         BIGINT NOT NULL COMMENT '评论文章ID',
                          `replied_comment_id` BIGINT DEFAULT 777 COMMENT '被回复的评论的ID',
                          `floor`              SMALLINT NOT NULL COMMENT '评论等级，分为三级',
                          `comment_text`       VARCHAR(255) NOT NULL COMMENT '评论内容',
                          `created_at`         TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '评论创建时间',
                          `deleted_at`         TIMESTAMP NULL COMMENT '评论删除时间',
                          INDEX `articleIdIdx` (`article_id`) COMMENT '根据文章ID，评论 ID 索引',
                          INDEX `repliedIdIdx` (`replied_comment_id`) COMMENT '根据被回复的评论ID，评论 ID 索引'
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '评论表';

-- ----------------------------
-- Table structure for comment_favorite
-- ----------------------------
DROP TABLE IF EXISTS `comment_favorite`;
CREATE TABLE `comment_favorite` (
                                  `id`         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                                  `article_id` BIGINT NOT NULL COMMENT '评论文章ID',
                                  `comment_id` BIGINT NOT NULL COMMENT '被点赞或踩的评论 ID',
                                  `user_id`    BIGINT NOT NULL COMMENT '用户ID',
                                  `status`     SMALLINT NOT NULL COMMENT '-1：踩, 1：点赞',
                                  INDEX `commentIdIdx` (`comment_id`),
                                  INDEX `userIdIdx` (`user_id`),
                                  INDEX `user_comment_IdIdx` (`user_id`, `comment_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '用户对评论的点赞或者踩';

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
                         `id`          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
                         `user_id`     BIGINT NOT NULL COMMENT '用户ID',
                         `follower_id` BIGINT NOT NULL COMMENT '粉丝ID',
                         `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '关注关系创建时间',
                         `deleted_at`  TIMESTAMP NULL COMMENT '关注关系删除时间',
                         INDEX `FollowerIdIdx` (`follower_id`),
                         INDEX `userIdIdx` (`user_id`),
                         INDEX `userIdToFollowerIdIdx` (`user_id`, `follower_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '关注表';

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                       `id`               BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
                       `user_name`        VARCHAR(255) NOT NULL UNIQUE COMMENT '用户名',
                       `password`         VARCHAR(255) NOT NULL COMMENT '用户密码',
                       `avatar`           VARCHAR(255) NOT NULL COMMENT '用户头像',
                       `background_image` VARCHAR(255) NOT NULL COMMENT '用户个人页顶部大图',
                       `signature`        VARCHAR(255) NOT NULL COMMENT '个人简介',
                       INDEX `user_name_password_idx` (`user_name`, `password`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '用户表';

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
                          `id`            BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '文章ID',
                          `user_id`       BIGINT NOT NULL COMMENT '作者ID',
                          `like_count`    BIGINT NOT NULL COMMENT '点赞用户数目',
                          `dislike_count` BIGINT NOT NULL COMMENT '踩的用户数目',
                          `comment_count` BIGINT NOT NULL COMMENT '评论数目',
                          `collect_count` BIGINT NOT NULL COMMENT '收藏数目',
                          `title`         VARCHAR(255) NOT NULL COMMENT '文章标题',
                          `note`          TINYTEXT NOT NULL COMMENT '文章小记',
                          `content`       MEDIUMTEXT NOT NULL COMMENT '文章内容',
                          `cover_url`     VARCHAR(255) NOT NULL COMMENT '背景图URL',
                          `publish_time` timestamp NOT NULL COMMENT '发布时间戳',
                          INDEX `userIdIdx` (`user_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '文章表';

-- ----------------------------
-- Table structure for article_favorite
-- ----------------------------
DROP TABLE IF EXISTS `article_favorite`;
CREATE TABLE `article_favorite` (
                                  `id`         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
                                  `user_id`    BIGINT NOT NULL COMMENT '点赞用户ID',
                                  `article_id` BIGINT NOT NULL COMMENT '被点赞的文章ID',
                                  `status`     SMALLINT NOT NULL COMMENT '-1：踩, 1：点赞',
                                  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '点赞创建时间',
                                  `deleted_at` TIMESTAMP NULL COMMENT '点赞删除时间',
                                  INDEX `articleIdx` (`article_id`),
                                  INDEX `userIdIdx` (`user_id`),
                                  INDEX `userIdtoArticleIdIdx` (`user_id`, `article_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '点赞表';

-- ----------------------------
-- Table structure for article_collect
-- ----------------------------
DROP TABLE IF EXISTS `article_collect`;
CREATE TABLE `article_collect` (
                                 `id`           BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
                                 `user_id`      BIGINT NOT NULL COMMENT '用户ID',
                                 `article_id`   BIGINT NOT NULL COMMENT '文章ID',
                                 `collect_name` VARCHAR(255) DEFAULT 'default' COMMENT '收藏的类型',
                                 `created_at`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '收藏创建时间',
                                 `deleted_at`   TIMESTAMP NULL COMMENT '收藏删除时间',
                                 INDEX `articleIdx` (`article_id`),
                                 INDEX `userIdIdx` (`user_id`)
) AUTO_INCREMENT = 1000 DEFAULT CHARSET = utf8mb3 COMMENT = '收藏表';

SET FOREIGN_KEY_CHECKS = 1;
