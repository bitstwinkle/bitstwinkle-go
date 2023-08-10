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

type Input struct {
	Logo    *Media `bson:"logo,omitempty" json:"logo,omitempty"`       //LOGO
	Avatar  *Media `bson:"avatar,omitempty" json:"avatar,omitempty"`   //Avatar
	Primary *Media `bson:"primary,omitempty" json:"primary,omitempty"` //Primary
	More    Array  `bson:"more" json:"more"`                           //More Media
}

func (m *Input) ToMore() More {
	dict := OfArray(m.More)
	if m.Logo != nil {
		dict["logo"] = Item{
			Code:       "logo",
			Seq:        0,
			MediaArray: []*Media{m.Logo},
		}
	}
	if m.Avatar != nil {
		dict["logo"] = Item{
			Code:       "avatar",
			Seq:        1,
			MediaArray: []*Media{m.Avatar},
		}
	}
	if m.Primary != nil {
		dict["primary"] = Item{
			Code:       "primary",
			Seq:        2,
			MediaArray: []*Media{m.Primary},
		}
	}
	return dict
}
