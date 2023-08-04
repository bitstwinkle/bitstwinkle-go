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

package ref

import (
	"github.com/bitstwinkle/bitstwinkle-go/tools/unique"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

// Lead Business content link unique value
type Lead struct {
	Owner Collar `bson:"owner" json:"owner"`
	Code  string `bson:"code" json:"code"`
}

const JustOne = "_uk_"

// LeadOfOneToOne 一个OWNER下只有一个
func LeadOfOneToOne(owner Collar) Lead {
	return Lead{
		Owner: owner,
		Code:  JustOne,
	}
}

// LeadOfOneToMany 一个OWNER下会有多个
func LeadOfOneToMany(owner Collar, prefix string) Lead {
	return Lead{
		Owner: owner,
		Code:  prefix + "_" + unique.ID(),
	}
}

func (m Lead) IsOneToOne() bool {
	return m.Code == JustOne
}

func (m Lead) IsOneToMany() bool {
	return m.Code != JustOne
}

func (m Lead) Verify() *errors.Error {
	if err := m.Owner.Verify(); err != nil {
		return err
	}
	if m.Code == strs.EMPTY {
		return errors.Verify("require scope.code")
	}
	return nil
}

func (m Lead) String() string {
	return m.Owner.String() + ":" + m.Code
}

func (m Lead) Same(other Lead) bool {
	return m.Owner.Same(other.Owner) && m.Code == other.Code
}
