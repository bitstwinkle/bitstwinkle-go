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
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"math/big"
	"strconv"
	"strings"
)

// NumbStrCompress 数字型字符串压缩。限制: 每个字符串的长度不能超过100
// 1.求出每个字符串的长度
// 2.将字符串和长度数字(最多两位)全部拼接位一个长的数字型字符串
// 3.对长的字符串进行62进制转码
func NumbStrCompress(numbStr ...string) (string, *errors.Error) {
	count := len(numbStr)
	if count == 0 {
		return EMPTY, nil
	}
	switch count {
	case 0:
		return EMPTY, nil
	case 1:
		return doNumStrTo62(numbStr[0])
	default:
		longStr, err := doGenLongArray(numbStr)
		if err != nil {
			return EMPTY, err
		}
		return doNumStrTo62(longStr)
	}
}

// NumbStrDecompress 解压
func NumbStrDecompress(infoStr string, count int) ([]string, *errors.Error) {
	if count < 0 || 2*count > len(infoStr) {
		return nil, errors.Assert("count<0 || 2 * count > strLen", "false")
	}
	bigNumb := big.NewInt(0)
	bigNumb, ok := bigNumb.SetString(infoStr, 62)
	if !ok {
		return nil, errors.Assert("bigNumb.SetString("+infoStr+", 62) ok", "false")
	}
	longStr := bigNumb.Text(10)
	strLen := len(longStr)
	weiStr := longStr[strLen-2*count:]
	weiArr, err := doSplitNumber(weiStr)
	if err != nil {
		return nil, err
	}
	strArr := make([]string, count)
	start := 0
	for i, wei := range weiArr {
		strArr[i] = longStr[start : start+wei]
		start += wei
	}
	return strArr, nil
}

func doGenLongArray(numbStr []string) (string, *errors.Error) {
	count := len(numbStr)
	arr := make([]string, 2*count)

	for i, item := range numbStr {
		arr[i] = item
	}
	for i := count; i < 2*count; i += 1 {
		strLen := len(numbStr[i-count])
		if strLen >= 100 {
			return EMPTY, errors.Assert("numbString.len < 100", fmt.Sprintf("%d", strLen))
		}
		arr[i] = fmt.Sprintf("%02d", strLen)
	}
	return strings.Join(arr, EMPTY), nil
}

func doNumStrTo62(str string) (string, *errors.Error) {
	bigNumb := big.NewInt(0)
	ok := false
	bigNumb, ok = bigNumb.SetString(str, 10)
	if !ok {
		return EMPTY, errors.Assert("bigNumb.SetString("+str+", 10) ok", "false")
	}
	return bigNumb.Text(62), nil
}

func doSplitNumber(numStr string) ([]int, *errors.Error) {
	var result []int

	for i := 0; i < len(numStr); i += 2 {
		end := i + 2
		if end > len(numStr) {
			end = len(numStr)
		}

		num, err := strconv.Atoi(numStr[i:end])
		if err != nil {
			return nil, errors.Verify("invalid number", err)
		}
		result = append(result, num)
	}

	return result, nil
}
