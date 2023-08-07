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
	"regexp"
	"strconv"
)

const (
	EMPTY = ""
)

// Int64Of String to number
// str: Corresponding string
// d: Default return value if not found
func Int64Of(str string, d int64) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return d
	}
	return i
}

func BoolOf(str string) bool {
	if str == "true" ||
		str == "TRUE" ||
		str == "yes" ||
		str == "YES" {
		return true
	}
	return false
}

func KeyVerify(str string) bool {
	if len(str) > 64 {
		return false
	}
	match, err := regexp.MatchString("^[a-zA-Z0-9_@-]+$", str)
	if err != nil {
		return false
	}
	return match
}

// FBCut Intercept before and after, remaining specified length
func FBCut(str string, max int) string {
	l := len(str)
	if l <= max {
		return str
	}
	if max <= 10 {
		return str[:max]
	}
	f := (max - 10) / 2
	b := l - ((max - 10) - f)
	return str[:f] + "**********" + str[b:]
}
