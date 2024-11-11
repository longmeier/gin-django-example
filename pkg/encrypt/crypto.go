package encrypt

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

//func Test(){
//	// 生成一个随机的盐值
//	salt, err := GenerateSalt(16) // 16 字节的盐值
//	if err != nil {
//		log.Fatalf("Failed to generate salt: %v", err)
//	}
//
//	// 输入密码
//	password := "my_secure_password"
//
//	// 生成 PBKDF2 哈希
//	hash, err := GeneratePBKDF2Hash(password, salt, 100000)
//	if err != nil {
//		log.Fatalf("Failed to generate hash: %v", err)
//	}
//}

// GeneratePBKDF2Hash 使用 PBKDF2 和 SHA-256 算法生成密码哈希
func GeneratePBKDF2Hash(password, salt string, iterations int) (string, error) {
	// 转换密码和盐为字节数组
	passBytes := []byte(password)
	saltBytes := []byte(salt)

	// 使用 PBKDF2 算法生成哈希值
	hash := pbkdf2.Key(passBytes, saltBytes, iterations, sha256.Size, sha256.New)

	// 将生成的哈希值编码为 Base64 格式
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	// 返回类似格式的密码哈希
	return fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", iterations, salt, hashBase64), nil
}

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}
