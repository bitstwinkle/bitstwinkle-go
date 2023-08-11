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

package media

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
)

type Input struct {
	Logo    *Media `bson:"logo,omitempty" json:"logo,omitempty"`       //LOGO
	Avatar  *Media `bson:"avatar,omitempty" json:"avatar,omitempty"`   //Avatar
	Primary *Media `bson:"primary,omitempty" json:"primary,omitempty"` //Primary
	More    Array  `bson:"more" json:"more"`                           //More Media

	_dict More //INNER USE
}

func (m *Input) Verify(requireKey ...string) *errors.Error {
	if err := collections.VerifyArray[Item](m.More); err != nil {
		return err
	}
	if m.Logo != nil {
		if err := m.Logo.Verify(); err != nil {
			return errors.Verify("invalid logo: " + err.Error())
		}
	}
	if m.Avatar != nil {
		if err := m.Avatar.Verify(); err != nil {
			return errors.Verify("invalid avatar: " + err.Error())
		}
	}
	if m.Primary != nil {
		if err := m.Primary.Verify(); err != nil {
			return errors.Verify("invalid primary: " + err.Error())
		}
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
	if m.Logo != nil {
		m._dict["logo"] = Item{
			Code:       "logo",
			Seq:        0,
			MediaArray: []*Media{m.Logo},
		}
	}
	if m.Avatar != nil {
		m._dict["avatar"] = Item{
			Code:       "avatar",
			Seq:        1,
			MediaArray: []*Media{m.Avatar},
		}
	}
	if m.Primary != nil {
		m._dict["primary"] = Item{
			Code:       "primary",
			Seq:        2,
			MediaArray: []*Media{m.Primary},
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
