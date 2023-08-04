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

package collections

// GetArrayIntersection 获取数组交集
// 返回值保证不为nil
func GetArrayIntersection[T any](arrays [][]T, getKey func(m T) string) []T {
	if len(arrays) == 0 {
		return []T{}
	}
	var intersection []T

	elements := make(map[string]T)
	for _, m := range arrays[0] {
		k := getKey(m)
		elements[k] = m
	}

	for _, arr := range arrays[1:] {
		temp := make(map[string]T)
		for _, m := range arr {
			k := getKey(m)
			_, has := elements[k]
			if has {
				temp[k] = m
			}
		}
		elements = temp
	}

	for _, m := range elements {
		intersection = append(intersection, m)
	}

	if intersection == nil {
		return []T{}
	}

	return intersection
}
