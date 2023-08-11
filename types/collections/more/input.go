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

package more

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

type Input struct {
	Alias string `bson:"alias" json:"alias"` //别名,可以设置多个别名,通过","(半角逗号)分割
	Intro string `bson:"intro" json:"intro"` //主介绍
	More  Array  `bson:"more" json:"more"`   //更多信息内容

	_dict More //INNER USE
}

func (m *Input) Verify(requireKey ...string) *errors.Error {
	if err := collections.VerifyArray[Item](m.More); err != nil {
		return err
	}
	dict := m.ToMore()
	if len(requireKey) > 0 {
		for _, key := range requireKey {
			_, ok := dict[key]
			if !ok {
				return errors.Verify("require " + key)
			}
		}
	}
	return nil
}

func (m *Input) ToMore() More {
	if m._dict != nil {
		return m._dict
	}
	m._dict = OfArray(m.More)
	if m.Alias != strs.EMPTY {
		m._dict["alias"] = Item{
			Key:   "alias",
			Value: m.Alias,
			Seq:   0,
		}
	}
	if m.Intro != strs.EMPTY {
		m._dict["intro"] = Item{
			Key:   "intro",
			Value: m.Intro,
			Seq:   1,
		}
	}
	return m._dict
}

func InputToMore(in *Input) More {
	if in == nil {
		return More{}
	}
	return in.ToMore()
}
