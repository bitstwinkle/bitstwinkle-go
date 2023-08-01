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

package label

import (
	"bitstwinkle-go/types/collections"
	"bitstwinkle-go/types/collections/more"
	"bitstwinkle-go/types/errors"
	"bitstwinkle-go/types/strs"
)

type Label struct {
	Label string     `bson:"label" json:"label"`                     //[*]标签
	Code  string     `bson:"code,omitempty" json:"code,omitempty"`   //[-]自定义编码
	Icon  string     `bson:"icon,omitempty" json:"icon,omitempty"`   //[-]图标
	Color string     `bson:"color,omitempty" json:"color,omitempty"` //[-]颜色标注
	Meta  more.Array `bson:"meta,omitempty" json:"meta,omitempty"`   //[-]更多元数据
	Seq   int        `bson:"seq" json:"seq"`                         //[-]排序
}

func (m Label) GetCode() string {
	return m.Label
}

func (m Label) GetSeq() int {
	return m.Seq
}

func (m Label) Verify() *errors.Error {
	if m.Label == strs.EMPTY {
		return errors.Verify("require label")
	}
	if len(m.Meta) > 0 {
		if err := m.Meta.Verify(); err != nil {
			return err
		}
	}
	return nil
}

type Array []Label

func (arr Array) Verify() *errors.Error {
	return collections.VerifyArray[Label](arr)
}
