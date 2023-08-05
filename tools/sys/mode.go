/*
 *
 *  *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *  *
 *  *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  *  * you may not use this file except in compliance with the License.
 *  *  * You may obtain a copy of the License at
 *  *  *
 *  *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *  *
 *  *  * Unless required by applicable law or agreed to in writing, software
 *  *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  *  * See the License for the specific language governing permissions and
 *  *  * limitations under the License.

 */

package sys

// Mode Run Mode Defined
type Mode string

// [dev|local|test|pre|prod]
const (
	LOCAL Mode = "local"
	DEV   Mode = "dev"
	TEST  Mode = "test"
	PRE   Mode = "pre"
	PROD  Mode = "prod"
)

func ModeValueOf(mode string) Mode {
	switch mode {
	case string(LOCAL):
		return LOCAL
	case string(DEV):
		return DEV
	case string(TEST):
		return TEST
	case string(PRE):
		return PRE
	case string(PROD):
		return PROD
	}
	return LOCAL
}

func (r Mode) IsRd() bool {
	return r.IsLocal() ||
		r.IsDev() ||
		r.IsTest()
}

func (r Mode) IsLocal() bool {
	return r == LOCAL
}

func (r Mode) IsDev() bool {
	return r == DEV
}

func (r Mode) IsTest() bool {
	return r == TEST
}

func (r Mode) IsPre() bool {
	return r == PRE
}

func (r Mode) IsProd() bool {
	return r == PROD
}
