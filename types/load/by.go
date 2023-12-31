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

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"strings"
)

type ByCode = string

type By struct {
	handlers map[ByCode]func() *errors.Error
}

func NewBy() *By {
	return &By{handlers: make(map[ByCode]func() *errors.Error)}
}

func (by *By) Verify(code ByCode) *errors.Error {
	_, ok := by.handlers[code]
	if !ok {
		return errors.Verify("can not use by: " + code)
	}
	return nil
}

func (by *By) Register(code ByCode, handle func() *errors.Error) *By {
	by.handlers[code] = handle
	return by
}

func (by *By) Do(byCode ByCode) *errors.Error {
	call, ok := by.handlers[byCode]
	if !ok {
		return errors.Verify("[by] must be: [" + strings.Join(by.AllByCode(), "|") + "], but: [" + byCode + "]")
	}
	return call()
}

func (by *By) AllByCode() []string {
	all := make([]string, len(by.handlers))
	i := 0
	for k, _ := range by.handlers {
		all[i] = k
		i += 1
	}
	return all
}
