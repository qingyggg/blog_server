/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package constants

// connection information
const (
	MySQLDefaultDSN = "storybook:blog_server123@tcp(127.0.0.1:18000)/storybook?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"

	MinioEndPoint        = "localhost:18001"
	MinioAccessKeyID     = "blog_server"
	MinioSecretAccessKey = "blog_server123"
	MiniouseSSL          = false

	RedisAddr     = "localhost:18003"
	RedisPassword = "blog_server123"

	MongoDefaultDSN = "mongodb://blog_server:blog_server123@localhost:18006/?connect=direct"
)

// constants in the project
const (
	UserTableName      = "users"
	FollowsTableName   = "follows"
	VideosTableName    = "videos"
	MessageTableName   = "messages"
	FavoritesTableName = "likes"
	CommentTableName   = "comments"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

	MinioVideoBucketName = "videobucket"
	MinioImgBucketName   = "imagebucket"

	TestSign       = "测试账号！ offer"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)
