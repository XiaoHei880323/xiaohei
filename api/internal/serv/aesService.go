package serv

import (
	"api/internal/svc"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"log"
	"strconv"
)

type AesService struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAesService(ctx context.Context, svcCtx *svc.ServiceContext) AesService {
	return AesService{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 加密
func (s AesService) AesEncryptECB(origData []byte, key []byte) string {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)

	encrypted := make([]byte, len(plain))
	// Block encryption
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return base64.StdEncoding.EncodeToString(encrypted)
}

// 解密
func (s AesService) AesDecryptECB(str string, key []byte) (decrypted []byte) {
	encrypted, _ := base64.StdEncoding.DecodeString(str)

	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// 加密
func (s AesService) AesEncryptECBSecond(origData []byte, key []byte) string {
	block, _ := aes.NewCipher(key)

	nonce, _ := hex.DecodeString("")
	aseGcm, _ := cipher.NewGCM(block)

	ciphertext := aseGcm.Seal(nil, nonce, origData, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

// 解密
func (s AesService) AesDecryptECBSecod(str string, key []byte) string {
	ciphertext, _ := hex.DecodeString(str)
	nonce, _ := hex.DecodeString("")
	block, _ := aes.NewCipher(key)

	aesGcm, _ := cipher.NewGCM(block)

	plaintext, _ := aesGcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext)
}

// AES加密
// iv为空则采用ECB模式，否则采用CBC模式
func (s AesService) AesEncrypt(value, secretKey, iv string) (string, error) {
	if value == "" {
		return "", nil
	}

	//根据秘钥生成16位的秘钥切片
	keyBytes := make([]byte, aes.BlockSize)
	copy(keyBytes, []byte(secretKey))
	//获取block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	blocksize := block.BlockSize()
	valueBytes := []byte(value)

	//填充
	fillsize := blocksize - len(valueBytes)%blocksize
	repeat := bytes.Repeat([]byte{byte(fillsize)}, fillsize)
	valueBytes = append(valueBytes, repeat...)

	result := make([]byte, len(valueBytes))

	//加密
	if iv == "" {
		temp := result
		for len(valueBytes) > 0 {
			block.Encrypt(temp, valueBytes[:blocksize])
			valueBytes = valueBytes[blocksize:]
			temp = temp[blocksize:]
		}
	} else {
		//向量切片
		ivBytes := make([]byte, aes.BlockSize)
		copy(ivBytes, []byte(iv))

		encrypter := cipher.NewCBCEncrypter(block, ivBytes)
		encrypter.CryptBlocks(result, valueBytes)
	}

	//以hex格式数值输出
	encryptText := fmt.Sprintf("%x", result)
	return encryptText, nil
}

// AES解密
// iv为空则采用ECB模式，否则采用CBC模式
func (s AesService) AesDecrypt(value, secretKey, iv string) (string, error) {
	if value == "" {
		return "", nil
	}

	//根据秘钥生成8位的秘钥切片
	keyBytes := make([]byte, aes.BlockSize)
	copy(keyBytes, []byte(secretKey))
	//获取block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	//将hex格式数据转换为byte切片
	valueBytes := []byte(value)
	var encryptedData = make([]byte, len(valueBytes)/2)
	for i := 0; i < len(encryptedData); i++ {
		fmt.Println(value[i*2 : i*2+2])
		b, err := strconv.ParseInt(value[i*2:i*2+2], 16, 10)
		if err != nil {
			return "", err
		}
		encryptedData[i] = byte(b)
	}

	result := make([]byte, len(encryptedData))

	if iv == "" {
		blocksize := block.BlockSize()
		temp := result
		for len(encryptedData) > 0 {
			block.Decrypt(temp, encryptedData[:blocksize])
			encryptedData = encryptedData[blocksize:]
			temp = temp[blocksize:]
		}
	} else {
		//向量切片
		ivBytes := make([]byte, aes.BlockSize)
		copy(ivBytes, []byte(iv))

		//解密
		blockMode := cipher.NewCBCDecrypter(block, ivBytes)
		blockMode.CryptBlocks(result, encryptedData)
	}

	//取消填充
	unpadding := int(result[len(result)-1])
	result = result[:(len(result) - unpadding)]
	return string(result), nil
}

// AES加密 CBC
func (s AesService) EncryptSecod(value, secretKey, iv string) (string, error) {
	//生成cipher.Block 数据块
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		log.Println("错误 -" + err.Error())
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad([]byte(value), blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// AES解密 CBC
func (s AesService) AesDecryptSecond(value, secretKey, iv string) (string, error) {
	decode_data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", nil
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher([]byte(secretKey))
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	return string(unpad(origin_data)), nil

}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
