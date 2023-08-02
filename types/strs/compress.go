/*
 *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 */

package strs

import (
	"fmt"
	"strings"
)

// NumbStrCompress 限制: 每个字符串的长度不能超过100
func NumbStrCompress(numbStr ...string) string {
	count := len(numbStr)
	arr := make([]string, 2*count)
	for i, item := range numbStr {
		arr[i] = item
	}
	for i := count; i < 2*count; i += 1 {
		arr[i] = fmt.Sprintf("%2d", len(numbStr[i-count]))
	}
	return strings.Join(arr, EMPTY)
}
