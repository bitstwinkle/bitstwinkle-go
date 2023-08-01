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

package endpoint

import (
	"bitstwinkle-go/types/errors"
	"bitstwinkle-go/types/strs"
)

// Container Terminal container
type Container = string

// Form Existence form
type Form = string

// Endpoint Terminal
type Endpoint struct {
	Container Container `bson:"container" json:"container"`
	Form      Form      `bson:"form" json:"form"`
}

func (ep Endpoint) Verify() *errors.Error {
	if ep.Container == strs.EMPTY {
		return errors.Verify("require container")
	}
	if ep.Form == strs.EMPTY {
		return errors.Verify("require form")
	}
	return nil
}

func (ep Endpoint) ToString() string {
	return ep.Container + ":" + ep.Form
}

func (ep Endpoint) ContainerIn(arr ...Container) bool {
	if len(arr) == 0 {
		return false
	}
	for _, c := range arr {
		if ep.Container == c {
			return true
		}
	}
	return false
}

func (ep Endpoint) ContainerNotIn(arr ...Container) bool {
	if len(arr) == 0 {
		return true
	}

	for _, c := range arr {
		if ep.Container == c {
			return false
		}
	}
	return true
}

func (ep Endpoint) FormIn(arr ...Form) bool {
	if len(arr) == 0 {
		return false
	}
	for _, c := range arr {
		if ep.Form == c {
			return true
		}
	}
	return false
}

func (ep Endpoint) FormNotIn(arr ...Form) bool {
	if len(arr) == 0 {
		return true
	}

	for _, c := range arr {
		if ep.Form == c {
			return false
		}
	}
	return true
}
