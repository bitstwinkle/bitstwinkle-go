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
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

// Link Connect to an external entity
type Link struct {
	Code  string `bson:"code" json:"code"`
	ID    string `bson:"id" json:"id"`
	Alias string `bson:"alias" json:"alias"`
	Seq   int    `bson:"seq" json:"seq"`
}

func (m Link) GetCode() string {
	return m.Code
}

func (m Link) Verify() *errors.Error {
	if m.Code == strs.EMPTY {
		return errors.Verify("require link.code")
	}
	if m.ID == strs.EMPTY {
		return errors.Verify("require link.id")
	}
	if m.Alias == strs.EMPTY {
		return errors.Verify("require link.alias")
	}
	return nil
}

func (m Link) GetSeq() int {
	return m.Seq
}

func (m Link) VerifyWithCode(code ...string) *errors.Error {
	if err := m.Verify(); err != nil {
		return err
	}
	if len(code) > 0 {
		include := false
		for _, c := range code {
			if m.Code == c {
				include = true
				break
			}
		}
		if !include {
			return errors.Verify(fmt.Sprintf("link.code not in %s", code))
		}
	}
	return nil
}

type LineArray []Link
type LinkMore map[string]Link
