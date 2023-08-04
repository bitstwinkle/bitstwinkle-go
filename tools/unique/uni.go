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

package unique

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/rs/xid"
	"time"
)

// ID 生成全网唯一ID
func ID() string {
	return xid.New().String()
}

// Rand 生成更具随机性的字符串
func Rand() string {
	var suffix string
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		timeStr := fmt.Sprintf("%d", time.Now().Unix())
		if len(timeStr) > 8 {
			suffix = timeStr[:8]
		} else {
			suffix = timeStr
		}
	} else {
		suffix = hex.EncodeToString(randomBytes)
	}
	return xid.New().String() + suffix
}
