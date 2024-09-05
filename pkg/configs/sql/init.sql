SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
create table comments
(
  id                 bigint auto_increment comment '评论 ID'
        primary key,
  user_id            bigint                              not null comment '评论发布用户ID',
  article_id         bigint                              not null comment '评论文章ID',
  replied_comment_id bigint    default 777               null comment '被回复的评论的ID',
  floor              smallint                            not null comment '评论等级，分为三级',
  comment_text       varchar(255)                        not null comment '评论内容',
  created_at         timestamp default CURRENT_TIMESTAMP not null comment '评论创建时间',
  deleted_at         timestamp                           null comment '评论删除时间'
)
  comment '评论表' charset = utf8mb3;

create index articleIdIdx
  on comments (article_id)
  comment '根据文章ID，评论 ID 索引';

create index repliedIdIdx
  on comments (replied_comment_id)
  comment '根据被回复的评论ID，评论 ID 索引';
-- ----------------------------
-- Table structure for comment_favorite
-- ----------------------------
DROP TABLE IF EXISTS `comment_favorite`;
create table comment_favorite
(
  id         bigint auto_increment comment '主键'
    primary key,
  article_id bigint   not null comment '评论文章ID',
  comment_id bigint   not null comment '被点赞或踩的评论 ID',
  user_id    bigint   not null comment '用户ID',
  status     smallint not null comment '-1：踩,1：点赞'
)
  comment '用户对评论的点赞或者踩' charset = utf8mb3;

create index commentIdIdx
  on comment_favorite (comment_id);

create index userIdIdx
  on comment_favorite (user_id);

create index user_comment_IdIdx
  on comment_favorite (user_id, comment_id);
-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
create table follows
(
  id          bigint auto_increment comment '自增主键'
        primary key,
  user_id     bigint                              not null comment '用户id',
  follower_id bigint                              not null comment '粉丝id',
  created_at  timestamp default CURRENT_TIMESTAMP not null comment '关注关系创建时间',
  deleted_at  timestamp                           null comment '关注关系删除时间'
)
  comment '关注表' charset = utf8mb3;

create index FollowerIdIdx
  on follows (follower_id);

create index userIdIdx
  on follows (user_id);

create index userIdToFollowerIdIdx
  on follows (user_id, follower_id);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
create table users
(
  id               bigint auto_increment comment '用户ID'
        primary key,
  user_name        varchar(255) not null comment '用户名',
  password         varchar(255) not null comment '用户密码',
  avatar           varchar(255) not null comment '用户头像',
  background_image varchar(255) not null comment '用户个人页顶部大图',
  signature        varchar(255) not null comment '个人简介',
  constraint user_name
    unique (user_name)
)
  comment '用户表' charset = utf8mb3;

create index user_name_password_idx
  on users (user_name, password);
-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
create table articles
(
  id            bigint auto_increment comment '文章ID'
    primary key,
  user_id       bigint       not null comment '作者id',
  like_count    bigint       not null comment '点赞用户数目',
  comment_count bigint       not null comment '评论数目',
  collect_count bigint       not null comment '收藏数目',
  title         varchar(255) not null comment '文章标题',
  note          tinytext     not null comment '文章小记',
  content       mediumtext   not null comment '文章内容',
  cover_url     varchar(255) not null comment '背景图url'
)
  comment '文章表' charset = utf8mb3;

create index userIdIdx
  on articles (user_id);

-- ----------------------------
-- Table structure for article_favorite
-- ----------------------------
create table article_favorite
(
  id         bigint auto_increment comment '自增主键'
    primary key,
  user_id    bigint                              not null comment '点赞用户id',
  article_id bigint                              not null comment '被点赞的文章id',
  status     smallint                            not null comment '-1：踩,1：点赞',
  created_at timestamp default CURRENT_TIMESTAMP not null comment '点赞创建时间',
  deleted_at timestamp                           null comment '点赞删除时间'
)
  comment '点赞表' charset = utf8mb3;

create index articleIdx
  on article_favorite (article_id);

create index userIdIdx
  on article_favorite (user_id);

create index userIdtoArticleIdIdx
  on article_favorite (user_id, article_id);


-- ----------------------------
-- Table structure for article_collect
-- ----------------------------
create table article_collect
(
  id           bigint auto_increment comment '自增主键'
    primary key,
  user_id      bigint                                 not null comment '用户id',
  article_id   bigint                                 not null comment '文章id',
  collect_name varchar(255) default 'default'         null comment '收藏的类型',
  created_at   timestamp    default CURRENT_TIMESTAMP not null comment '收藏创建时间',
  deleted_at   timestamp                              null comment '收藏删除时间'
)
  comment '收藏表' charset = utf8mb3;

create index articleIdx
  on article_collect (article_id);

create index userIdIdx
  on article_collect (user_id);

SET FOREIGN_KEY_CHECKS = 1;


