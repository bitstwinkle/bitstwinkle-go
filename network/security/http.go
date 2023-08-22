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
	"github.com/bitstwinkle/bitstwinkle-go/tools/logger"
	"github.com/bitstwinkle/bitstwinkle-go/tools/sign"
	"github.com/bitstwinkle/bitstwinkle-go/tools/sys"
	"github.com/bitstwinkle/bitstwinkle-go/tools/unique"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"go.uber.org/zap"
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
	HeaderPrefix        = "Twinkle-" //Uniform prefix

	HeaderSecretPub   = HeaderPrefix + "Secret-Pub" //Secret Pub
	HeaderTokenPub    = HeaderPrefix + "Token-Pub"  //Token Public
	HeaderNonce       = HeaderPrefix + "Nonce"      //Nonce
	HeaderTimestamp   = HeaderPrefix + "Timestamp"  //Timestamp
	HeaderSignature   = HeaderPrefix + "Signature"  //Signature
	HeaderTokenExpire = HeaderPrefix + "Expiration" //Token Expiration

	BodyInside       = "__b_o_d_y__"
	TurnBySecretURL  = "/security/access/secret"
	TurnByRefreshURL = "/security/access/refresh"
)

var signWithHeaderKey = []string{HeaderNonce, HeaderTimestamp}

// TokenSignature Sign in accordance with the agreement
func TokenSignature(req *http.Request, bodyData []byte, tokenPub string, tokenPri string) *errors.Error {
	nonce := unique.Rand()
	timestamp := time.Now().Unix()
	req.Header.Set(HeaderTokenPub, tokenPub)
	req.Header.Set(HeaderNonce, nonce)
	req.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signStr, err := GenSignature(req, bodyData, tokenPri)
	if err != nil {
		return err
	}
	req.Header.Set(HeaderSignature, signStr)
	return nil
}

// SecretSignature Use signatures when exchanging protocols
func SecretSignature(req *http.Request, bodyData []byte, secretPub string, secretPri string) *errors.Error {
	nonce := unique.Rand()
	timestamp := time.Now().Unix()
	req.Header.Set(HeaderSecretPub, secretPub)
	req.Header.Set(HeaderNonce, nonce)
	req.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signStr, err := GenSignature(req, bodyData, secretPri)
	if err != nil {
		return err
	}
	req.Header.Set(HeaderSignature, signStr)
	return nil
}

// GenSignature Sign the request data
func GenSignature(req *http.Request, bodyData []byte, priKey string) (string, *errors.Error) {
	wrapper := make(map[string]string)

	if len(req.URL.Query()) > 0 {
		err := urlValuesToMap(req.URL.Query(), wrapper)
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

	for _, headerKey := range signWithHeaderKey {
		_, _ = joinBuf.WriteString(headerKey + "=")
		_, err := joinBuf.WriteString(req.Header.Get(headerKey))
		if err != nil {
			return strs.EMPTY, errors.Sys(err.Error(), err)
		}
		_, _ = joinBuf.WriteString(";")
	}

	for _, key := range sortedKeys {
		_, _ = joinBuf.WriteString(key + "=")
		_, err := joinBuf.WriteString(wrapper[key])
		if err != nil {
			return strs.EMPTY, errors.Sys(err.Error(), err)
		}
		_, _ = joinBuf.WriteString(";")
	}

	if bodyData != nil && len(bodyData) > 0 {
		_, _ = joinBuf.WriteString(BodyInside + "=")
		_, err := joinBuf.Write(bodyData)
		if err != nil {
			return strs.EMPTY, errors.Sys(err.Error(), err)
		}
	}

	joinStr := joinBuf.String()

	if sys.RunMode.IsRd() {
		logger.Logger.Info("joinStr is ", zap.String("joinStr", joinStr))
	}

	signStr, err := sign.Sign(joinStr, priKey)
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
