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

package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

// Sign 使用TOKEN对数据进行签名
func Sign(data string, token string) (string, *errors.Error) {
	key := []byte(token)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(data))
	if err != nil {
		return strs.EMPTY, errors.Sys(err.Error(), err)
	}
	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}
