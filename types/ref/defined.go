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
	"bitstwinkle-go/types/errors"
	"bitstwinkle-go/types/strs"
	"fmt"
)

type Defined struct {
	Code string `bson:"code" json:"code"`
	Key  string `bson:"key" json:"key"`
}

func DefinedOf(code string, key string) Defined {
	return Defined{
		Code: code,
		Key:  key,
	}
}

func (d Defined) String() string {
	return d.Code + ":" + d.Key
}

func (d Defined) Verify(code ...string) *errors.Error {
	if d.Code == strs.EMPTY {
		return errors.Verify("require defined.code")
	}
	if d.Key == strs.EMPTY {
		return errors.Verify("require defined.key")
	}
	if len(code) > 0 {
		include := false
		for _, c := range code {
			if d.Code == c {
				include = true
				break
			}
		}
		if !include {
			return errors.Verify(fmt.Sprintf("defined.code not in %s", code))
		}
	}
	return nil
}
