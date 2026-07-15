package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

type Md5Helper struct {
}

var Md5HelperObject Md5Helper

/**
 *  @param value
 *  @return string
 */
func (i *Md5Helper) Md5ToString(value string) string {
	// 创建MD5哈希实例
	hasher := md5.New()

	// 写入数据到哈希实例
	hasher.Write([]byte(value)) // 注意：这里需要将字符串转换为字节切片

	// 获取哈希值，并转换为十六进制字符串
	hashInBytes := hasher.Sum(nil) // Sum方法返回一个字节切片，代表最终的哈希值
	hashString := fmt.Sprintf("%x", hashInBytes)
	return hashString
}

func (i *Md5Helper) Sha256ToString(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))    // 注意：这里需要将字符串转换为字节切片
	hashInBytes := hasher.Sum(nil) // Sum方法返回一个字节切片，代表最终的哈希值
	hashString := fmt.Sprintf("%x", hashInBytes)
	return hashString
}
