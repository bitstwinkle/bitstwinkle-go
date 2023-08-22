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
	"encoding/base64"
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"strings"
)

type Collar struct {
	Code string `bson:"code" json:"code"`
	ID   string `bson:"id" json:"id"`
}

type Loader interface {
	Get(collar *Collar) (*Ref, *errors.Error)
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

func (c Collar) Unique() string {
	uniStr := c.Code + ":" + c.ID
	return base64.StdEncoding.EncodeToString([]byte(uniStr))
}

func CollarOf(uni string) Collar {
	bData, err := base64.StdEncoding.DecodeString(uni)
	if err != nil {
		return Collar{
			Code: strs.EMPTY,
			ID:   strs.EMPTY,
		}
	}
	arr := strings.Split(string(bData), ":")
	if len(arr) != 2 {
		return Collar{
			Code: strs.EMPTY,
			ID:   strs.EMPTY,
		}
	}
	return Collar{
		Code: arr[0],
		ID:   arr[1],
	}
}

func (c Collar) Same(other Collar) bool {
	return c.Code == other.Code && c.ID == other.ID
}

type CollarSet struct {
	Yes   bool    `bson:"yes" json:"yes"`
	Value *Collar `bson:"value,omitempty" json:"value,omitempty"`
}
