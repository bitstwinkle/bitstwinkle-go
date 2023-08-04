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

package address

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
)

type Service interface {
	AreaLoad(req AreaLoadRequest) ([]Area, load.Paging, *errors.Error)
	ZoneLoad(req ZoneLoadRequest) ([]Zone, load.Paging, *errors.Error)
	AddrRegister(req AddrRegisterRequest) (*Address, *errors.Error)
	AddrLoad(req AddrLoadRequest) ([]Area, load.Paging, *errors.Error)
	AddrSet(req AddrSetRequest) (*Address, *errors.Error)
}
