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
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
)

// Option 可选值
type Option struct {
	Code  string      `bson:"code" json:"code"`
	Label label.Label `bson:"label" json:"label"`
}

func (opt Option) Verify() *errors.Error {
	if opt.Code == strs.EMPTY {
		return errors.Verify("require code")
	}
	return opt.Label.Verify()
}

type OptionArray []Option

func (arr OptionArray) Verify() *errors.Error {
	for _, item := range arr {
		if err := item.Verify(); err != nil {
			return err
		}
	}
	return nil
}

func (arr OptionArray) GetOption(optionCode string) (Option, *errors.Error) {
	for _, defOption := range arr {
		if optionCode == defOption.Code {
			return defOption, nil
		}
	}
	return Option{}, nil
}

// Definition 规格定义
type Definition struct {
	Code        string      `bson:"code" json:"code"`                 //键值
	Name        string      `bson:"name" json:"name"`                 //名称
	OptionArray OptionArray `bson:"option_array" json:"option_array"` //枚举备选值
	Seq         int         `bson:"seq" json:"seq"`                   //排序值
}

func (d Definition) GetCode() string {
	return d.Code
}

func (d Definition) Verify() *errors.Error {
	if d.Code == strs.EMPTY {
		return errors.Verify("require code")
	}
	if d.Name == strs.EMPTY {
		return errors.Verify("require name")
	}
	if len(d.OptionArray) == 0 {
		return errors.Verify("require option_array")
	}
	if err := d.OptionArray.Verify(); err != nil {
		return errors.Verify("invalid option_array: " + err.Error())
	}
	return nil
}

func (d Definition) GetSeq() int {
	return d.Seq
}

type DefinitionSet struct {
	Newest  []Definition `json:"spec"`    //要新增或更新的规格
	Removed []string     `json:"removed"` //要删除的规格
}

func (m DefinitionSet) Verify() *errors.Error {
	if len(m.Newest) > 0 {
		for _, item := range m.Newest {
			if err := item.Verify(); err != nil {
				return err
			}
		}
	}
	return nil
}
