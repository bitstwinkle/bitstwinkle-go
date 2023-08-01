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
)

// Scope Business scope definition
type Scope struct {
	Owner Collar `bson:"owner" json:"owner"`
	Code  string `bson:"code" json:"code"`
}

func (m Scope) Verify() *errors.Error {
	if err := m.Owner.Verify(); err != nil {
		return err
	}
	if m.Code == strs.EMPTY {
		return errors.Verify("require scope.code")
	}
	return nil
}

func (m Scope) String() string {
	return m.Owner.String() + ":" + m.Code
}

func (m Scope) Same(other Scope) bool {
	return m.Owner.Same(other.Owner) && m.Code == other.Code
}
