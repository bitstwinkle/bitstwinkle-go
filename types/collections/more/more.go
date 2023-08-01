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

package more

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

type Item struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"value" json:"value"`
	Seq   int    `bson:"seq" json:"seq"`
}

func (m Item) GetCode() string {
	return m.Key
}

func (m Item) GetSeq() int {
	return m.Seq
}

func (m Item) Verify() *errors.Error {
	if m.Key == strs.EMPTY {
		return errors.Verify("require key")
	}
	if m.Value == strs.EMPTY {
		return errors.Verify("require value")
	}
	return nil
}

type Array []Item

func (arr Array) Verify() *errors.Error {
	return collections.VerifyArray[Item](arr)
}

type More map[string]Item

func NewMore() More {
	return make(More)
}

func OfArray(arr Array) (More, *errors.Error) {
	return collections.MapOfArray[Item](arr)
}

func (m More) Put(key string, value string, seq ...int) More {
	item := Item{
		Key:   key,
		Value: value,
		Seq:   0,
	}
	if len(seq) > 0 {
		item.Seq = seq[0]
	}
	m[key] = item
	return m
}

func (m More) Get(key string) string {
	item, ok := m[key]
	if !ok {
		return strs.EMPTY
	}
	return item.Value
}

func (m More) ToArray() Array {
	return collections.ArrayOfMap[Item](m)
}

type Set struct {
	Newest  Array    `json:"newest"`  //新增或更新
	Removed []string `json:"removed"` //需要移除的
}
