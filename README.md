# blog_server

## 项目简介

本项目基于 TikTok Demo 进行开发，吸收了该项目的代码组织方式和设计思路。通过对现有代码的分析与改进，我们致力于构建一个高效、可扩展的博客服务器。

## 技术栈

- **Hertz**: 高性能的网络框架
- **Gorm_gen**: 方便的 ORM 工具
- **MariaDB**: 关系型数据库
- **MongoDB**: 用于处理多级评论和存储文章内容的文档数据库
- **Minio**: 高性能的对象存储，用以存储头像或者背景照片
- **Redis**: 高速缓存和数据存储
- **Swagger**: API 文档生成工具

## 数据库设计

### Redis + MariaDB

我们采用了 Redis 和 MariaDB 的组合，设计了定时同步与异步同步策略。这种设计使得我们在进行数据库操作时，能够高效地处理大量数据请求，同时保证数据的一致性和完整性。

### MongoDB

MongoDB 被用于处理多级评论的数据库操作，采用了闭包表技术。这种设计允许我们在存储和检索评论时，能够高效地管理评论的层级关系，提高了数据访问的灵活性。

## 部署逻辑

### 前置环境

- Docker

### 部署步骤

1. 使用 [mkcert](https://github.com/FiloSottile/mkcert) 生成证书，分别命名为 `blog_server.pem` 和 `blog_server-key.pem`。
2. 将证书文件放在 `pkg/configs/nginx/certs` 下。
3. 运行以下 Docker Compose 命令：

   ```bash
   docker compose -f ./docker-build/dep_setup.yml up -d
   ```



## API 文档

项目中集成了 Swagger 和 ApiFox 文档，便于开发者理解和使用 API。详细的 API 接口信息放置在该项目的根目录中的docs文件夹，可将里面的文件导入到apifox

或者可以通过以下链接访问：
#### apifox: https://apifox.com/apidoc/shared-a0db5be8-f02e-48d8-be07-7f06b3993ec2


## 结语
作为一名大学生，在没有大量的工作经验的前提下，感谢tiktok demo的作者让我学习到了hertz的项目设计技术，让我在称为合格的后端工程师的路上更近一部，也衷心感谢字节跳动，写了那么多开源项目，以及详细的文档



