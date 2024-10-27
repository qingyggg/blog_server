# blog_server

## Project Introduction

This project is developed based on TikTok Demo, absorbing the code organization and design ideas of the project. Through the analysis and improvement of the existing code, we are committed to building an efficient and scalable blog server.

## Technology Stack

- **Hertz**: High-performance network framework

- **Gorm_gen**: Convenient ORM tool

- **MariaDB**: Relational database

- **MongoDB**: Document database for processing multi-level comments and storing article content

- **Minio**: High-performance object storage for storing avatars or background photos

- **Redis**: Cache and data storage

- **Swagger**: API document generation tool

## Database Design

### Redis + MariaDB

We use a combination of Redis and MariaDB and design timed synchronization and asynchronous synchronization strategies. This design enables us to efficiently handle a large number of data requests while ensuring data consistency and integrity when performing database operations.

### MongoDB

MongoDB is used to handle database operations for multi-level comments, using closure table technology. This design allows us to efficiently manage the hierarchical relationship of comments when storing and retrieving comments, improving the flexibility of data access.

## Deployment logic

### Prerequisites

- Docker

### Deployment steps

1. Use [mkcert](https://github.com/FiloSottile/mkcert) to generate certificates, named `blog_server.pem` and `blog_server-key.pem` respectively.
2. Put the certificate files under `pkg/configs/nginx/certs`.
3. Run the following Docker Compose command:

```bash
docker compose -f ./docker-build/dep_setup.yml up -d
```

## API documentation

Swagger and ApiFox documents are integrated in the project to facilitate developers to understand and use the API. Detailed API interface information is placed in the docs folder in the root directory of the project, and the files in it can be imported into apifox

Or you can access it through the following link:
#### apifox: https://apifox.com/apidoc/shared-a0db5be8-f02e-48d8-be07-7f06b3993ec2

## Conclusion
As a college student, without a lot of work experience, I would like to thank the author of tiktok demo for letting me learn hertz's project design technology, which has brought me one step closer to becoming a qualified backend engineer. I would also like to sincerely thank ByteDance for writing so many open source projects and detailed documents.
