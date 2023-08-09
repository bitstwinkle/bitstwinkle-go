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

package media

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

type Type = string

const (
	Photo Type = "photo"
	Video Type = "video"
)

type Media struct {
	Type Type       `bson:"t" json:"t"`
	Link string     `bson:"link" json:"link"`
	Meta more.Array `bson:"meta" json:"meta"`
}

func (m Media) Verify() *errors.Error {
	if m.Type == strs.EMPTY {
		return errors.Verify("media.type is empty")
	}
	if m.Link == strs.EMPTY {
		return errors.Verify("media.link is empty")
	}
	if err := collections.VerifyArray[more.Item](m.Meta); err != nil {
		return err
	}
	return nil
}

type Item struct {
	Code       string   `bson:"code" json:"code"`
	Seq        int      `bson:"seq" json:"seq"`
	MediaArray []*Media `bson:"media_array" json:"media_array"`
}

func (i Item) GetCode() string {
	return i.Code
}

func (i Item) Verify() *errors.Error {
	if len(i.MediaArray) == 0 {
		return errors.Verify("require media")
	}
	for _, each := range i.MediaArray {
		if err := each.Verify(); err != nil {
			return err
		}
	}
	return nil
}

func (i Item) GetSeq() int {
	return i.Seq
}

type Array []Item

type More map[string]Item

func NewMore() More {
	return make(More)
}

func OfArray(arr Array) More {
	m, _ := collections.MapOfArray[Item](arr)
	return m
}

func (m More) ToArray() Array {
	return collections.ArrayOfMap[Item](m)
}

type Set struct {
	Newest  Array    `json:"newest"`  //新增或更新
	Removed []string `json:"removed"` //需要移除的
}
