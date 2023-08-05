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

package security

import (
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/tools/sign"
	"github.com/bitstwinkle/bitstwinkle-go/tools/unique"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	Security            = "security"
	Turn                = "turn"
	Auth                = "auth"
	RefreshTokenPub     = "refresh_token_pub"
	RefreshTokenPubSign = "refresh_token_pub_sign"
	HeaderPrefix        = "Twinkle-"                  //Uniform prefix
	HeaderTokenPub      = HeaderPrefix + "Token-Pub"  //Public
	HeaderNonce         = HeaderPrefix + "Nonce"      //Nonce
	HeaderTimestamp     = HeaderPrefix + "Timestamp"  //Timestamp
	HeaderSignature     = HeaderPrefix + "Signature"  //Signature
	HeaderTokenExpire   = HeaderPrefix + "Expiration" //Token Expiration
)

var signWithHeaderKey = []string{HeaderNonce, HeaderTimestamp, HeaderTokenPub}

// Signature 按照协议签名
func Signature(req *http.Request, tokenPub string, tokenPri string) *errors.Error {
	nonce := unique.Rand()
	timestamp := time.Now().Unix()
	req.Header.Set(HeaderTokenPub, tokenPub)
	req.Header.Set(HeaderNonce, nonce)
	req.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signStr, err := GenSignature(req, tokenPri)
	if err != nil {
		return err
	}
	req.Header.Set(HeaderSignature, signStr)
	return nil
}

// GenSignature 对请求数据进行签名
func GenSignature(req *http.Request, tokenPri string) (string, *errors.Error) {
	wrapper := make(map[string]string)

	if len(req.URL.Query()) > 0 {
		err := urlValuesToMap(req.URL.Query(), wrapper)
		if err != nil {
			return strs.EMPTY, err
		}
	}

	if len(req.PostForm) == 0 {
		err := urlValuesToMap(req.PostForm, wrapper)
		if err != nil {
			return strs.EMPTY, err
		}
	}

	sortedKeys := make([]string, 0, len(wrapper))
	for key := range wrapper {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	joinBuf := strings.Builder{}
	for _, key := range sortedKeys {
		_, err := joinBuf.WriteString(wrapper[key])
		if err != nil {
			return strs.EMPTY, errors.Sys(err.Error(), err)
		}
	}

	if req.Body != nil {
		byteData, _ := io.ReadAll(req.Body)
		_, err := joinBuf.Write(byteData)
		if err != nil {
			return strs.EMPTY, errors.Sys(err.Error(), err)
		}
	}

	for _, headerKey := range signWithHeaderKey {
		_, err := joinBuf.WriteString(req.Header.Get(headerKey))
		if err != nil {
			return strs.EMPTY, errors.Sys(err.Error(), err)
		}
	}

	joinStr := joinBuf.String()
	signStr, err := sign.Sign(joinStr, tokenPri)
	if err != nil {
		return strs.EMPTY, err
	}

	return signStr, nil
}

func urlValuesToMap(inputData url.Values, wrapper map[string]string) *errors.Error {
	if len(inputData) == 0 {
		return nil
	}

	for k, values := range inputData {
		switch len(values) {
		case 0:
			continue
		case 1:
			wrapper[k] = values[0]
			continue
		}
		buf := strings.Builder{}
		for _, v := range values {
			_, err := buf.WriteString(v)
			if err != nil {
				return errors.Sys(err.Error(), err)
			}
		}
		wrapper[k] = buf.String()
	}
	return nil
}
