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

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	cost := 5

	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GetSHA256Byte(data string) []byte {
	// 生成 SHA256 哈希
	hash := sha256.New()
	hash.Write([]byte(data))
	hashedData := hash.Sum(nil)

	return hashedData
}

func GetSHA256String(data string) string {
	// 生成 SHA256 哈希
	hash := sha256.New()
	hash.Write([]byte(data))
	hashedData := hash.Sum(nil)

	return hex.EncodeToString(hashedData)
}

// ConvertByteHashToString 将 byte 类型的哈希转换为字符串哈希
func ConvertByteHashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}

// ConvertStringHashToByte 将字符串哈希转换为 byte 类型的哈希
func ConvertStringHashToByte(hash string) []byte {
	data, _ := hex.DecodeString(hash)
	return data
}
