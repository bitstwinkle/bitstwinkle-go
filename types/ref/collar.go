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

type Collar struct {
	Code string `bson:"code" json:"code"`
	ID   string `bson:"id" json:"id"`
}

func (c Collar) Verify(code ...string) *errors.Error {
	if c.Code == strs.EMPTY {
		return errors.Verify("require code")
	}
	if c.ID == strs.EMPTY {
		return errors.Verify("require id")
	}
	if len(code) > 0 {
		include := false
		for _, item := range code {
			if c.Code == item {
				include = true
				break
			}
		}
		if !include {
			return errors.Verify(fmt.Sprintf("collar.code expect %s, but %s", code, c.Code))
		}
	}
	return nil
}

func (c Collar) String() string {
	return c.Code + ":" + c.ID
}

func (c Collar) Same(other Collar) bool {
	return c.Code == other.Code && c.ID == other.ID
}
