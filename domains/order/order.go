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

package order

import "github.com/bitstwinkle/bitstwinkle-go/types/errors"

type Service interface {
	Create(req CreateRequest) (*Order, *errors.Error)
	Get(req GetRequest) (*Order, *errors.Error)
	Set(req SetRequest) *errors.Error
	Advance(req AdvanceRequest) *errors.Error
}
