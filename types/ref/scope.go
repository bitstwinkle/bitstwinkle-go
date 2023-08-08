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
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
)

// Scope Business scope definition
type Scope struct {
	VN   ww.VN  `bson:"vn" json:"vn"`
	JD   ww.JD  `bson:"jd" json:"jd"`
	Code string `bson:"code" json:"code"`
}

func (m Scope) Verify() *errors.Error {
	if m.VN == strs.EMPTY {
		return errors.Verify("require scope.vn")
	}
	if m.JD == strs.EMPTY {
		return errors.Verify("require scope.jd")
	}
	if m.Code == strs.EMPTY {
		return errors.Verify("require scope.code")
	}
	return nil
}

func (m Scope) String() string {
	return "bitstwinkle://" + m.VN + "/" + m.JD + "/" + m.Code
}

func (m Scope) Same(other Scope) bool {
	return m.VN == other.VN && m.JD == other.JD && m.Code == other.Code
}
