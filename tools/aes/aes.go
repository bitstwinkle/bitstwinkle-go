/*
 *
 *  *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *  *
 *  *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  *  * you may not use this file except in compliance with the License.
 *  *  * You may obtain a copy of the License at
 *  *  *
 *  *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *  *
 *  *  * Unless required by applicable law or agreed to in writing, software
 *  *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  *  * See the License for the specific language governing permissions and
 *  *  * limitations under the License.

 */

package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"io"
)

// Encrypt 加密
func Encrypt(content string, priKey []byte) (string, *errors.Error) {
	plaintext := []byte(content)
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return strs.EMPTY, errors.Sys("aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedPlaintext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return strs.EMPTY, errors.Sys("init iv failed: " + err.Error())
	}

	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)

	copy(ciphertext[:aes.BlockSize], iv)

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt 解密
func Decrypt(content string, priKey []byte) (string, *errors.Error) {
	ciphertext, err := hex.DecodeString(content)
	if err != nil {
		return strs.EMPTY, errors.Sys("hex.DecodeString(content):" + err.Error())
	}

	block, err := aes.NewCipher(priKey)
	if err != nil {
		return strs.EMPTY, errors.Sys("aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}

	iv := ciphertext[:aes.BlockSize]

	// 使用CBC模式进行解密
	decrypted := make([]byte, len(ciphertext)-aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext[aes.BlockSize:])

	// 去除填充数据
	padding := int(decrypted[len(decrypted)-1])

	decrypted = decrypted[:len(decrypted)-padding]

	return string(decrypted), nil
}
