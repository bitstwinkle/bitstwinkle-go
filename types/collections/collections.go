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

package collections

import (
	"bitstwinkle-go/types/errors"
	"sort"
)

type ItemAble interface {
	GetCode() string
	Verify() *errors.Error
	GetSeq() int
}

type ItemArray []ItemAble

func (arr ItemArray) Len() int           { return len(arr) }
func (arr ItemArray) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }
func (arr ItemArray) Less(i, j int) bool { return arr[i].GetSeq() < arr[j].GetSeq() }

func VerifyArray[T ItemAble](arr []T, notEmpty ...bool) *errors.Error {
	if len(arr) == 0 {
		if len(notEmpty) > 0 && notEmpty[0] {
			return errors.Verify("it is empty")
		}
		return nil
	}
	for _, item := range arr {
		if err := item.Verify(); err != nil {
			return err
		}
	}
	return nil
}

func MapOfArray[T ItemAble](arr []T) (map[string]T, *errors.Error) {
	if len(arr) == 0 {
		return map[string]T{}, nil
	}
	dict := make(map[string]T)
	for _, item := range arr {
		if err := item.Verify(); err != nil {
			return nil, err
		}
		_, ok := dict[item.GetCode()]
		if ok {
			return nil, errors.Assert("single "+item.GetCode(), "repetitive")
		}
		dict[item.GetCode()] = item
	}
	return dict, nil
}

func ArrayOfMap[T ItemAble](dict map[string]T) []T {
	if len(dict) == 0 {
		return []T{}
	}
	arr := make(ItemArray, len(dict))
	i := 0
	for _, item := range dict {
		arr[i] = item
		i += 1
	}
	sort.Sort(arr)
	tArr := make([]T, arr.Len())
	for i, m := range arr {
		tArr[i] = m.(T)
	}
	return tArr
}
