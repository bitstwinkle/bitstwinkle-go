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

package load

import "bitstwinkle-go/types/errors"

type By struct {
	handlers map[string]func() *errors.Error
}

func NewBy() *By {
	return &By{handlers: make(map[string]func() *errors.Error)}
}

func (by *By) Register(code string, handle func() *errors.Error) *By {
	by.handlers[code] = handle
	return by
}

func (by *By) Do(byCode string) *errors.Error {
	call, ok := by.handlers[byCode]
	if !ok {
		return errors.Assert(byCode+".call", "nil")
	}
	return call()
}
