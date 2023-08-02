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

package spec

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

type Value struct {
	Code   string `bson:"code" json:"code"`     //规格定义编码
	Option string `bson:"option" json:"option"` //规格定义中备选值编码
	Seq    int    `bson:"seq" json:"seq"`
}

func (s Value) GetCode() string {
	return s.Code
}

func (s Value) Verify() *errors.Error {
	if s.Code == strs.EMPTY {
		return errors.Verify("require code")
	}
	if s.Option == strs.EMPTY {
		return errors.Verify("require option")
	}
	return nil
}

func (s Value) GetSeq() int {
	return s.Seq
}

type Spec struct {
	Code   string `bson:"code" json:"code"`     //规格定义的编码
	Name   string `bson:"name" json:"name"`     //规格定义的名称
	Option Option `bson:"option" json:"option"` //备选值
	Seq    int    `bson:"seq" json:"seq"`       //排序
}

func (s Spec) GetCode() string {
	return s.Option.Code
}

func (s Spec) Verify() *errors.Error {
	if s.Code == strs.EMPTY {
		return errors.Verify("require code")
	}
	if s.Name == strs.EMPTY {
		return errors.Verify("require name")
	}
	if err := s.Option.Verify(); err != nil {
		return errors.Verify("invalid option: " + err.Error())
	}
	return nil
}

func (s Spec) GetSeq() int {
	return s.Seq
}

func Of(def Definition, val Value) (Spec, *errors.Error) {
	if def.Code != val.Code {
		return Spec{}, errors.Assert("def.Code == val.Code", "def.Code != val.Code")
	}
	if len(def.OptionArray) == 0 {
		return Spec{}, errors.Assert("def.OptionArray.len > 0", "def.OptionArray.len == 0")
	}
	matchedDefOption, err := def.OptionArray.GetOption(val.Option)
	if err != nil {
		return Spec{}, err
	}
	return Spec{
		Code:   def.Code,
		Name:   def.Name,
		Option: matchedDefOption,
		Seq:    def.Seq,
	}, nil
}

type Set struct {
	Set     bool     `json:"set"`     //是否已设置
	Newest  []Value  `json:"spec"`    //要新增或更新的规格
	Removed []string `json:"removed"` //要删除的规格
}

func (m Set) Verify() *errors.Error {
	if len(m.Newest) > 0 {
		for _, item := range m.Newest {
			if err := item.Verify(); err != nil {
				return err
			}
		}
	}
	return nil
}
