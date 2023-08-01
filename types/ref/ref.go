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
	"time"
)

// Able Every principal structure should implement this interface, such as users, merchants, orders, and so on
type Able interface {
	Ref() Ref
}

type Ref struct {
	VN    string    `bson:"vn" json:"vn"`
	Code  string    `bson:"code" json:"code"`
	ID    string    `bson:"id" json:"id"`
	Alias string    `bson:"alias" json:"alias"`
	Sync  time.Time `bson:"sync" json:"sync"`
}

func Of(vn string, code string, id string, alias string) Ref {
	return Ref{
		VN:    vn,
		Code:  code,
		ID:    id,
		Alias: alias,
		Sync:  time.Now(),
	}
}

func (r Ref) String() string {
	return fmt.Sprintf("vn:%s|code:%s|id:%salias:%s", r.VN, r.Code, r.ID, r.Alias)
}

func (r Ref) Verify(code ...string) *errors.Error {
	if r.VN == strs.EMPTY {
		return errors.Verify("require vn")
	}
	if r.Code == strs.EMPTY {
		return errors.Verify("require code")
	}
	if r.ID == strs.EMPTY {
		return errors.Verify("require id")
	}
	if r.Alias == strs.EMPTY {
		return errors.Verify("require alias")
	}

	if len(code) > 0 {
		include := false
		for _, c := range code {
			if r.Code == c {
				include = true
				break
			}
		}
		if !include {
			return errors.Verify(fmt.Sprintf("rer.code expect %s, but %s", code, r.Code))
		}
	}

	return nil
}

func (r Ref) Collar() Collar {
	return Collar{
		Code: r.Code,
		ID:   r.ID,
	}
}

func (r Ref) MustInVN(vnID string) *errors.Error {
	if r.VN != vnID {
		return errors.Verify("not in vn: " + vnID)
	}
	return nil
}

func (r Ref) Same(other Ref) bool {
	return r.VN == other.VN && r.Code == other.Code && r.ID == other.ID
}

type Getter interface {
	Get(collar *Collar) (*Ref, *errors.Error)
}
